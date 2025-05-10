package vm

import (
	"math"

	"github.com/pehringer/mapper/internal/types"
)

type (
	State struct {
		flag byte
		accumulator float64
		memory []float64
	}
)

const (
	flagNone = 0x0
	flagZ    = 0x1
	flagN    = 0x2

	Condition   = 0xE000
	ConditionLT = 0x2000
	ConditionLE = 0x4000
	ConditionEQ = 0x6000
	ConditionNE = 0x8000
	ConditionGE = 0xA000
	ConditionGT = 0xC000
	ConditionNV = 0xE000

	Operation   = 0x1E00
	OperationLD = 0x0000
	OperationST = 0x0200
	OperationAD = 0x0400
	OperationSB = 0x0600
	OperationML = 0x0800
	OperationDV = 0x0A00
	OperationMX = 0x0C00
	OperationMN = 0x0E00
	OperationAB = 0x1000
	OperationPW = 0x1200
	OperationSQ = 0x1400
	OperationEX = 0x1600
	OperationLG = 0x1800
	OperationSN = 0x1A00
	OperationCS = 0x1C00
	OperationTN = 0x1E00

	SetFlag  = 0x0100
	SetFlagS = 0x0100

	Address = 0x00FF
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
	condition := instruction & Condition
	switch {
	case condition == ConditionLT && s.flag != flagN:
		return
	case condition == ConditionLE && s.flag == flagNone:
		return
	case condition == ConditionEQ && s.flag != flagZ:
		return
	case condition == ConditionNE && s.flag == flagZ:
		return
	case condition == ConditionGE && s.flag == flagN:
		return
	case condition == ConditionGT && s.flag != flagNone:
		return
	case condition == ConditionNV:
		return
	}
	operation := instruction & Operation
	switch operation {
	case OperationLD:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator = s.memory[address]
	case OperationST:
		address := int(instruction & Address) % len(s.memory)
		s.memory[address] = s.accumulator
	case OperationAD:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator += s.memory[address]
	case OperationSB:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator -= s.memory[address]
	case OperationML:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator *= s.memory[address]
	case OperationDV:
		address := int(instruction & Address) % len(s.memory)
		operand := s.memory[address]
		s.accumulator /= guardZero(operand)
	case OperationMX:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator = math.Max(s.accumulator, s.memory[address])
	case OperationMN:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator = math.Min(s.accumulator, s.memory[address])
	case OperationAB:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator = math.Abs(s.memory[address])
	case OperationPW:
		address := int(instruction & Address) % len(s.memory)
		result := math.Pow(s.accumulator, s.memory[address])
		s.accumulator = guardEdge(result)
	case OperationSQ:
		address := int(instruction & Address) % len(s.memory)
		result := math.Sqrt(s.memory[address])
		s.accumulator = guardEdge(result)
	case OperationEX:
		address := int(instruction & Address) % len(s.memory)
		result := math.Exp(s.memory[address])
		s.accumulator = guardEdge(result)
	case OperationLG:
		address := int(instruction & Address) % len(s.memory)
		result := math.Log(s.memory[address])
		s.accumulator = guardEdge(result)
	case OperationSN:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator = math.Sin(s.memory[address])
	case OperationCS:
		address := int(instruction & Address) % len(s.memory)
		s.accumulator = math.Cos(s.memory[address])
	case OperationTN:
		address := int(instruction & Address) % len(s.memory)
		result := math.Tan(s.memory[address])
		s.accumulator = guardEdge(result)
	}
	setFlag := instruction & SetFlag
	switch {
	case setFlag == SetFlagS && s.accumulator < 0:
		s.flag = flagN
	case setFlag == SetFlagS && s.accumulator == 0:
		s.flag = flagZ
	case setFlag == SetFlagS && s.accumulator > 0:
		s.flag = flagNone
	}
	return
}

func (s *State) Run(input float64, program types.Program) float64 {
	s.flag = flagNone
	s.accumulator = input
	if len(s.memory) != len(program.Data) {
		s.memory = make([]float64, len(program.Data))
	}
	copy(s.memory, program.Data)
	for i := range program.Instructions {
		s.execute(program.Instructions[i])
	}
	return s.accumulator
}
