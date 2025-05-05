package vm

import (
	"math"
)

type (
	Machine struct {
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

func (m *Machine) Set(accumulator float64, memory []float64) {
	m.flag = flagNone
	m.accumulator = accumulator
	if m.memory == nil {
		m.memory = make([]float64, len(memory))
	}
	copy(m.memory, memory)
}

func (m *Machine) Get() float64 {
	return m.accumulator
}

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

func (m *Machine) Execute(instruction uint16) {
	condition := instruction & Condition
	switch {
	case condition == ConditionLT && m.flag != flagN:
		return
	case condition == ConditionLE && m.flag == flagNone:
		return
	case condition == ConditionEQ && m.flag != flagZ:
		return
	case condition == ConditionNE && m.flag == flagZ:
		return
	case condition == ConditionGE && m.flag == flagN:
		return
	case condition == ConditionGT && m.flag != flagNone:
		return
	case condition == ConditionNV:
		return
	}
	operation := instruction & Operation
	switch operation {
	case OperationLD:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator = m.memory[address]
	case OperationST:
		address := int(instruction & Address) % len(m.memory)
		m.memory[address] = m.accumulator
	case OperationAD:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator += m.memory[address]
	case OperationSB:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator -= m.memory[address]
	case OperationML:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator *= m.memory[address]
	case OperationDV:
		address := int(instruction & Address) % len(m.memory)
		operand := m.memory[address]
		m.accumulator /= guardZero(operand)
	case OperationMX:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator = math.Max(m.accumulator, m.memory[address])
	case OperationMN:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator = math.Min(m.accumulator, m.memory[address])
	case OperationAB:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator = math.Abs(m.memory[address])
	case OperationPW:
		address := int(instruction & Address) % len(m.memory)
		result := math.Pow(m.accumulator, m.memory[address])
		m.accumulator = guardEdge(result)
	case OperationSQ:
		address := int(instruction & Address) % len(m.memory)
		result := math.Sqrt(m.memory[address])
		m.accumulator = guardEdge(result)
	case OperationEX:
		address := int(instruction & Address) % len(m.memory)
		result := math.Exp(m.memory[address])
		m.accumulator = guardEdge(result)
	case OperationLG:
		address := int(instruction & Address) % len(m.memory)
		result := math.Log(m.memory[address])
		m.accumulator = guardEdge(result)
	case OperationSN:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator = math.Sin(m.memory[address])
	case OperationCS:
		address := int(instruction & Address) % len(m.memory)
		m.accumulator = math.Cos(m.memory[address])
	case OperationTN:
		address := int(instruction & Address) % len(m.memory)
		result := math.Tan(m.memory[address])
		m.accumulator = guardEdge(result)
	}
	setFlag := instruction & SetFlag
	switch {
	case setFlag == SetFlagS && m.accumulator < 0:
		m.flag = flagN
	case setFlag == SetFlagS && m.accumulator == 0:
		m.flag = flagZ
	case setFlag == SetFlagS && m.accumulator > 0:
		m.flag = flagNone
	}
	return
}
