package main

import (
	"math"
)

const (
	secondShift = 0
	secondMask  = 0x000F

	firstShift  = 4
	firstMask   = 0x000F

	resultShift = 8
	resultMask  = 0x000F

	opcodeMask  = 0xF000
	opcodeAD    = 0x0000
	opcodeSB    = 0x1000
	opcodeML    = 0x2000
	opcodeDV    = 0x3000
	opcodeMN    = 0x4000
	opcodeMX    = 0x5000
	opcodePW    = 0x6000
	opcodeLG    = 0x7000
	opcodeSN    = 0x8000
	opcodeAS    = 0x9000
	opcodeCS    = 0xA000
	opcodeAC    = 0xB000
	opcodeTN    = 0xC000
	opcodeAT    = 0xD000
	opcodeSG    = 0xE000
	opcodeLT    = 0xF000
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
	second := int(instruction >> secondShift & secondMask)
	first := int(instruction >> firstShift & firstMask)
	result := int(instruction >> resultShift & resultMask)
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
	case opcodeLG:
		r[result] = math.Log(r[first]) / math.Log(r[second])
	case opcodeSN:
		r[result] = math.Sin(r[first])
	case opcodeAS:
		r[result] = math.Asin(r[first])
	case opcodeCS:
		r[result] = math.Cos(r[first])
	case opcodeAC:
		r[result] = math.Acos(r[first])
	case opcodeTN:
		r[result] = math.Tan(r[first])
	case opcodeAT:
		r[result] = math.Atan(r[first])
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
