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

func parseFloats(values []string) ([]float32, error) {
	result := make([]float32, len(values))
	for i := range values {
		trimmed := strings.TrimSpace(values[i])
		float, err := strconv.ParseFloat(trimmed, 32)
		if err != nil {
			return nil, fmt.Errorf("column %d: %w", i+1, err)
		}
		result[i] = float32(float)
	}
	return result, nil
}

func ReadTests(filename string) ([]ga.Test, error) {
	values, err := readCSV(filename)
	if err != nil {
		return nil, err
	}
	inputs := 0
	expected := 0
	for i := range values[0] {
		trimmed := strings.TrimSpace(values[0][i])
		if len(trimmed) >= 5 && trimmed[:5] == "input" {
			inputs++
			continue
		}
		if len(trimmed) >= 8 && trimmed[:8] == "expected" {
			expected++
			continue
		}
		return nil, fmt.Errorf("invalid header name: %s", trimmed)
	}
	result := make([]ga.Test, len(values)-1)
	for i := 1; i < len(values); i++ {
		columns, err := parseFloats(values[i])
		if err != nil {
			return nil, fmt.Errorf("row %d: %w", i+1, err)
		}
		result[i-1].Inputs = columns[:inputs]
		result[i-1].Expected = columns[inputs:]
	}
	return result, nil
}

func parseInstruction(registers int, instruction uint32) string {
	result := ""
	switch instruction & vm.Condition {
	case vm.ConditionAlways:
		result += "  "
	case vm.ConditionLT:
		result += "LT"
	case vm.ConditionGT:
		result += "GT"
	case vm.ConditionEQ:
		result += "EQ"
	}
	switch instruction & vm.Operation {
	case vm.OperationADD:
		result += "ADD"
	case vm.OperationSUB:
		result += "SUB"
	case vm.OperationMUL:
		result += "MUL"
	case vm.OperationDIV:
		result += "DIV"
	case vm.OperationNOP4:
		return "  NOP\n"
	case vm.OperationNOP5:
		return "  NOP\n"
	case vm.OperationNOP6:
		return "  NOP\n"
	case vm.OperationNOP7:
		return "  NOP\n"
	}
	switch instruction & vm.SetFlags {
	case vm.SetFlagsNo:
		result += "  "
	case vm.SetFlagsS:
		result += "S "
	}
	destination := int((instruction & vm.Destination)>>12) % registers
	result += fmt.Sprintf("R%02d, ", destination)
	source1 := int((instruction & vm.Source1)>>6) % registers
	result += fmt.Sprintf("R%02d, ", source1)
	source2 := int((instruction & vm.Source2)>>0) % registers
	result += fmt.Sprintf("R%02d\n", source2)
	return result
}

func WriteProgram(filepath string, program *ga.Program) error {
	assembly := ""
	for i := range program.Registers {
		assembly += fmt.Sprintf("R%02d %f\n", i, program.Registers[i])
	}
	for i := range program.Instructions {
		assembly += parseInstruction(len(program.Registers), program.Instructions[i])
	}
	err := os.WriteFile(filepath, []byte(assembly), 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}
	return nil
}
