package ga

import (
	"math/rand"
)

func crossoverSinglePoint(parent1, parent2, offspring *Program) {
	i := rand.Intn(len(offspring.memory) + 1)
	copy(offspring.memory[:i], parent1.memory[:i])
	copy(offspring.memory[i:], parent2.memory[i:])
	i = rand.Intn(len(offspring.instructions) + 1)
	copy(offspring.instructions[:i], parent1.instructions[:i])
	copy(offspring.instructions[i:], parent2.instructions[i:])
}

func mutationBitFlips(bits int, offspring *Program) {
	for i := 0; i < bits; i++ {
		j := rand.Intn(len(offspring.instructions))
		offspring.instructions[j] ^= uint16(1) << rand.Intn(16)
	}
}

func mutationPerturbation(min, max float64, offspring *Program) {
	i := rand.Intn(len(offspring.memory))
	offspring.memory[i] += rand.Float64() * (max - min) + min
}

func mutationQuantization(scale float64, offspring *Program) {
	i := rand.Intn(len(offspring.memory))
	scaled := int64(offspring.memory[i] * scale)
	offspring.memory[i] = float64(scaled) / scale
}

func mutationSwap(offspring *Program) {
	i := rand.Intn(len(offspring.instructions))
	j := rand.Intn(len(offspring.instructions))
	temporary := offspring.instructions[i]
	offspring.instructions[i] = offspring.instructions[j]
	offspring.instructions[j] = temporary
}
