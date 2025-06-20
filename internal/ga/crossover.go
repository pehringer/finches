package ga

import (
	"math/rand"
)

func crossover(offspring, parent1, parent2 *individual) {
        i := rand.Intn(len(offspring.constants) + 1)
        copy(offspring.constants[:i], parent1.constants[:i])
        copy(offspring.constants[i:], parent2.constants[i:])
        i = rand.Intn(len(offspring.instructions) + 1)
        copy(offspring.instructions[:i], parent1.instructions[:i])
        copy(offspring.instructions[i:], parent2.instructions[i:])
}
