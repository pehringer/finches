package ga

import (
	"math"

	"github.com/pehringer/mapper/internal/vm"
	"github.com/pehringer/mapper/internal/types"
)


func evaluateFitness(mappings []types.Mapping, candidate *individual) {
        candidate.fitness = 0
        simulation := vm.State{}
        for i := range mappings {
                output := simulation.Run(mappings[i].Input, candidate.Program)
                candidate.fitness += math.Abs(output - mappings[i].Output)
        }
}
