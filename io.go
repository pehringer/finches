package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFloats(values []string) ([]float64, error) {
	result := make([]float64, len(values))
	for i := range values {
		trimmed := strings.TrimSpace(values[i])
		float, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return nil, fmt.Errorf("column %d: %w", i+1, err)
		}
		result[i] = float
	}
	return result, nil
}

func readExamples(filename string) ([][]float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()
	values, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("could not read csv: %w", err)
	}
	if len(values[0]) < 2 {
		return nil, nil, fmt.Errorf("expected 2 or more columns")
	}
	inputs := make([][]float64, len(values))
	outputs := make([]float64, len(values))
	for i := range len(values) {
		columns, err := parseFloats(values[i])
		if err != nil {
			return nil, nil, fmt.Errorf("row %d: %w", i+1, err)
		}
		output := len(columns) - 1
		inputs[i] = columns[:output]
		outputs[i] = columns[output]
	}
	return inputs, outputs, nil
}

func instructionString(instruction uint16) string {
	line := ""
	opcode := instruction & opcodeMask
	switch opcode {
	case opcodeAD:
		line += "AD "
	case opcodeSB:
		line += "SB "
	case opcodeML:
		line += "ML "
	case opcodeDV:
		line += "DV "
	case opcodePW:
		line += "PW "
	case opcodeSQ:
		line += "SQ "
	case opcodeEX:
		line += "EX "
	case opcodeLG:
		line += "LG "
	case opcodeSN:
		line += "SN "
	case opcodeAS:
		line += "AS "
	case opcodeCS:
		line += "CS "
	case opcodeAC:
		line += "AC "
	case opcodeMN:
		line += "MN "
	case opcodeMX:
		line += "MX "
	case opcodeLT:
		line += "LT "
	case opcodeGT:
		line += "GT "
	}
	second := int(instruction >> secondShift & shiftMask)
	first := int(instruction >> firstShift & shiftMask)
	result := int(instruction >> resultShift & shiftMask)
	line += fmt.Sprintf("%02d %02d %02d\n", result, first, second)
	return line
}

func writeProgram(path string, constants []float64, instructions []uint16) error {
	program := ""
	for i := range constants {
		program += fmt.Sprintf("%02d %f\n", i, constants[i])
	}
	for i := range instructions {
		program += instructionString(instructions[i])
	}
	err := os.WriteFile(path, []byte(program), 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}
	return nil
}
