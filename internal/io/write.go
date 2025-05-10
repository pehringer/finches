package io

import (
	"fmt"
	"os"

	"github.com/pehringer/mapper/internal/vm"
	"github.com/pehringer/mapper/internal/types"
)

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

func WriteProgram(filepath string, program types.Program) error {
	assembly := ""
		for i := range program.Data {
		assembly += fmt.Sprintf("%02d %f\n", i, program.Data[i])
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
