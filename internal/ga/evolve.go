package ga

import (
	"math"
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

func initialization(data, instructions, individuals int) []individual {
	population := make([]individual, individuals)
	for i := range population {
		population[i].fitness = math.MaxFloat64
		population[i].Program = types.EmptyProgram(data, instructions)
	}
	return population
}

func selection(neighborhood int, population []individual) (*individual, *individual) {
	i := rand.Intn(len(population))
	j := (i + rand.Intn(neighborhood) + 1) % len(population)
	return &population[i], &population[j]
}

func evaluation(mappings []types.Mapping, candidate *individual) {
	candidate.fitness = 0
	simulation := vm.State{}
	for i := range mappings {
		output := simulation.Run(mappings[i].Input, candidate.Program)
		candidate.fitness += math.Abs(output - mappings[i].Output)
	}
}

func replacement(mappings []types.Mapping, parent1, parent2 *individual) {
	offspring := parent1
	if parent1.fitness < parent2.fitness {
		offspring = parent2
	}
	crossoverSinglePoint(parent1, parent2, offspring)
	percent := rand.Float32()
	switch {
	case percent < 0.35:
		mutationBitFlips(1, offspring)
	case percent < 0.70:
		mutationPerturbation(-0.001, +0.001, offspring)
	case percent < 0.80:
		mutationBitFlips(3, offspring)
	case percent < 0.90:
		mutationPerturbation(-0.1, +0.1, offspring)
	case percent < 0.99:
		mutationSwap(offspring)
	case percent < 1.00:
		mutationQuantization(10.0, offspring)
	}
	evaluation(mappings, offspring)
}

func termination(mappings []types.Mapping, accuracy float64) float64 {
	target := 0.0
	for i := range mappings {
		target += mappings[i].AbsoluteOutput()
	}
	return target - target * accuracy
}

func Evolution(mappings []types.Mapping, accuracy float64, data, instructions, individuals int) types.Program {
	target := termination(mappings, accuracy)
	population := initialization(data, instructions, individuals)
	solution := &population[0]
	for solution.fitness > target {
		for range len(population) {
			parent1, parent2 := selection(10, population)
			replacement(mappings, parent1, parent2)
		}
		for i := range population {
			if population[i].fitness < solution.fitness {
				solution = &population[i]
			}
		}
		io.PrintProgress(target / solution.fitness)
	}
	io.PrintComplete()
	return solution.Program
}

