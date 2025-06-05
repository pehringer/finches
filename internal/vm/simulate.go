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
	OpcodeAS   = 0x9000
	OpcodeCS   = 0xA000
	OpcodeAC   = 0xB000
	OpcodeMN   = 0xC000
	OpcodeMX   = 0xD000
	OpcodeLT   = 0xE000
	OpcodeGT   = 0xF000

	ResultShift = 8
	FirstShift  = 4
	SecondShift = 0
	ShiftMask   = 0x000F
)

func float(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func divide(n, d float64) float64 {
	if math.Abs(d) < 1e-9 {
		if math.Abs(n) < 1e-9 {
			return math.NaN() // zero/0
		} else if n > 0 {
			return math.Inf(1) // positive/0
		} else if n < 0 {
			return math.Inf(-1) // negative/0
		}
	}
	return n / d
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
		s[result] = divide(s[first], s[second])
	case OpcodePW:
		s[result] = math.Pow(s[first], s[second])
	case OpcodeSQ:
		s[result] = math.Sqrt(s[first])
	case OpcodeEX:
		s[result] = math.Exp(s[first])
	case OpcodeLG:
		s[result] = math.Log(s[first])
	case OpcodeSN:
		s[result] = math.Sin(s[first])
	case OpcodeAS:
		s[result] = math.Asin(s[first])
	case OpcodeCS:
		s[result] = math.Cos(s[first])
	case OpcodeAC:
		s[result] = math.Acos(s[first])
	case OpcodeMN:
		s[result] = math.Min(s[first], s[second])
	case OpcodeMX:
		s[result] = math.Max(s[first], s[second])
	case OpcodeLT:
		s[result] = float(s[first] < s[second])
	case OpcodeGT:
		s[result] = float(s[first] > s[second])
	}
	return
}

func (s *State) Run(inputs []float64, program types.Program) float64 {
	copy(s[:], program.Data)
	copy(s[:], inputs)
	for i := range program.Instructions {
		s.execute(program.Instructions[i])
	}
	return s[15]
}
