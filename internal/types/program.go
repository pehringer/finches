package types

type (
	Program struct {
		Data         []float64
		Instructions []uint16
	}
)

func EmptyProgram(data, instructions int) Program {
	return Program{
		Data:         make([]float64, data),
		Instructions: make([]uint16, instructions),
	}
}
