package io

import (
	"fmt"
	"os"

	"github.com/pehringer/mapper/internal/vm"
)

func parseInstruction(instruction uint16) string {
	line := ""
	opcode := instruction & vm.OpcodeMask
	switch opcode {
	case vm.OpcodeAD:
		line += "AD "
	case vm.OpcodeSB:
		line += "SB "
	case vm.OpcodeML:
		line += "ML "
	case vm.OpcodeDV:
		line += "DV "
	case vm.OpcodePW:
		line += "PW "
	case vm.OpcodeSQ:
		line += "SQ "
	case vm.OpcodeEX:
		line += "EX "
	case vm.OpcodeLG:
		line += "LG "
	case vm.OpcodeSN:
		line += "SN "
	case vm.OpcodeAS:
		line += "AS "
	case vm.OpcodeCS:
		line += "CS "
	case vm.OpcodeAC:
		line += "AC "
	case vm.OpcodeMN:
		line += "MN "
	case vm.OpcodeMX:
		line += "MX "
	case vm.OpcodeLT:
		line += "LT "
	case vm.OpcodeGT:
		line += "GT "
	}
	second := int(instruction >> vm.SecondShift & vm.ShiftMask)
	first := int(instruction >> vm.FirstShift & vm.ShiftMask)
	result := int(instruction >> vm.ResultShift & vm.ShiftMask)
	line += fmt.Sprintf("%02d %02d %02d\n", result, first, second)
	return line
}

func WriteSolution(path string, constants []float64, instructions []uint16) error {
	assembly := ""
		for i := range constants {
		assembly += fmt.Sprintf("%02d %f\n", i, constants[i])
	}
	for i := range instructions {
		assembly += parseInstruction(instructions[i])
	}
	err := os.WriteFile(path, []byte(assembly), 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}
	return nil
}
