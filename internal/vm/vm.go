package vm

type (
	Machine struct {
		flags     byte
		registers []float32
	}
)

const (
	flagNone = 0x00
	flagZ    = 0x01
	flagN    = 0x02

	Condition       = 0x00C00000
	ConditionAlways = 0x00000000
	ConditionLT     = 0x00400000
	ConditionGT     = 0x00800000
	ConditionEQ     = 0x00C00000

	Operation     = 0x00380000
	OperationADD  = 0x00000000
	OperationSUB  = 0x00080000
	OperationMUL  = 0x00100000
	OperationDIV  = 0x00180000
	OperationNOP4 = 0x00200000
	OperationNOP5 = 0x00280000
	OperationNOP6 = 0x00300000
	OperationNOP7 = 0x00380000

	SetFlags   = 0x00040000
	SetFlagsNo = 0x00000000
	SetFlagsS  = 0x00040000

	Destination = 0x0003F000
	Source1     = 0x00000FC0
	Source2     = 0x0000003F
)

func SetState(registers []float32) *Machine {
	if len(registers) < 1 || len(registers) > 64 {
		return nil
	}
	m := &Machine{
		registers: make([]float32, len(registers)),
	}
	m.flags = flagNone
	copy(m.registers, registers)
	return m
}

func (m *Machine) ResetState(registers []float32) {
	m.flags = flagNone
	copy(m.registers, registers)
}

func (m *Machine) SetRegister(destination int, value float32) {
	m.registers[destination%len(m.registers)] = value
}

func (m *Machine) GetRegister(source int) float32 {
	return m.registers[source%len(m.registers)]
}

func (m *Machine) Execute(instruction uint32) {
	condition := instruction & Condition
	operation := instruction & Operation
	setFlags := instruction & SetFlags == SetFlagsS
	destination := int((instruction & Destination)>>12) % len(m.registers)
	source1 := int((instruction & Source1)>>6) % len(m.registers)
	source2 := int((instruction & Source2)>>0) % len(m.registers)
	switch {
	case condition == ConditionLT && m.flags != flagN:
		return
	case condition == ConditionGT && m.flags != flagNone:
		return
	case condition == ConditionEQ && m.flags != flagZ:
		return
	}
	result := m.registers[destination]
	first := m.registers[source1]
	second := m.registers[source2]
	switch operation {
	case OperationADD:
		result = first + second
	case OperationSUB:
		result = first - second
	case OperationMUL:
		result = first * second
	case OperationDIV:
		if second == 0 {
			second = 1
		}
		result = first / second
	case OperationNOP4:
		return
	case OperationNOP5:
		return
	case OperationNOP6:
		return
	case OperationNOP7:
		return
	}
	switch {
	case setFlags && result == 0:
		m.flags = flagZ
	case setFlags && result < 0:
		m.flags = flagN
	case setFlags && result > 0:
		m.flags = flagNone
	}
	m.registers[destination] = result
	return
}
