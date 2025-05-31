package types

import (
	"math/rand"
)

type (
	Program struct {
		Data         []float64
		Instructions []uint16
	}
)

func EmptyProgram(data, instructions int) Program {
	result := Program{
		Data:         make([]float64, data),
		Instructions: make([]uint16, instructions),
	}
	for i := range result.Data {
		result.Data[i] = rand.Float64() * 20 - 10
	}
	for i := range result.Instructions {
		result.Instructions[i] = uint16(rand.Int())
	}
	return result
}
