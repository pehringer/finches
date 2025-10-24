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

type registers [16]float64

func setupRegisters(constants []float64) *registers {
	state := registers{}
	copy(state[:], constants)
	return &state
}

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

func (r *registers) executeInstruction(instruction uint16) {
	second := int(instruction >> secondShift & shiftMask)
	first := int(instruction >> firstShift & shiftMask)
	result := int(instruction >> resultShift & shiftMask)
	opcode := instruction & opcodeMask
	switch opcode {
	case opcodeAD:
		r[result] = r[first] + r[second]
	case opcodeSB:
		r[result] = r[first] - r[second]
	case opcodeML:
		r[result] = r[first] * r[second]
	case opcodeDV:
		r[result] = safeDivision(r[first], r[second])
	case opcodePW:
		r[result] = math.Pow(r[first], r[second])
	case opcodeSQ:
		r[result] = math.Sqrt(r[first])
	case opcodeEX:
		r[result] = math.Exp(r[first])
	case opcodeLG:
		r[result] = math.Log(r[first])
	case opcodeSN:
		r[result] = math.Sin(r[first])
	case opcodeAS:
		r[result] = math.Asin(r[first])
	case opcodeCS:
		r[result] = math.Cos(r[first])
	case opcodeAC:
		r[result] = math.Acos(r[first])
	case opcodeMN:
		r[result] = math.Min(r[first], r[second])
	case opcodeMX:
		r[result] = math.Max(r[first], r[second])
	case opcodeLT:
		r[result] = castFloat(r[first] < r[second])
	case opcodeGT:
		r[result] = castFloat(r[first] > r[second])
	}
}

func (r *registers) executeInstructions(inputs []float64, instructions []uint16) {
	copy(r[:], inputs)
	for i := range instructions {
		r.executeInstruction(instructions[i])
	}
}

func (r *registers) resetRegisters(constants []float64) float64 {
	output := r[15]
	copy(r[:], constants)
	return output
}
