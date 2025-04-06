package main

import (
	"fmt"
)

type (
	MachineState struct {
		flagN          bool
		flagO          bool
		flagZ          bool
                machineCode    []uint32
		programCounter int
		registers      [32]int32
	}
)

const (
	LSB_1  = 0x1
	LSB_3  = 0x7
	LSB_5  = 0x1F
	LSB_12 = 0xFFF
)

func (m *MachineState) fetch() (instruction uint32) {
	instruction = m.machineCode[m.programCounter]
	m.programCounter = (m.programCounter + 1) % len(m.machineCode)
	return
}

func decode(instruction uint32) (operation, typ, set, condition, result, firstOperand, secondOperand int) {
	operation = int(instruction >> 27 & LSB_5)
	typ = int(instruction >> 26 & LSB_1)
	set = int(instruction >> 25 & LSB_1)
	condition = int(instruction >> 22 & LSB_3)
	result = int(instruction >> 17 & LSB_5)
	firstOperand = int(instruction >> 12 & LSB_5)
	secondOperand = int(instruction >> 0 & LSB_12)
	if typ == 0 {
		secondOperand &= LSB_5
	}
	return
}

func main() {
	fmt.Println("Testing...")
}
