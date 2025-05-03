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
	Program struct {
		fitness      float64
		Registers    []float32
		Instructions []uint16
	}
)

func initialization(instructions, programs int) []Program {
	population := make([]Program, programs)
	for i := range population {
		population[i].fitness = math.MaxFloat64
		population[i].Registers = make([]float32, vm.Registers)
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
	i := rand.Intn(len(offspring.Registers) + 1)
	copy(offspring.Registers[:i], parent1.Registers[:i])
	copy(offspring.Registers[i:], parent2.Registers[i:])
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

func mutationPerturbation(min, max float32, offspring *Program) {
	i := rand.Intn(len(offspring.Registers))
	offspring.Registers[i] += rand.Float32() * (max - min) + min
}

func mutationQuantization(scale float32, offspring *Program) {
	i := rand.Intn(len(offspring.Registers))
	scaled := int(offspring.Registers[i] * scale)
	offspring.Registers[i] = float32(scaled) / scale
}

//func mutationAnnihilation(offspring *Program) {
//	i := rand.Intn(len(offspring.Registers))
//	offspring.Registers[i] = 0
//}

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
		m.Reset(candidate.Registers)
		for j := range tests[i].Inputs {
			m.Set(j, tests[i].Inputs[j])
		}
		for j := range candidate.Instructions {
			m.Execute(candidate.Instructions[j])
		}
		for j := range tests[i].Expected {
			actual := float64(m.Get(j))
			expected := float64(tests[i].Expected[j])
			candidate.fitness += math.Abs(actual - expected)
		}
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
		//mutationAnnihilation(offspring)
	}
	evaluation(tests, offspring)
}

func Evolution(tests []Test, target float64, instructions, programs int) *Program {
	population := initialization(instructions, programs)
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
		m.Reset(solution.Registers)
		for j := range tests[i].Inputs {
			m.Set(j, tests[i].Inputs[j])
		}
		for j := range solution.Instructions {
			m.Execute(solution.Instructions[j])
		}
		fmt.Print("answer: ")
		for j := range tests[i].Expected {
			fmt.Print(m.Get(j), " ")
		}
		fmt.Println()
	}
	return solution
}

