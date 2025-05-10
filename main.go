package main

import (
	"github.com/pehringer/mapper/internal/io"
	"github.com/pehringer/mapper/internal/ga"
)

func main() {
	mappings, err := io.ReadMappings("examples/piecewise_complex.csv")//piecewise_complex.csv")
	if err != nil {
		panic(err)
	}
	program := ga.Evolution(mappings, 0.70, 10, 40, 1024)//1048576)
	err = io.WriteProgram("solution.asm", program)
	if err != nil {
		panic(err)
	}
}
