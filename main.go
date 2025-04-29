package main

import (
	"github.com/pehringer/fungen/internal/io"
	"github.com/pehringer/fungen/internal/ga"
)

func main() {
	// f(x,y)= 0.26(x^2 + y^2) − 0.48xy
	tests, err := io.ReadTests("examples/MatyasFunction.csv")
	if err != nil {
		panic(err)
	}
	solution := ga.Evolution(tests, 1.0, 8, 64, 4096)
	err = io.WriteProgram("solution.asm", solution)
	if err != nil {
		panic(err)
	}
}
