package ga

import (
	"fmt"
	"math"
	"math/rand"
	"github.com/pehringer/fungen/internal/vm"
)

type (
	Test struct {
		Inputs   []float32
		Expected []float32
	}
	individual struct {
		fitness      float64
		registers    []float32
		instructions []uint32
	}
)

func initialization(registers, instructions, programs int) []individual {
	population := make([]individual, programs)
	for i := range population {
		population[i].fitness = math.MaxFloat64
		population[i].registers = make([]float32, registers)
		population[i].instructions = make([]uint32, instructions)
		for j := range population[i].registers {
			population[i].registers[j] = rand.Float32() * 20 - 10
		}
		for j := range population[i].instructions {
			population[i].instructions[j] = rand.Uint32()
		}
	}
	return population
}

func selectionNeighbors(neighborhood int, population []individual) (*individual, *individual) {
	i := rand.Intn(len(population))
	j := (i + rand.Intn(neighborhood) + 1) % len(population)
	return &population[i], &population[j]
}

func crossoverSinglePoint(parent1, parent2, offspring *individual) {
	i := rand.Intn(len(offspring.registers)+1)
	copy(offspring.registers[:i], parent1.registers[:i])
	copy(offspring.registers[i:], parent2.registers[i:])
	i = rand.Intn(len(offspring.instructions)+1)
	copy(offspring.instructions[:i], parent1.instructions[:i])
	copy(offspring.instructions[i:], parent2.instructions[i:])
}

func mutationBitFlips(bits int, offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	for j := 0; j < bits; j++ {
		offspring.instructions[i] ^= uint32(1) << rand.Intn(32)
	}
}

func mutationPerturbation(min, max float32, offspring *individual) {
	i := rand.Intn(len(offspring.registers))
	offspring.registers[i] += rand.Float32() * (max - min) + min
}

func mutationSwap(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	j := rand.Intn(len(offspring.instructions))
	value := offspring.instructions[i]
	offspring.instructions[i] = offspring.instructions[j]
	offspring.instructions[j] = value
}

func evaluation(tests []Test, candidate *individual) {
	candidate.fitness = 0
	m := vm.SetState(candidate.registers)
	for i := range tests {
		for j := range tests[i].Inputs {
			m.SetRegister(j, tests[i].Inputs[j])
		}
		for j := range candidate.instructions {
			m.Execute(candidate.instructions[j])
		}
		for j := range tests[i].Expected {
			actual := float64(m.GetRegister(j))
			expected := float64(tests[i].Expected[j])
			candidate.fitness += math.Abs(actual - expected)
		}
		m.ResetState(candidate.registers)
	}
}

func dualStrategy(tests []Test, population []individual) {
	parent1, parent2 := selectionNeighbors(8, population)
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
	case percent < 1.00:
		mutationSwap(offspring)
	}
	evaluation(tests, offspring)
}

func Evolution(tests []Test, target float64, registers, instructions, programs int) {
	population := initialization(registers, instructions, programs)
	best := math.MaxFloat64
	var solution *individual
	for solution == nil {
		for range len(population) {
			dualStrategy(tests, population)
		}
		for i := range population {
			if population[i].fitness < best {
				best = population[i].fitness
				fmt.Println(best)
			}
			if population[i].fitness <= target {
				solution = &population[i]
			}
		}
	}
	m := vm.SetState(solution.registers)
	for i := range tests {
		for j := range tests[i].Inputs {
			m.SetRegister(j, tests[i].Inputs[j])
		}
		for j := range solution.instructions {
			m.Execute(solution.instructions[j])
		}
		for j := range tests[i].Expected {
			fmt.Print(m.GetRegister(j))
		}
		fmt.Println()
		m.ResetState(solution.registers)
	}
	return
}
