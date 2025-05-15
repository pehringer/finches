package ga

import (
	"math/rand"
)

func crossoverSinglePoint(parent1, parent2, offspring *individual) {
	i := rand.Intn(len(offspring.Data) + 1)
	copy(offspring.Data[:i], parent1.Data[:i])
	copy(offspring.Data[i:], parent2.Data[i:])
	i = rand.Intn(len(offspring.Instructions) + 1)
	copy(offspring.Instructions[:i], parent1.Instructions[:i])
	copy(offspring.Instructions[i:], parent2.Instructions[i:])
}
