package main

import (
	"math"
)

const (
	opcodeMask  = 0xF000
	opcodeAD    = 0x0000
	opcodeSB    = 0x1000
	opcodeML    = 0x2000
	opcodeDV    = 0x3000
	opcodePW    = 0x4000
	opcodeSQ    = 0x5000
	opcodeEX    = 0x6000
	opcodeLG    = 0x7000
	opcodeSN    = 0x8000
	opcodeAS    = 0x9000
	opcodeCS    = 0xA000
	opcodeAC    = 0xB000
	opcodeMN    = 0xC000
	opcodeMX    = 0xD000
	opcodeLT    = 0xE000
	opcodeGT    = 0xF000
	resultShift = 8
	firstShift  = 4
	secondShift = 0
	shiftMask   = 0x000F
)

func castFloat(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func safeDivision(n, d float64) float64 {
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

func executeInstruction(registers *[16]float64, instruction uint16) {
	second := int(instruction >> secondShift & shiftMask)
	first := int(instruction >> firstShift & shiftMask)
	result := int(instruction >> resultShift & shiftMask)
	opcode := instruction & opcodeMask
	switch opcode {
	case opcodeAD:
		registers[result] = registers[first] + registers[second]
	case opcodeSB:
		registers[result] = registers[first] - registers[second]
	case opcodeML:
		registers[result] = registers[first] * registers[second]
	case opcodeDV:
		registers[result] = safeDivision(registers[first], registers[second])
	case opcodePW:
		registers[result] = math.Pow(registers[first], registers[second])
	case opcodeSQ:
		registers[result] = math.Sqrt(registers[first])
	case opcodeEX:
		registers[result] = math.Exp(registers[first])
	case opcodeLG:
		registers[result] = math.Log(registers[first])
	case opcodeSN:
		registers[result] = math.Sin(registers[first])
	case opcodeAS:
		registers[result] = math.Asin(registers[first])
	case opcodeCS:
		registers[result] = math.Cos(registers[first])
	case opcodeAC:
		registers[result] = math.Acos(registers[first])
	case opcodeMN:
		registers[result] = math.Min(registers[first], registers[second])
	case opcodeMX:
		registers[result] = math.Max(registers[first], registers[second])
	case opcodeLT:
		registers[result] = castFloat(registers[first] < registers[second])
	case opcodeGT:
		registers[result] = castFloat(registers[first] > registers[second])
	}
}

func simulateProgram(inputs, constants []float64, instructions []uint16) float64 {
	registers := [16]float64{}
	copy(registers[:], constants)
	copy(registers[:], inputs)
	for i := range instructions {
		executeInstruction(&registers, instructions[i])
	}
	return registers[15]
}
