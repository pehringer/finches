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
	operationMAX        // 00100
	operationMIN        // 00101
	operationABD        // 00110
	operationAVG        // 00111
	operationNOP8       // 01000
	operationNOP9       // 01001
	operationNOP10      // 01010
	operationNOP11      // 01011
	operationNOP12      // 01100
	operationNOP13      // 01101
	operationNOP14      // 01110
	operationNOP15      // 01111
	operationNOP16      // 10000
	operationNOP17      // 10001
	operationNOP18      // 10010
	operationNOP19      // 10011
	operationNOP20      // 10100
	operationNOP21      // 10101
	operationNOP22      // 10110
	operationNOP23      // 10111
	operationNOP24      // 11000
	operationNOP25      // 11001
	operationNOP26      // 11010
	operationNOP27      // 11011
	operationNOP28      // 11100
	operationNOP29      // 11101
	operationNOP30      // 11110
	operationNOP31      // 11111
)

func NewMachine(registers []float32) *Machine {
	if len(registers) < 1 || len(registers) > 256 {
		return nil
	}
	result := Machine {
		flags:     flagNone,
		registers: make([]float32, len(registers)),
	}
	copy(result.registers, registers)
	return &result
}

func (m *Machine) SetRegister(destination int, value float32) {
	m.registers[destination % len(m.registers)] = value
}

func (m *Machine) GetRegister(source int) float32 {
	return m.registers[source % len(m.registers)]
}

func (m *Machine) GetRegisters() []float32 {
	result := make([]float32, len(m.registers))
	copy(result, m.registers)
	return result
}

func (m *Machine) Execute(instruction uint32) {
	condition := instruction >> 30 & 3
	operation := instruction >> 25 & 31
	setFlags := instruction >> 24 & 1 == 1
	destination := int(instruction >> 16 & 255) % len(m.registers)
	source1 := int(instruction >> 8 & 255) % len(m.registers)
	source2 := int(instruction >> 0 & 255) % len(m.registers)
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
	case operationMAX:
		result = first
		if first < second {
			result = second
		}
	case operationMIN:
		result = first
		if first > second {
			result = second
		}
	case operationABD:
		result = first - second
		if result < 0 {
			result = -result
		}
	case operationAVG:
		result = (first + second) / 2
	case operationNOP8:
		return
	case operationNOP9:
		return
	case operationNOP10:
		return
	case operationNOP11:
		return
	case operationNOP12:
		return
	case operationNOP13:
		return
	case operationNOP14:
		return
	case operationNOP15:
		return
	case operationNOP16:
		return
	case operationNOP17:
		return
	case operationNOP18:
		return
	case operationNOP19:
		return
	case operationNOP20:
		return
	case operationNOP21:
		return
	case operationNOP22:
		return
	case operationNOP23:
		return
	case operationNOP24:
		return
	case operationNOP25:
		return
	case operationNOP26:
		return
	case operationNOP27:
		return
	case operationNOP28:
		return
	case operationNOP29:
		return
	case operationNOP30:
		return
	case operationNOP31:
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
