package vm

import (
	"math"
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

func execute(registers *[16]float64, instruction uint16) {
	second := int(instruction >> SecondShift & ShiftMask)
	first := int(instruction >> FirstShift & ShiftMask)
	result := int(instruction >> ResultShift & ShiftMask)
	opcode := instruction & OpcodeMask
	switch opcode {
	case OpcodeAD:
		registers[result] = registers[first] + registers[second]
	case OpcodeSB:
		registers[result] = registers[first] - registers[second]
	case OpcodeML:
		registers[result] = registers[first] * registers[second]
	case OpcodeDV:
		registers[result] = divide(registers[first], registers[second])
	case OpcodePW:
		registers[result] = math.Pow(registers[first], registers[second])
	case OpcodeSQ:
		registers[result] = math.Sqrt(registers[first])
	case OpcodeEX:
		registers[result] = math.Exp(registers[first])
	case OpcodeLG:
		registers[result] = math.Log(registers[first])
	case OpcodeSN:
		registers[result] = math.Sin(registers[first])
	case OpcodeAS:
		registers[result] = math.Asin(registers[first])
	case OpcodeCS:
		registers[result] = math.Cos(registers[first])
	case OpcodeAC:
		registers[result] = math.Acos(registers[first])
	case OpcodeMN:
		registers[result] = math.Min(registers[first], registers[second])
	case OpcodeMX:
		registers[result] = math.Max(registers[first], registers[second])
	case OpcodeLT:
		registers[result] = float(registers[first] < registers[second])
	case OpcodeGT:
		registers[result] = float(registers[first] > registers[second])
	}
	return
}

func Run(inputs, constants []float64, instructions []uint16) float64 {
	registers := [16]float64{}
	copy(registers[:], constants)
	copy(registers[:], inputs)
	for i := range instructions {
		execute(&registers, instructions[i])
	}
	return registers[15]
}
