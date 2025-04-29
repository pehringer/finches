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

	SetFlags   = 0x80000000
	SetFlagsNo = 0x00000000
	SetFlagsS  = 0x80000000

	Condition       = 0x70000000
	ConditionAlways = 0x00000000
	ConditionLT     = 0x10000000
	ConditionLE     = 0x20000000
	ConditionEQ     = 0x30000000
	ConditionNE     = 0x40000000
	ConditionGE     = 0x50000000
	ConditionGT     = 0x60000000
	ConditionNOP    = 0x70000000

	Operation      = 0x0F000000
	OperationADD   = 0x00000000
	OperationSUB   = 0x01000000
	OperationMUL   = 0x02000000
	OperationDIV   = 0x03000000
	OperationMOV4  = 0x04000000
	OperationMOV5  = 0x05000000
	OperationMOV6  = 0x06000000
	OperationMOV7  = 0x07000000
	OperationMOV8  = 0x08000000
	OperationMOV9  = 0x09000000
	OperationMOV10 = 0x0A000000
	OperationMOV11 = 0x0B000000
	OperationMOV12 = 0x0C000000
	OperationMOV13 = 0x0D000000
	OperationMOV14 = 0x0E000000
	OperationMOV15 = 0x0F000000

	Destination = 0x00FF0000
	Source1     = 0x0000FF00
	Source2     = 0x000000FF
)

func SetState(registers []float32) *Machine {
	if len(registers) < 1 || len(registers) > 256 {
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
	switch {
	case condition == ConditionLT && m.flags != flagN:
		return
	case condition == ConditionLE && m.flags == flagNone:
		return
	case condition == ConditionEQ && m.flags != flagZ:
		return
	case condition == ConditionNE && m.flags == flagZ:
		return
	case condition == ConditionGE && m.flags == flagN:
		return
	case condition == ConditionGT && m.flags != flagNone:
		return
	case condition == ConditionNOP:
		return
	}
	destination := int((instruction & Destination) >> 16) % len(m.registers)
	result := m.registers[destination]
	source1 := int((instruction & Source1) >> 8) % len(m.registers)
	first := m.registers[source1]
	source2 := int((instruction & Source2) >> 0) % len(m.registers)
	second := m.registers[source2]
	operation := instruction & Operation
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
	case OperationMOV4:
		result = first
	case OperationMOV5:
		result = first
	case OperationMOV6:
		result = first
	case OperationMOV7:
		result = first
	case OperationMOV8:
		result = first
	case OperationMOV9:
		result = first
	case OperationMOV10:
		result = first
	case OperationMOV11:
		result = first
	case OperationMOV12:
		result = first
	case OperationMOV13:
		result = first
	case OperationMOV14:
		result = first
	case OperationMOV15:
		result = first
	}
	setFlags := instruction & SetFlags == SetFlagsS
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
