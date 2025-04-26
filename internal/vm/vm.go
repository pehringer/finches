package vm

type (
	Machine struct {
		flags     byte
		registers []float32
	}
)

const (
	flagNone  = iota // 00
	flagZ            // 01
	flagN            // 10
)

const (
	conditionAlways = iota // 00
	conditionLT            // 01
	conditionGT            // 10
	conditionEQ            // 11
)

const (
	operationADD = iota // 00000
	operationSUB        // 00001
	operationMUL        // 00010
	operationDIV        // 00011
	operationNOP4       // 00100
	operationNOP5       // 00101
	operationNOP6       // 00110
	operationNOP7       // 00111
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
	condition := instruction >> 22 & 3
	operation := instruction >> 19 & 7
	setFlags := instruction >> 18 & 1 == 1
	destination := int(instruction >> 12 & 63) % len(m.registers)
	source1 := int(instruction >> 6 & 63) % len(m.registers)
	source2 := int(instruction >> 0 & 63) % len(m.registers)
	switch {
	case condition == conditionLT && m.flags != flagN:
		return
	case condition == conditionGT && m.flags != flagNone:
		return
	case condition == conditionEQ && m.flags != flagZ:
		return
	}
	result := m.registers[destination]
	first := m.registers[source1]
	second := m.registers[source2]
	switch operation {
	case operationADD:
		result = first + second
	case operationSUB:
		result = first - second
	case operationMUL:
		result = first * second
	case operationDIV:
		if second == 0 {
			second = 1
		}
		result = first / second
	case operationNOP4:
		return
	case operationNOP5:
		return
	case operationNOP6:
		return
	case operationNOP7:
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
