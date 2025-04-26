package main

import (
	"github.com/pehringer/fungen/internal/io"
	"github.com/pehringer/fungen/internal/ga"
)

func main() {
	tests, err := io.ReadTests("examples/GoldsteinPriceFunction.csv")
	if err != nil {
		panic(err)
	}
	ga.Evolution(tests, 1000.0, 16, 128, 4096 * 4)
}
