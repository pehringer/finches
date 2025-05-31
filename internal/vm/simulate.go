package vm

import (
	"math"

	"github.com/pehringer/mapper/internal/types"
)

type (
	State [16]float64
)

const (
	OpcodeMask = 0xF000
	OpcodeAD   = 0x0000
	OpcodeSB   = 0x1000
	OpcodeML   = 0x2000
	OpcodeDV   = 0x3000
	OpcodePW   = 0x4000
	OpcodeSQ   = 0x5000
	OpcodeEX   = 0x6000
	OpcodeLG   = 0x7000
	OpcodeSN   = 0x8000
	OpcodeCS   = 0x9000
	OpcodeMN   = 0xA000
	OpcodeMX   = 0xB000
	OpcodeLT   = 0xC000
	OpcodeGT   = 0xD000
	OpcodeN0   = 0xE000
	OpcodeN1   = 0xF000

	ResultShift = 8
	FirstShift  = 4
	SecondShift = 0
	ShiftMask   = 0x000F
)

func guardZero(value float64) float64 {
	if value == 0 {
		return 1
	}
	return value
}

func guardEdge(value float64) float64 {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0
	}
	return value
}

func (s *State) execute(instruction uint16) {
	second := int(instruction >> SecondShift & ShiftMask)
	first := int(instruction >> FirstShift & ShiftMask)
	result := int(instruction >> ResultShift & ShiftMask)
	opcode := instruction & OpcodeMask
	switch opcode {
	case OpcodeAD:
		s[result] = s[first] + s[second]
	case OpcodeSB:
		s[result] = s[first] - s[second]
	case OpcodeML:
		s[result] = s[first] * s[second]
	case OpcodeDV:
		s[result] = s[first] / guardZero(s[second])
	case OpcodePW:
		s[result] = guardEdge(math.Pow(s[first], s[second]))
	case OpcodeSQ:
		s[result] = guardEdge(math.Sqrt(s[first]))
	case OpcodeEX:
		s[result] = guardEdge(math.Exp(s[first]))
	case OpcodeLG:
		s[result] = guardEdge(math.Log(s[first]))
	case OpcodeSN:
		s[result] = math.Sin(s[first])
	case OpcodeCS:
		s[result] = math.Cos(s[first])
	case OpcodeMN:
		s[result] = math.Min(s[first], s[second])
	case OpcodeMX:
		s[result] = math.Max(s[first], s[second])
	case OpcodeLT:
		if s[first] < s[second] {
			s[result] = 1
		} else {
			s[result] = 0
		}
	case OpcodeGT:
		if s[first] > s[second] {
			s[result] = 1
		} else {
			s[result] = 0
		}
	case OpcodeN0:
	case OpcodeN1:
	}
	return
}

func (s *State) Run(input float64, program types.Program) float64 {
	copy(s[:], program.Data)
	s[0] = input
	for i := range program.Instructions {
		s.execute(program.Instructions[i])
	}
	return s[1]
}
