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

	opcodeMask = 0x7000
	opcodeADD  = 0x0000
	opcodeSUB  = 0x1000
	opcodeMUL  = 0x2000
	opcodeDIV  = 0x3000
	opcodeMAX  = 0x4000
	opcodeMIN  = 0x5000
	opcodeMIX  = 0x6000
	opcodeSIG  = 0x7000
)

type registers [16]float64

func setup(constants []float64) *registers {
	state := registers{}
	copy(state[:], constants)
	return &state
}

func mix(x, y, w float64) float64 {
	if w >= 1.0 {
		return y
	}
	if w <= 0.0 {
		return x
	}
	return (1.0 - w) * x + w * y
}

func sig(x, y float64) float64 {
	return 1.0 / (1.0 + math.Exp(x - y))
}

func (r *registers) execute(instruction int) {
	second := instruction >> secondShift & secondMask
	first := instruction >> firstShift & firstMask
	result := instruction >> resultShift & resultMask
	opcode := instruction & opcodeMask
	switch opcode {
	case opcodeADD:
		r[result] = r[first] + r[second]
	case opcodeSUB:
		r[result] = r[first] - r[second]
	case opcodeMUL:
		r[result] = r[first] * r[second]
	case opcodeDIV:
		r[result] = r[first] / r[second]
	case opcodeMAX:
		r[result] = math.Max(r[first], r[second])
	case opcodeMIN:
		r[result] = math.Min(r[first], r[second])
	case opcodeMIX:
		r[result] = mix(r[result], r[first], r[second])
	case opcodeSIG:
		r[result] = sig(r[first], r[second])
	}
}

func (r *registers) process(inputs []float64, instructions []int) {
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
