package vm

import (
	"math"

	"github.com/pehringer/mapper/internal/types"
)

type (
	State struct {
		data [8]float64
		flag byte
	}
)

const (
	OperationMask = 0xF000
	OperationAD   = 0x0000
	OperationSB   = 0x1000
	OperationML   = 0x2000
	OperationDV   = 0x3000
	OperationPW   = 0x4000
	OperationSQ   = 0x5000
	OperationEX   = 0x6000
	OperationLG   = 0x7000
	OperationSN   = 0x8000
	OperationCS   = 0x9000
	OperationTN   = 0xA000
	OperationAB   = 0xB000
	OperationLT   = 0xC000
	OperationLE   = 0xD000
	OperationEQ   = 0xE000
	OperationNE   = 0xF000

	ResultShift    = 9
	FirstShift     = 6
	SecondShift    = 3
	PredicateShift = 0
	ShiftMask      = 0x0007
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
	predicate := int(instruction >> PredicateShift & ShiftMask)
	if 1 << predicate & s.flag == 0 {
		return
	}
	second := int(instruction >> SecondShift & ShiftMask)
	first := int(instruction >> FirstShift & ShiftMask)
	result := int(instruction >> ResultShift & ShiftMask)
	operation := instruction & OperationMask
	switch operation {
	case OperationAD:
		s.data[result] = guardEdge(s.data[first] + s.data[second])
	case OperationSB:
		s.data[result] = guardEdge(s.data[first] - s.data[second])
	case OperationML:
		s.data[result] = guardEdge(s.data[first] * s.data[second])
	case OperationDV:
		s.data[result] = guardEdge(s.data[first] / guardZero(s.data[second]))
	case OperationPW:
		s.data[result] = guardEdge(math.Pow(s.data[first], s.data[second]))
	case OperationSQ:
		s.data[result] = guardEdge(math.Sqrt(s.data[first]))
	case OperationEX:
		s.data[result] = guardEdge(math.Exp(s.data[first]))
	case OperationLG:
		s.data[result] = guardEdge(math.Log(s.data[first]))
	case OperationSN:
		s.data[result] = guardEdge(math.Sin(s.data[first]))
	case OperationCS:
		s.data[result] = guardEdge(math.Cos(s.data[first]))
	case OperationTN:
		s.data[result] = guardEdge(math.Tan(s.data[first]))
	case OperationAB:
		//s.data[result] = guardEdge(math.Abs(s.data[first]))
		if 1 << first & s.flag != 0 && 1 << second & s.flag != 0 {
			s.flag |= 1 << result
		} else {
			s.flag &^= 1 << result
		}
	case OperationLT:
		if s.data[first] < s.data[second] {
			s.flag |= 1 << result
		} else {
			s.flag &^= 1 << result
		}
	case OperationLE:
		if s.data[first] <= s.data[second] {
			s.flag |= 1 << result
		} else {
			s.flag &^= 1 << result
		}
	case OperationEQ:
		if s.data[first] == s.data[second] {
			s.flag |= 1 << result
		} else {
			s.flag &^= 1 << result
		}
	case OperationNE:
		if s.data[first] != s.data[second] {
			s.flag |= 1 << result
		} else {
			s.flag &^= 1 << result
		}
	}
	return
}

func (s *State) Run(input float64, program types.Program) float64 {
	s.data[0] = input
	s.flag = 255
	for i := range program.Instructions {
		s.execute(program.Instructions[i])
	}
	return s.data[0]
}
