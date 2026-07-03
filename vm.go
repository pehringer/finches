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
	opcodeMN    = 0x4000
	opcodeMX    = 0x5000
	opcodePW    = 0x6000
	opcodeSQ    = 0x7000
	opcodeEX    = 0x8000
	opcodeLG    = 0x9000
	opcodeSN    = 0xA000
	opcodeAS    = 0xB000
	opcodeCS    = 0xC000
	opcodeAC    = 0xD000
	opcodeSG    = 0xE000
	opcodeLT    = 0xF000
	resultShift = 8
	firstShift  = 4
	secondShift = 0
	shiftMask   = 0x000F
)

type registers [16]float64

func setup(constants []float64) *registers {
	state := registers{}
	copy(state[:], constants)
	return &state
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func logit(x float64) float64 {
	return math.Log(x / (1 - x))
}

func (r *registers) execute(instruction uint16) {
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
		r[result] = r[first] / r[second]
	case opcodeMN:
		r[result] = math.Min(r[first], r[second])
	case opcodeMX:
		r[result] = math.Max(r[first], r[second])
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
	case opcodeSG:
		r[result] = sigmoid(r[first])
	case opcodeLT:
		r[result] = logit(r[first])
	}
}

func (r *registers) process(inputs []float64, instructions []uint16) {
	copy(r[:], inputs)
	for i := range instructions {
		r.execute(instructions[i])
	}
}

func (r *registers) reset(constants []float64) float64 {
	output := r[15]
	copy(r[:], constants)
	return output
}
