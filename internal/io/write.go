package io

import (
	"fmt"
	"os"

	"github.com/pehringer/mapper/internal/vm"
	"github.com/pehringer/mapper/internal/types"
)

func parseInstruction(instruction uint16) string {
	predicate := int(instruction >> vm.PredicateShift & vm.ShiftMask)
	line := fmt.Sprintf("%02d ", predicate)
	operation := instruction & vm.OperationMask
	switch operation {
	case vm.OperationAD:
		line += "AD "
	case vm.OperationSB:
		line += "SB "
	case vm.OperationML:
		line += "ML "
	case vm.OperationDV:
		line += "DV "
	case vm.OperationPW:
		line += "PW "
	case vm.OperationSQ:
		line += "SQ "
	case vm.OperationEX:
		line += "EX "
	case vm.OperationLG:
		line += "LG "
	case vm.OperationSN:
		line += "SN "
	case vm.OperationCS:
		line += "CS "
	case vm.OperationTN:
		line += "TN "
	case vm.OperationAB:
		line += "AB "
	case vm.OperationLT:
		line += "LT "
	case vm.OperationLE:
		line += "LE "
	case vm.OperationEQ:
		line += "EQ "
	case vm.OperationNE:
		line += "NE "
	}
	second := int(instruction >> vm.SecondShift & vm.ShiftMask)
	first := int(instruction >> vm.FirstShift & vm.ShiftMask)
	result := int(instruction >> vm.ResultShift & vm.ShiftMask)
	line += fmt.Sprintf("%02d %02d %02d\n", result, first, second)
	return line
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
