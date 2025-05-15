package ga

import (
	"fmt"
	"math/rand"
	"github.com/pehringer/mapper/internal/io"
	"github.com/pehringer/mapper/internal/vm"
	"github.com/pehringer/mapper/internal/types"
)

type (
	individual struct {
		fitness float64
		types.Program
	}
)

func termination(mappings []types.Mapping, accuracy float64) float64 {
	target := 0.0
	for i := range mappings {
		target += mappings[i].AbsoluteOutput()
	}
	return target - target * accuracy
}

func Evolve(mappings []types.Mapping, accuracy float64, data, instructions, individuals int) types.Program {
	target := termination(mappings, accuracy)
	population := initialize(data, instructions, individuals)
	solution := &population[0]
	io.PrintStarting()
	for solution.fitness > target {
		for range len(population) {
			parent1, parent2 := selectNeighbors(10, population)
			offspring := replaceDuel(parent1, parent2)
			crossoverSinglePoint(parent1, parent2, offspring)
			percent := rand.Float32()
			switch {
			case percent < 0.35:
				mutateBitFlips(1, offspring)
			case percent < 0.70:
				mutatePerturbation(-0.001, +0.001, offspring)
			case percent < 0.80:
				mutateBitFlips(3, offspring)
			case percent < 0.90:
				mutatePerturbation(-0.1, +0.1, offspring)
			case percent < 0.99:
				mutateSwap(offspring)
			case percent < 1.00:
				mutateQuantization(10.0, offspring)
			}
			evaluateFitness(mappings, offspring)
		}
		for i := range population {
			if population[i].fitness < solution.fitness {
				solution = &population[i]
			}
		}
		io.PrintProgress(target / solution.fitness)
	}
	io.PrintComplete()
	simulation := vm.State{}
	fmt.Println(mappings[0].Input, "->", simulation.Run(mappings[0].Input, solution.Program))
	fmt.Println(mappings[1].Input, "->", simulation.Run(mappings[1].Input, solution.Program))
	fmt.Println(mappings[2].Input, "->", simulation.Run(mappings[2].Input, solution.Program))
	fmt.Println(mappings[3].Input, "->", simulation.Run(mappings[3].Input, solution.Program))
	fmt.Println(mappings[4].Input, "->", simulation.Run(mappings[4].Input, solution.Program))
	fmt.Println(mappings[5].Input, "->", simulation.Run(mappings[5].Input, solution.Program))
	fmt.Println(mappings[6].Input, "->", simulation.Run(mappings[6].Input, solution.Program))
	fmt.Println(mappings[7].Input, "->", simulation.Run(mappings[7].Input, solution.Program))
	fmt.Println(mappings[8].Input, "->", simulation.Run(mappings[8].Input, solution.Program))
	fmt.Println(mappings[9].Input, "->", simulation.Run(mappings[9].Input, solution.Program))
	return solution.Program
}

