package main

import (
	"fmt"

	"github.com/pehringer/fungen/internal/io"
	"github.com/pehringer/fungen/internal/ga"
)

func main() {
	tests, err := io.ReadTests("examples/unflattenable.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(tests)
	solution := ga.Evolution(tests, 20.0, 256, 4096 * 4)
	err = io.WriteProgram("solution.asm", solution)
	if err != nil {
		panic(err)
	}
}
