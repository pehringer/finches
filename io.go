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

func float(b bool) float64 {
  if b {
    return 1
  }
  return 0
}

func divide(n, d float64) float64 {
  if math.Abs(d) > 1e-9 {
    return n / d
  }
  if math.Abs(n) < 1e-9 {
    return math.NaN()
  }
  if n > 0 {
    return math.Inf(1)
  }
  return math.Inf(-1)
}

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

func main() {
  for i := 1; i < len(os.Args); i += %d {
%s  }
  fmt.Println(r[15])
}`

func write(path string, inputs int, constants []float64, instructions []uint16) error {
	code := ""
	for i := range inputs {
		code += fmt.Sprintf("    r[%d], _ = strconv.ParseFloat(os.Args[i+%d], 64)\n", i, i)
	}
	for i := range instructions {
		result := fmt.Sprintf("r[%d]", instructions[i] >> resultShift & shiftMask)
		first := fmt.Sprintf("r[%d]", instructions[i] >> firstShift & shiftMask)
		second := fmt.Sprintf("r[%d]", instructions[i] >> secondShift & shiftMask)
		switch instructions[i] & opcodeMask {
		case opcodeAD:
			code += fmt.Sprintf("    %s = %s + %s\n", result, first, second)
		case opcodeSB:
			code += fmt.Sprintf("    %s = %s - %s\n", result, first, second)
		case opcodeML:
			code += fmt.Sprintf("    %s = %s * %s\n", result, first, second)
		case opcodeDV:
			code += fmt.Sprintf("    %s = divide(%s, %s)\n", result, first, second)
		case opcodePW:
			code += fmt.Sprintf("    %s = math.Pow(%s, %s)\n", result, first, second)
		case opcodeSQ:
			code += fmt.Sprintf("    %s = math.Sqrt(%s)\n",  result, first)
		case opcodeEX:
			code += fmt.Sprintf("    %s = math.Exp(%s)\n", result, first)
		case opcodeLG:
			code += fmt.Sprintf("    %s = math.Log(%s)\n", result, first)
		case opcodeSN:
			code += fmt.Sprintf("    %s = math.Sin(%s)\n", result, first)
		case opcodeAS:
			code += fmt.Sprintf("    %s = math.Asin(%s)\n", result, first)
		case opcodeCS:
			code += fmt.Sprintf("    %s = math.Cos(%s)\n", result, first)
		case opcodeAC:
			code += fmt.Sprintf("    %s = math.Acos(%s)\n", result, first)
		case opcodeMN:
			code += fmt.Sprintf("    %s = math.Min(%s, %s)\n", result, first, second)
		case opcodeMX:
			code += fmt.Sprintf("    %s = math.Max(%s, %s)\n", result, first, second)
		case opcodeLT:
			code += fmt.Sprintf("    %s = float(%s < %s)\n", result, first, second)
		case opcodeGT:
			code += fmt.Sprintf("    %s = float(%s > %s)\n", result, first, second)
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
