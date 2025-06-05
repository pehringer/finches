package ga

import (
	"math"

	"github.com/pehringer/mapper/internal/vm"
	"github.com/pehringer/mapper/internal/types"
)


func evaluateFitness(mappings []types.Mapping, candidate *individual, penalty float64) {
	candidate.fitness = 0
        simulation := vm.State{}
        for i := range mappings {
        	output := simulation.Run(mappings[i].Inputs, candidate.Program)
        	delta := math.Abs(output - mappings[i].Output)
		if math.IsNaN(delta) || math.IsInf(delta, 0) {
			candidate.fitness += penalty
		} else {
			candidate.fitness += delta
		}
	}
}
