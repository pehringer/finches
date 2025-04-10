package fungen

import (
	"errors"
	"unsafe"
)

type Machine struct {
	stack []float32
	flags byte // bit 0 = Z, bit 1 = N
}

var (
	ErrorStackUnderflow   = errors.New("stack underflow")
	ErrorDivideByZero     = errors.New("divide by zero")
	ErrorIndexOutOfBounds = errors.New("index out of bounds")
)

func SetInputs(inputs []float32) *Machine {
	stack := make([]float32, len(inputs))
	copy(stack, inputs)
	return &Machine {
		stack: stack,
		flags: 0x00,
	}
}

func asInt32(b []byte) int32 {
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

func asFloat32(i int32) float32 {
	return *(*float32)(unsafe.Pointer(&i))
}

func (m *Machine) Execute(instruction [5]byte) error {
	switch instruction[0] & 0xE0 {
	case 0x00:
	case 0x20: // LT
		if m.flags != 0x02 {
			return nil
		}
	case 0x40: // LE
		if m.flags == 0x00 {
			return nil
		}
	case 0x60: // EQ
		if m.flags != 0x01 {
			return nil
		}
	case 0x80: // NE
		if m.flags == 0x01 {
			return nil
		}
	case 0xA0: // GE
		if m.flags == 0x02 {
			return nil
		}
	case 0xC0: // GT
		if m.flags != 0x00 {
			return nil
		}
	case 0xE0: // NV
		return nil
	}
	n := len(m.stack)
	switch instruction[0] & 0x1E {
	case 0x00: // ADD
		if n < 2 {
			return ErrorStackUnderflow
		}
		m.stack[n-2] += m.stack[n-1]
		m.stack = m.stack[:n-1]
	case 0x02: // ADDI
		if n < 1 {
			return ErrorStackUnderflow
		}
		v := asInt32(instruction[1:5])
		m.stack[n-1] += asFloat32(v)
	case 0x04: // SUB
		if n < 2 {
			return ErrorStackUnderflow
		}
		m.stack[n-2] -= m.stack[n-1]
		m.stack = m.stack[:n-1]
	case 0x06: // SUBI
		if n < 1 {
			return ErrorStackUnderflow
		}
		v := asInt32(instruction[1:5])
		m.stack[n-1] -= asFloat32(v)
	case 0x08: // MUL
		if n < 2 {
			return ErrorStackUnderflow
		}
		m.stack[n-2] *= m.stack[n-1]
		m.stack = m.stack[:n-1]
	case 0x0A: // MULI
		if n < 1 {
			return ErrorStackUnderflow
		}
		v := asInt32(instruction[1:5])
		m.stack[n-1] *= asFloat32(v)
	case 0x0C: // DIV
		if n < 2 {
			return ErrorStackUnderflow
		}
		v := m.stack[n-1]
		if v == 0 {
			return ErrorDivideByZero
		}
		m.stack[n-2] /= v
		m.stack = m.stack[:n-1]
	case 0x0E: // DIVI
		if n < 1 {
			return ErrorStackUnderflow
		}
		v := asInt32(instruction[1:5])
		if v == 0 {
			return ErrorDivideByZero
		}
		m.stack[n-1] /= asFloat32(v)
	case 0x10: // MAX
		if n < 2 {
			return ErrorStackUnderflow
		}
		m.stack[n-2] = max(m.stack[n-2], m.stack[n-1])
		m.stack = m.stack[:n-1]
	case 0x12: // MAXI
		if n < 1 {
			return ErrorStackUnderflow
		}
		v := asInt32(instruction[1:5])
		m.stack[n-1] = max(m.stack[n-1], asFloat32(v))
	case 0x14: // MIN
		if n < 2 {
			return ErrorStackUnderflow
		}
		m.stack[n-2] = min(m.stack[n-2], m.stack[n-1])
		m.stack = m.stack[:n-1]
	case 0x16: // MINI
		if n < 1 {
			return ErrorStackUnderflow
		}
		v := asInt32(instruction[1:5])
		m.stack[n-1] = min(m.stack[n-1], asFloat32(v))
	case 0x18: // POP
		if n < 1 {
			return ErrorStackUnderflow
		}
		m.stack = m.stack[:n-1]
	case 0x1A: // PUSH
		v := asInt32(instruction[1:5])
		m.stack = append(m.stack, asFloat32(v))
	case 0x1C: // SWP
		if n < 2 {
			return ErrorStackUnderflow
		}
		v := m.stack[n-2]
		m.stack[n-2] = m.stack[n-1]
		m.stack[n-1] = v
	case 0x1E: //PICK
		v := int(asInt32(instruction[1:5]))
		if v < 1 || v > n {
			return ErrorIndexOutOfBounds
		}
		m.stack = append(m.stack, m.stack[n-v])
	}
	n = len(m.stack)
	switch instruction[0] & 0x01 {
	case 0x00:
	case 0x01:
		if n < 1 {
			return ErrorStackUnderflow
		}
		v := m.stack[n-1]
		switch {
		case v == 0:
			m.flags = 0x01
		case v < 0:
			m.flags = 0x02
		case v > 0:
			m.flags = 0x00
		}
	}
	return nil
}

func (m *Machine) GetOutputs() []float32 {
	outputs := make([]float32, len(m.stack))
	copy(outputs, m.stack)
	return outputs
}

