package ga

import (
	"fmt"
	"math"
	"math/rand"
	"github.com/pehringer/fungen/internal/vm"
)

type (
	Program struct {
		fitness      float64
		memory       []float64
		instructions []uint16
	}
)

func initialization(memory, instructions, programs int) []Program {
	population := make([]Program, programs)
	for i := range population {
		population[i].fitness = math.MaxFloat64
		population[i].memory = make([]float64, memory)
		population[i].instructions = make([]uint16, instructions)
	}
	return population
}

func selection(neighborhood int, population []Program) (*Program, *Program) {
	i := rand.Intn(len(population))
	j := (i + rand.Intn(neighborhood) + 1) % len(population)
	return &population[i], &population[j]
}

func evaluation(inputs, outputs []float64, candidate *Program) {
	candidate.fitness = 0
	m := vm.Machine{}
	for i := range inputs {
		m.Set(inputs[i], candidate.memory)
		for j := range candidate.instructions {
			m.Execute(candidate.instructions[j])
		}
		candidate.fitness += math.Abs(m.Get() - outputs[i])
	}
}

func dualStrategy(inputs, outputs []float64, population []Program) {
	parent1, parent2 := selection(10, population)
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
	evaluation(inputs, outputs, offspring)
}

func Evolution(inputs, outputs []float64, target float64, memory, instructions, programs int) ([]float64, []uint16) {
	population := initialization(memory, instructions, programs)
	solution := &population[0]
	for solution.fitness > target {
		for range len(population) {
			dualStrategy(inputs, outputs, population)
		}
		for i := range population {
			if population[i].fitness < solution.fitness {
				solution = &population[i]
				fmt.Println("solution error:", solution.fitness)
			}
		}
	}
	m := vm.Machine{}
	for i := range inputs {
		m.Set(inputs[i], solution.memory)
		for j := range solution.instructions {
			m.Execute(solution.instructions[j])
		}
		fmt.Println("answer:", m.Get())
	}
	return solution.memory, solution.instructions
}

