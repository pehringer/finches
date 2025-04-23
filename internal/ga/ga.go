package ga

import (
	"math"
	"math/rand"
	"github.com/pehringer/fungen/internal/vm"
)

type (
	test struct {
		inputs   []float32
		expected []float32
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

func crossoverSinglePoint(parent1, parent2, offspring *individual) {
	point := rand.Intn(len(offspring.registers)+1)
	copy(offspring.registers[:point], parent1.registers[:point])
	copy(offspring.registers[point:], parent2.registers[point:])
	point = rand.Intn(len(offspring.instructions)+1)
	copy(offspring.instructions[:point], parent1.instructions[:point])
	copy(offspring.instructions[point:], parent2.instructions[point:])
}

func mutationPerturbationBitFlip(offspring *individual) {
	point := rand.Intn(len(offspring.registers))
	offspring.registers[point] += rand.Float32() * 2 - 1
	point = rand.Intn(len(offspring.instructions))
	offspring.instructions[point] ^= uint32(1) << rand.Intn(32)
}

func evaluation(tests []test, candidate *individual) {
	candidate.fitness = 0
	for i := range tests {
		m := vm.NewMachine(candidate.registers)
		for j := range tests[i].inputs {
			m.SetRegister(j, tests[i].inputs[j])
		}
		for j := range candidate.instructions {
			m.Execute(candidate.instructions[j])
		}
		for j := range tests[i].expected {
			actual := float64(m.GetRegister(j))
			expected := float64(tests[i].expected[j])
			candidate.fitness += math.Abs(actual - expected)
		}
	}
}

func Evolution(tests []test, target float64, registers, instructions, programs int) {
}
