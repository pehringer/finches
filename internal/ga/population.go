package ga

import (
	"math"
	"sync"
	"math/rand"
)

type (
	individual struct {
		mu           sync.Mutex
		fitness      float64
		constants    []float64
		instructions []uint16
	}
)

func initialization(constants, instructions, individuals int) []individual {
	population := make([]individual, individuals)
	for i := range population {
		population[i].fitness = math.MaxFloat64
		population[i].constants = make([]float64, constants)
		for j := range population[i].constants {
			population[i].constants[j] = rand.Float64() * 200.0 - 100.0
		}
		population[i].instructions = make([]uint16, instructions)
		for j := range population[i].instructions {
			population[i].instructions[j] = uint16(rand.Int())
		}
	}
	return population
}

func incrementation(population []individual)  {
	for i := range population {
		population[i].fitness = math.MaxFloat64
		population[i].instructions = append(population[i].instructions, uint16(rand.Int()))
	}
}
