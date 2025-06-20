package ga

import (
	"math/rand"
)

func flip(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	mask := uint16(1) << rand.Intn(16)
	offspring.instructions[i] ^= mask
}

func step(offspring *individual) {
	i := rand.Intn(len(offspring.constants))
	amount := rand.Float64() * 0.00002 - 0.00001
	offspring.constants[i] += amount
}

func leap(offspring *individual) {
	i := rand.Intn(len(offspring.constants))
	amount := rand.Float64() * 2.0 - 1.0
	offspring.constants[i] += amount
}

func swap(offspring *individual) {
	i := rand.Intn(len(offspring.instructions))
	j := rand.Intn(len(offspring.instructions))
	temporary := offspring.instructions[i]
	offspring.instructions[i] = offspring.instructions[j]
	offspring.instructions[j] = temporary
}

func mutation(offspring *individual) {
	percent := rand.Float64()
	switch {
	case percent < 0.50:
		flip(offspring)
	case percent < 0.70:
		swap(offspring)
	case percent < 0.90:
		step(offspring)
	case percent < 1.00:
		leap(offspring)
	}
}
