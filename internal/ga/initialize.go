package ga

import (
	"math"

	"github.com/pehringer/mapper/internal/types"
)

func initialize(data, instructions, individuals int) []individual {
	population := make([]individual, individuals)
	for i := range population {
		population[i].fitness = math.MaxFloat64
		population[i].Program = types.EmptyProgram(data, instructions)
	}
	return population
}

