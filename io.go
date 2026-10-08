package main

import (
	"math"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(values []string) ([]float64, error) {
	result := make([]float64, len(values))
	for i := range values {
		result[i] = math.NaN()
		trimmed := strings.TrimSpace(values[i])
		if trimmed == "" {
			continue
		}
		float, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return nil, fmt.Errorf("column %d: %w", i+1, err)
		}
		result[i] = float
	}
	return result, nil
}

func read(filename string) ([][]float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()
	values, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("could not read csv: %w", err)
	}
	if len(values) < 1 {
		return nil, nil, fmt.Errorf("csv format: must contain a least one input-output example row")
	}
	if len(values[0]) < 2 {
		return nil, nil, fmt.Errorf("csv format: missing expected output column (last column)")
	}
	if len(values[0]) > 9 {
		return nil, nil, fmt.Errorf("csv format: too many example input columns (maximum first eight columns)")
	}
	inputs := make([][]float64, len(values))
	outputs := make([]float64, len(values))
	for i := range len(values) {
		inouts, err := parse(values[i])
		if err != nil {
			return nil, nil, fmt.Errorf("row %d: %w", i+1, err)
		}
		index := len(inouts) - 1
		inputs[i] = inouts[:index]
		outputs[i] = inouts[index]
	}
	return inputs, outputs, nil
}

const programFormat = `
package main

import "os"
import "fmt"
import "math"
import "strconv"

var r = [16]float64{
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
  %f,
}

func mix(x, y, w float64) float64 {
  if w >= 1.0 {
    return y
  }
  if w <= 0.0 {
    return x
  }
  return (1.0 - w) * x + w * y
}

func sig(x, y float64) float64 {
  return 1.0 / (1.0 + math.Exp(x - y))
}

func main() {
  for i := 1; i < len(os.Args); i += %d {
%s  }
  fmt.Println(r[15])
}`

func write(path string, inputs int, constants []float64, instructions []int) error {
	code := ""
	for i := range inputs {
		code += fmt.Sprintf("    r[%d], _ = strconv.ParseFloat(os.Args[i+%d], 64)\n", i, i)
	}
	for i := range instructions {
		second := fmt.Sprintf("r[%d]", instructions[i] >> secondShift & secondMask)
		first := fmt.Sprintf("r[%d]", instructions[i] >> firstShift & firstMask)
		result := fmt.Sprintf("r[%d]", instructions[i] >> resultShift & resultMask)
		opcode := instructions[i] & opcodeMask
		switch opcode {
		case opcodeADD:
			code += fmt.Sprintf("    %s = %s + %s\n", result, first, second)
		case opcodeSUB:
			code += fmt.Sprintf("    %s = %s - %s\n", result, first, second)
		case opcodeMUL:
			code += fmt.Sprintf("    %s = %s * %s\n", result, first, second)
		case opcodeDIV:
			code += fmt.Sprintf("    %s = %s / %s\n", result, first, second)
		case opcodeMAX:
			code += fmt.Sprintf("    %s = math.Max(%s, %s)\n", result, first, second)
		case opcodeMIN:
			code += fmt.Sprintf("    %s = math.Min(%s, %s)\n", result, first, second)
		case opcodeMIX:
			code += fmt.Sprintf("    %s = mix(%s, %s, %s)\n", result, result, first, second)
		case opcodeSIG:
			code += fmt.Sprintf("    %s = sig(%s, %s)\n",  result, first, second)
		}
	}
	program := fmt.Sprintf(programFormat,
		constants[0],
		constants[1],
		constants[2],
		constants[3],
		constants[4],
		constants[5],
		constants[6],
		constants[7],
		constants[8],
		constants[9],
		constants[10],
		constants[11],
		constants[12],
		constants[13],
		constants[14],
		constants[15],
		inputs,
		code,
	)
	err := os.WriteFile(path, []byte(program), 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}
	return nil
}
