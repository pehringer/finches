package ga

import (
	"math/rand"
)

func selectNeighbors(neighborhood int, population []individual) (*individual, *individual) {
	i := rand.Intn(len(population))
	j := (i + rand.Intn(neighborhood) + 1) % len(population)
	return &population[i], &population[j]
}
