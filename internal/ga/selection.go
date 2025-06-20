package ga

import (
	"math/rand"
)

func selection(population []individual, distance int) (*individual, *individual) {
	i := rand.Intn(len(population))
	j := (i + rand.Intn(distance) + 1) % len(population)
	return &population[i], &population[j]
}
