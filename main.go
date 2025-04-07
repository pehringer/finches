package main

import (
	"fmt"
)

type (
	machine struct {
		negative	bool
		overflow	bool
		zero		bool
		counter		int
		register	[32]uint32
		code		[]uint32
	}
	instruction struct {
		operation	int
		immediate	bool
		setFlags	bool
		condition	int
		destination	int
		source1		int
		source2		int
	}
)

const (
	BIT_11 = 0x00000800
	LSB_1  = 0x00000001
	LSB_3  = 0x00000007
	LSB_5  = 0x0000001F
	LSB_12 = 0x00000FFF
	MSB_20 = 0xFFFFF000
)

func (m *machine) fetch() uint32 {
	result := m.code[m.counter]
	m.counter = (m.counter + 1) % len(m.code)
	return result
}

func decode(operation uint32) *instruction {
	result := instruction {
		operation:	int(operation >> 27 & LSB_5),
                immediate:	0 != (operation >> 26 & LSB_1),
                setFlags:	0 != (operation >> 25 & LSB_1),
                condition:	int(operation >> 22 & LSB_3),
                destination:	int(operation >> 17 & LSB_5),
                source1:	int(operation >> 12 & LSB_5),
                source2:	int(operation & LSB_12),
	}
	if !result.immediate {
		result.source2 &= LSB_5
	}
	return &result
}

func (m *machine) execute(i *instruction) error {
	//TODO check condition before execution.
	first := m.register[i.source1]
	second := uint32(i.source2)
	if !i.immediate {
		second = m.register[i.source2]
	}
	result := uint32(0)
	switch i.operation {
	case 0x00:
		if i.immediate {
                	second = uint32(int32(second << 20) >> 20)
		}
		result = uint32(int32(first) + int32(second))
	case 0x01:
		if i.immediate {
                	second = uint32(int32(second << 20) >> 20)
		}
		result = uint32(int32(first) - int32(second))
	case 0x02:
		if i.immediate {
                	second = uint32(int32(second << 20) >> 20)
		}
		result = uint32(int32(first) * int32(second))
	case 0x03:
		if i.immediate {
                	second = uint32(int32(second << 20) >> 20)
		}
		if second == 0 {
			return fmt.Errorf("division by zero")
		}
		result = uint32(int32(first) / int32(second))
	case 0x04:
		if i.immediate {
                	second = uint32(int32(second << 20) >> 20)
		}
		if second == 0 {
			return fmt.Errorf("modulo by zero")
		}
		result = uint32(int32(first) % int32(second))
	case 0x05:
		if i.immediate {
                	second = uint32(int32(second << 20) >> 20)
		}
		result = uint32(max(int32(first), int32(second)))
	case 0x06:
		if i.immediate {
                	second = uint32(int32(second << 20) >> 20)
		}
		result = uint32(min(int32(first), int32(second)))
	case 0x07:
		result = first & second
	case 0x08:
		result = ^(first & second)
	case 0x09:
		result = first ^ second
	case 0x0A:
		result = first | second
	case 0x0B:
		result = ^(first | second)
	case 0x0C:
		result = first << second
	case 0x0D:
		result = first >> second
	case 0x0E:
	case 0x0F:
	case 0x10:
	case 0x11:
	case 0x12:
	case 0x13:
	case 0x14:
	case 0x15:
	case 0x16:
	case 0x17:
	case 0x18:
	case 0x19:
	case 0x1A:
	case 0x1B:
	case 0x1C:
	case 0x1D:
	case 0x1E:
	case 0x1F:
	}
	fmt.Println(result)
	//TODO set flags if required.
	return nil
}

func main() {
	m := machine{
		code: []uint32{0, 0, 0, 0},
	}
	err := m.execute(decode(m.fetch()))
	fmt.Println(err)
}
