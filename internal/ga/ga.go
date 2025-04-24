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

func crossoverSinglePoint(parent1, parent2, offspring *individual) {
	i := rand.Intn(len(offspring.registers)+1)
	copy(offspring.registers[:i], parent1.registers[:i])
	copy(offspring.registers[i:], parent2.registers[i:])
	i = rand.Intn(len(offspring.instructions)+1)
	copy(offspring.instructions[:i], parent1.instructions[:i])
	copy(offspring.instructions[i:], parent2.instructions[i:])
}

func mutationBitFlip(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	offspring.instructions[i] ^= uint32(1) << rand.Intn(32)
}

func mutationBitFlipCondition(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	offspring.instructions[i] ^= uint32(rand.Intn(4)) << 30
}

func mutationBitFlipOperation(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	offspring.instructions[i] ^= uint32(rand.Intn(32)) << 25
}

func mutationBitFlipSetFlags(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	offspring.instructions[i] ^= uint32(rand.Intn(2)) << 24
}

func mutationBitFlipDestination(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	offspring.instructions[i] ^= uint32(rand.Intn(256)) << 16
}

func mutationBitFlipSource1(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	offspring.instructions[i] ^= uint32(rand.Intn(256)) << 8
}

func mutationBitFlipSource2(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	offspring.instructions[i] ^= uint32(rand.Intn(256)) << 0
}

func mutationPerturbationOnes(offspring *individual) {
	i := rand.Intn(len(offspring.registers))
	offspring.registers[i] += rand.Float32() * 2.0 - 1.0
}

func mutationPerturbationTenths(offspring *individual) {
	i := rand.Intn(len(offspring.registers))
	offspring.registers[i] += rand.Float32() * 0.2 - 0.1
}

func mutationPerturbationHundredths(offspring *individual) {
	i := rand.Intn(len(offspring.registers))
	offspring.registers[i] += rand.Float32() * 0.02 - 0.01
}

func mutationPerturbationThousandths(offspring *individual) {
	i := rand.Intn(len(offspring.registers))
	offspring.registers[i] += rand.Float32() * 0.002 - 0.001
}

func mutationSwapInstructions(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	j := rand.Intn(len(offspring.instructions))
	value := offspring.instructions[i]
	offspring.instructions[i] = offspring.instructions[j]
	offspring.instructions[j] = value
}

func evaluation(tests []Test, candidate *individual) {
	candidate.fitness = 0
	for i := range tests {
		m := vm.NewMachine(candidate.registers)
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
	}
}

func dualStrategy(tests []Test, population []individual) float64 {
	i := rand.Intn(len(population))
	parent1 := &population[i]
	i = (i + rand.Intn(5) + 1) % len(population)
	parent2 := &population[i]
	offspring := parent1
	if parent1.fitness < parent2.fitness {
		offspring = parent2
	}
	crossoverSinglePoint(parent1, parent2, offspring)
	if rand.Intn(2) > 0 {
		switch rand.Intn(3) {
		case 0:
			mutationBitFlip(offspring)
		case 1:
			mutationPerturbationHundredths(offspring)
		case 2:
			mutationPerturbationThousandths(offspring)
		}
	} else {
		switch rand.Intn(9) {
		case 0:
			mutationBitFlipCondition(offspring)
		case 1:
			mutationBitFlipOperation(offspring)
		case 2:
			mutationBitFlipSetFlags(offspring)
		case 3:
			mutationBitFlipDestination(offspring)
		case 4:
			mutationBitFlipSource1(offspring)
		case 5:
			mutationBitFlipSource2(offspring)
		case 6:
			mutationPerturbationTenths(offspring)
		case 7:
			mutationPerturbationOnes(offspring)
		case 8:
			mutationSwapInstructions(offspring)
		}
	}
	evaluation(tests, offspring)
	return min(parent1.fitness, parent2.fitness)
}

func Evolution(tests []Test, target float64, registers, instructions, programs int) {
	population := initialization(registers, instructions, programs)
	best := 1000000000.0
	for {
		actual := dualStrategy(tests, population)
		if actual < best {
			best = actual
			fmt.Println(best)
		}
		if actual <= target {
			break
		}
	}
	return
}
