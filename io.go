package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInputs(values []string) ([]float64, error) {
	result := make([]float64, len(values))
	for i := range values {
		trimmed := strings.TrimSpace(values[i])
		if trimmed == "" {
			return nil, fmt.Errorf("column %d: do data", i+1)
		}
		float, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return nil, fmt.Errorf("column %d: %w", i+1, err)
		}
		result[i] = float
	}
	return result, nil
}

func parseOutputs(values []string) ([]*float64, error) {
	result := make([]*float64, len(values))
	for i := range values {
		trimmed := strings.TrimSpace(values[i])
		if trimmed == "" {
			continue
		}
		float, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return nil, fmt.Errorf("column %d: %w", i+1, err)
		}
		result[i] = &float
	}
	return result, nil
}

func readExamples(filename string) ([][]float64, [][]*float64, error) {
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
	index := len(values[0]) - 1
	inputs := make([][]float64, len(values))
	outputs := make([][]*float64, len(values))
	for i := range len(values) {
		inputs[i], err = parseInputs(values[i][:index])
		if err != nil {
			return nil, nil, fmt.Errorf("row %d: %w", i+1, err)
		}
		outputs[i], err = parseOutputs(values[i][index:])
		if err != nil {
			return nil, nil, fmt.Errorf("row %d: %w", i+1, err)
		}
	}
	return inputs, outputs, nil
}


var registerVariable map[int]string = map[int]string{
	0:  "A",
	1:  "B",
	2:  "C",
	3:  "D",
	4:  "E",
	5:  "F",
	6:  "G",
	7:  "H",
	8:  "I",
	9:  "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "O",
	15: "P",
}

var instructionArguments map[uint16]int = map[uint16]int{
	opcodeAD: 3,
	opcodeSB: 3,
	opcodeML: 3,
	opcodeDV: 3,
	opcodePW: 3,
	opcodeSQ: 2,
	opcodeEX: 2,
	opcodeLG: 2,
	opcodeSN: 2,
	opcodeAS: 2,
	opcodeCS: 2,
	opcodeAC: 2,
	opcodeMN: 3,
	opcodeMX: 3,
	opcodeLT: 3,
	opcodeGT: 3,
}

var instructionFormat map[uint16]string = map[uint16]string{
	opcodeAD: "\t\t%s = %s + %s\n",
	opcodeSB: "\t\t%s = %s - %s\n",
	opcodeML: "\t\t%s = %s * %s\n",
	opcodeDV: "\t\t%s = divide(%s, %s)\n",
	opcodePW: "\t\t%s = math.Pow(%s, %s)\n",
	opcodeSQ: "\t\t%s = math.Sqrt(%s)\n",
	opcodeEX: "\t\t%s = math.Exp(%s)\n",
	opcodeLG: "\t\t%s = math.Log(%s)\n",
	opcodeSN: "\t\t%s = math.Sin(%s)\n",
	opcodeAS: "\t\t%s = math.Asin(%s)\n",
	opcodeCS: "\t\t%s = math.Cos(%s)\n",
	opcodeAC: "\t\t%s = math.Acos(%s)\n",
	opcodeMN: "\t\t%s = math.Min(%s, %s)\n",
	opcodeMX: "\t\t%s = math.Max(%s, %s)\n",
	opcodeLT: "\t\t%s = float(%s < %s)\n",
	opcodeGT: "\t\t%s = float(%s > %s)\n",
}

const programFormat = `package main
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

%s
func main() {
	for i := 1; i < len(os.Args); i += %s {
%s
%s
		fmt.Println(P)
	}
}`

func writeProgram(path string, inputs int, constants []float64, instructions []uint16) error {
	variables := ""
	for i := range constants {
		name := registerVariable[i]
		value := constants[i]
		variables += fmt.Sprintf("var %s float64 = %f\n", name, value)
	}
	increment := fmt.Sprintf("%d", inputs)
	arguments := ""
	for i := range inputs {
		name := registerVariable[i]
		arguments += fmt.Sprintf("\t\t%s, _ = strconv.ParseFloat(os.Args[i+%d], 64)\n", name, i)
	}
	code := ""
	for i := range instructions {
		argument1 := registerVariable[int(instructions[i] >> resultShift & shiftMask)]
		argument2 := registerVariable[int(instructions[i] >> firstShift & shiftMask)]
		argument3 := registerVariable[int(instructions[i] >> secondShift & shiftMask)]
		opcode := instructions[i] & opcodeMask
		switch instructionArguments[opcode] {
		case 1:
			code += fmt.Sprintf(instructionFormat[opcode], argument1)
		case 2:
			code += fmt.Sprintf(instructionFormat[opcode], argument1, argument2)
		case 3:
			code += fmt.Sprintf(instructionFormat[opcode], argument1, argument2, argument3)
		}
	}
	program := fmt.Sprintf(programFormat, variables, increment, arguments, code)
	err := os.WriteFile(path, []byte(program), 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}
	return nil
}
