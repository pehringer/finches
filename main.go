package main

import (
	"fmt"

	"github.com/pehringer/fungen/internal/io"
	"github.com/pehringer/fungen/internal/ga"
)

func main() {
	tests, err := io.ReadTests("examples/piecewise.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(tests)
	solution := ga.Evolution(tests, 0.1, 4, 32, 4096 * 4)
	err = io.WriteProgram("solution.asm", solution)
	if err != nil {
		panic(err)
	}
}
