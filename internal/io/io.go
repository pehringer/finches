package io

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pehringer/fungen/internal/ga"
	"github.com/pehringer/fungen/internal/vm"
)

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	values, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read csv: %w", err)
	}
	return values, nil
}

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

func ReadTests(filename string) ([]ga.Test, error) {
	values, err := readCSV(filename)
	if err != nil {
		return nil, err
	}
	if len(values[0]) != 2 {
		return nil, fmt.Errorf("expected 2 columns but has: %d", len(values[0]))
	}
	result := make([]ga.Test, len(values)-1)
	for i := 1; i < len(values); i++ {
		columns, err := parseFloats(values[i])
		if err != nil {
			return nil, fmt.Errorf("row %d: %w", i+1, err)
		}
		result[i-1].Input = columns[0]
		result[i-1].Output = columns[1]
	}
	return result, nil
}

func parseInstruction(instruction uint16) string {
	result := ""
	condition := instruction & vm.Condition
	switch condition {
	case vm.ConditionLT:
		result += "LT"
	case vm.ConditionLE:
		result += "LE"
	case vm.ConditionEQ:
		result += "EQ"
	case vm.ConditionNE:
		result += "NE"
	case vm.ConditionGE:
		result += "GE"
	case vm.ConditionGT:
		result += "GT"
	case vm.ConditionNV:
		result += "NV"
	}
	operation := instruction & vm.Operation
	switch operation {
	case vm.OperationLD:
		result += "LD"
	case vm.OperationST:
		result += "ST"
	case vm.OperationAD:
		result += "AD"
	case vm.OperationSB:
		result += "SB"
	case vm.OperationML:
		result += "ML"
	case vm.OperationDV:
		result += "DV"
	case vm.OperationMX:
		result += "MX"
	case vm.OperationMN:
		result += "MN"
	case vm.OperationAB:
		result += "AB"
	case vm.OperationPW:
		result += "PW"
	case vm.OperationSQ:
		result += "SQ"
	case vm.OperationEX:
		result += "EX"
	case vm.OperationLG:
		result += "LG"
	case vm.OperationSN:
		result += "SN"
	case vm.OperationCS:
		result += "CS"
	case vm.OperationTN:
		result += "TN"
	}
	setFlag := instruction & vm.SetFlag
	switch setFlag {
	case vm.SetFlagS:
        	result += "S"
	}
	address := instruction & vm.Address
	result += fmt.Sprintf(" %02d\n", address)
	return result
}

func WriteProgram(filepath string, program *ga.Program) error {
	assembly := ""
		for i := range program.Memory {
		assembly += fmt.Sprintf("%02d %f\n", i, program.Memory[i])
	}
	for i := range program.Instructions {
		assembly += parseInstruction(program.Instructions[i])
	}
	err := os.WriteFile(filepath, []byte(assembly), 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}
	return nil
}
