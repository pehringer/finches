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

func parseInstruction(instruction uint16) string {
	state := instruction >> vm.State
	operand := instruction & vm.Operand
	switch instruction & vm.Opcode {
	case vm.OpcodeLD:
		return fmt.Sprintf("%02d: LD R%02d\n", state, operand)
	case vm.OpcodeST:
		return fmt.Sprintf("%02d: ST R%02d\n", state, operand)
	case vm.OpcodeAD:
		return fmt.Sprintf("%02d: AD R%02d\n", state, operand)
	case vm.OpcodeSB:
		return fmt.Sprintf("%02d: SB R%02d\n", state, operand)
	case vm.OpcodeML:
		return fmt.Sprintf("%02d: ML R%02d\n", state, operand)
	case vm.OpcodeDV:
		return fmt.Sprintf("%02d: DV R%02d\n", state, operand)
	case vm.OpcodeLT:
		return fmt.Sprintf("%02d: LT R%02d\n", state, operand)
	case vm.OpcodeGT:
		return fmt.Sprintf("%02d: GT R%02d\n", state, operand)
	case vm.OpcodeEQ:
		return fmt.Sprintf("%02d: EQ R%02d\n", state, operand)
	case vm.OpcodeNE:
		return fmt.Sprintf("%02d: NE R%02d\n", state, operand)
	}
	return fmt.Sprintf("    NOP\n")
}

func WriteProgram(filepath string, program *ga.Program) error {
	assembly := ""
		for i := range program.Registers {
		assembly += fmt.Sprintf("R%02d %f\n", i, program.Registers[i])
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
