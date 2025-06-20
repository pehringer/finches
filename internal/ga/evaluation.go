package ga

import (
	"math"

	"github.com/pehringer/mapper/internal/vm"
)


func evaluation(candidate *individual, penalty float64, inputs [][]float64, outputs []float64) {
	candidate.fitness = 0
        for i := range min(len(inputs), len(outputs)) {
        	output := vm.Run(inputs[i], candidate.constants, candidate.instructions)
        	delta := math.Abs(output - outputs[i])
		if math.IsNaN(delta) || math.IsInf(delta, 0) {
			candidate.fitness += penalty
		} else {
			candidate.fitness += delta
		}
	}
}

func finalist(population []individual) *individual {
	minimum := &population[0]
	for i := range population {
		if population[i].fitness < minimum.fitness {
			minimum = &population[i]
		}
	}
	return minimum
}
