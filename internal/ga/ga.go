package ga

import (
	"fmt"
	"math"
	"math/rand"
	"github.com/pehringer/fungen/internal/vm"
)

type (
	Test struct {
		Input  float64
		Output float64
	}
	Program struct {
		fitness      float64
		Memory       []float64
		Instructions []uint16
	}
)

func initialization(memory, instructions, programs int) []Program {
	population := make([]Program, programs)
	for i := range population {
		population[i].fitness = math.MaxFloat64
		population[i].Memory = make([]float64, memory)
		population[i].Instructions = make([]uint16, instructions)
	}
	return population
}

func selectionNeighbors(neighborhood int, population []Program) (*Program, *Program) {
	i := rand.Intn(len(population))
	j := (i + rand.Intn(neighborhood) + 1) % len(population)
	return &population[i], &population[j]
}

func crossoverSinglePoint(parent1, parent2, offspring *Program) {
	i := rand.Intn(len(offspring.Memory) + 1)
	copy(offspring.Memory[:i], parent1.Memory[:i])
	copy(offspring.Memory[i:], parent2.Memory[i:])
	i = rand.Intn(len(offspring.Instructions) + 1)
	copy(offspring.Instructions[:i], parent1.Instructions[:i])
	copy(offspring.Instructions[i:], parent2.Instructions[i:])
}

func mutationBitFlips(bits int, offspring *Program) {
	for i := 0; i < bits; i++ {
		j := rand.Intn(len(offspring.Instructions))
		offspring.Instructions[j] ^= uint16(1) << rand.Intn(16)
	}
}

func mutationPerturbation(min, max float64, offspring *Program) {
	i := rand.Intn(len(offspring.Memory))
	offspring.Memory[i] += rand.Float64() * (max - min) + min
}

func mutationQuantization(scale float64, offspring *Program) {
	i := rand.Intn(len(offspring.Memory))
	scaled := int64(offspring.Memory[i] * scale)
	offspring.Memory[i] = float64(scaled) / scale
}

func mutationSwap(offspring *Program) {
	i := rand.Intn(len(offspring.Instructions))
	j := rand.Intn(len(offspring.Instructions))
	temporary := offspring.Instructions[i]
	offspring.Instructions[i] = offspring.Instructions[j]
	offspring.Instructions[j] = temporary
}

func evaluation(tests []Test, candidate *Program) {
	candidate.fitness = 0
	m := vm.Machine{}
	for i := range tests {
		m.Set(tests[i].Input, candidate.Memory)
		for j := range candidate.Instructions {
			m.Execute(candidate.Instructions[j])
		}
		candidate.fitness += math.Abs(m.Get() - tests[i].Output)
	}
}

func dualStrategy(tests []Test, population []Program) {
	parent1, parent2 := selectionNeighbors(10, population)
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
	evaluation(tests, offspring)
}

func Evolution(tests []Test, target float64, memory, instructions, programs int) *Program {
	population := initialization(memory, instructions, programs)
	solution := &population[0]
	for solution.fitness > target {
		for range len(population) {
			dualStrategy(tests, population)
		}
		for i := range population {
			if population[i].fitness < solution.fitness {
				solution = &population[i]
				fmt.Println("solution error:", solution.fitness)
			}
		}
	}
	m := vm.Machine{}
	for i := range tests {
		m.Set(tests[i].Input, solution.Memory)
		for j := range solution.Instructions {
			m.Execute(solution.Instructions[j])
		}
		fmt.Println("answer:", m.Get())
	}
	return solution
}

