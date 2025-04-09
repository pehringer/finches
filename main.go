package main

import (
	"fmt"
)

type machine struct {
	stack []float32
	flags byte
}

func setInputs(inputs []float32) *machine {
	return nil
}

func (m *machine) execute(instruction [5]byte) error {
	return nil
}

func (m *machine) getOutputs() []float32 {
	return nil
}

func main() {
	fmt.Println("fungen")
}
