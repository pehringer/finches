package vm

type (
	Machine struct {
		state uint16
		accumulator float32
		registers [64]float32
	}
)

const (
	Registers = 64

	State = 10
	Opcode  = 0x03C0
	Operand = 0x003F

	OpcodeLD    = 0x0000
	OpcodeST    = 0x0040
	OpcodeAD    = 0x0080
	OpcodeSB    = 0x00C0
	OpcodeML    = 0x0100
	OpcodeDV    = 0x0140
	OpcodeLT    = 0x0180
	OpcodeGT    = 0x01C0
	OpcodeEQ    = 0x0200
	OpcodeNE    = 0x0240
	OpcodeNOP10 = 0x0280
	OpcodeNOP11 = 0x02C0
	OpcodeNOP12 = 0x0300
	OpcodeNOP13 = 0x0340
	OpcodeNOP14 = 0x0380
	OpcodeNOP15 = 0x03C0
)

func (m *Machine) Reset(registers []float32) {
	m.state = 0
	m.accumulator = 0
	copy(m.registers[:], registers)
}

func (m *Machine) Set(register int, value float32) {
		m.registers[register % Registers] = value
}

func (m *Machine) Get(register int) float32 {
		return m.registers[register % Registers]
}

func (m *Machine) Execute(instruction uint16) {
	if m.state != instruction >> State {
		return
	}
	switch instruction & Opcode {
	case OpcodeLD:
		m.accumulator = m.registers[int(instruction & Operand)]
	case OpcodeST:
		m.registers[int(instruction & Operand)] = m.accumulator
	case OpcodeAD:
		m.accumulator += m.registers[int(instruction & Operand)]
	case OpcodeSB:
		m.accumulator -= m.registers[int(instruction & Operand)]
	case OpcodeML:
		m.accumulator *= m.registers[int(instruction & Operand)]
	case OpcodeDV:
		protected := m.registers[int(instruction & Operand)]
		if protected == 0 {
			protected = 1
		}
		m.accumulator /= protected
	case OpcodeLT:
		if m.accumulator < 0 {
				m.state = instruction & Operand
		}
	case OpcodeGT:
		if m.accumulator > 0 {
				m.state = instruction & Operand
		}
	case OpcodeEQ:
		if m.accumulator == 0 {
				m.state = instruction & Operand
		}
	case OpcodeNE:
		if m.accumulator != 0 {
				m.state = instruction & Operand
		}
	case OpcodeNOP10:
	case OpcodeNOP11:
	case OpcodeNOP12:
	case OpcodeNOP13:
	case OpcodeNOP14:
	case OpcodeNOP15:
	}
	return
}
