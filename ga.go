package main

import (
	"math"
	"math/rand"
	"sync"
)

type solution struct {
	fitness      float64
	constants    []float64
	instructions []uint16
}

type individual struct {
	mu sync.Mutex
	solution
}

func initialize(population int) []individual {
	individuals := make([]individual, population)
	for i := range population {
		individuals[i].fitness = math.MaxFloat64
		individuals[i].constants = []float64{
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
			rand.Float64()*200.0 - 100.0,
		}
		individuals[i].instructions = []uint16{
			uint16(rand.Int()),
		}
	}
	return individuals
}

func seleCt(individuals []individual) (*individual, *individual, *individual) {
	i := rand.Intn(len(individuals))
	j := (i + 1) % len(individuals)
	k := rand.Intn(len(individuals))
	for k == i || k == j {
		k = rand.Intn(len(individuals))
	}
	return &individuals[i], &individuals[j], &individuals[k]
}

func replace(parentX, parentY *individual) (*individual, *individual) {
	if parentX.fitness < parentY.fitness {
		return parentX, parentY
	}
	if parentY.fitness < parentX.fitness {
		return parentY, parentX
	}
	if len(parentX.instructions) < len(parentY.instructions) {
		return parentX, parentY
	}
	return parentY, parentX
}

func fission(parent, offspring *individual) *individual {
	copy(offspring.constants, parent.constants)
	offspring.instructions = make([]uint16, len(parent.instructions))
	copy(offspring.instructions, parent.instructions)
	return offspring
}

func mutate(offspring *individual) *individual {
	switch rand.Intn(4) {
	case 0:
		i := rand.Intn(len(offspring.constants))
		offspring.constants[i] += rand.Float64()*0.002 - 0.001
	case 1:
		i := rand.Intn(len(offspring.instructions))
		offspring.instructions[i] = uint16(rand.Int())
	case 2:
		if len(offspring.instructions) <= 1 {
			break
		}
		i := rand.Intn(len(offspring.instructions))
		head := offspring.instructions[:i]
		tail := offspring.instructions[i+1:]
		offspring.instructions = append(head, tail...)
	case 3:
		i := rand.Intn(len(offspring.instructions))
		head := offspring.instructions[:i]
		body := []uint16{uint16(rand.Int())}
		tail := offspring.instructions[i:]
		offspring.instructions = append(head, append(body, tail...)...)
	}
	return offspring
}

func transfer(donor, offspring *individual) *individual {
	if rand.Intn(100) == 0 {
		n := rand.Intn(min(len(donor.instructions), 10)) + 1
		i := rand.Intn(len(donor.instructions) + 1 - n)
		body := donor.instructions[i : i+n]
		j := rand.Intn(len(offspring.instructions) + 1)
		head := offspring.instructions[:j]
		tail := offspring.instructions[j:]
		offspring.instructions = append(head, append(body, tail...)...)
	}
	return offspring
}

func evaluate(inputs [][]float64, outputs [][]*float64, penalty float64, offspring *individual) *individual {
	offspring.fitness = 0
	machine := setupRegisters(offspring.constants)
	for i := range min(len(inputs), len(outputs)) {
		machine.executeInstructions(inputs[i], offspring.instructions)
		if len(outputs[i]) < 1 || outputs[i][0] == nil {
			continue
		}
		output := machine.resetRegisters(offspring.constants)
		delta := math.Abs(output - *outputs[i][0])
		if math.IsNaN(delta) || math.IsInf(delta, 0) {
			delta = penalty
		}
		offspring.fitness += delta
	}
	return offspring
}

func evolve(population int, inputs [][]float64, outputs [][]*float64, solutions chan<- solution) {
	total := 0.0
	for i := range outputs {
		for j := range outputs[i] {
			if outputs[i][j] != nil {
				total += math.Abs(*outputs[i][j])
			}
		}
	}
	mu := sync.Mutex{}
	lowest := total
	individuals := initialize(population)
	for {
		parentX, parentY, donor := seleCt(individuals)
		parentX.mu.Lock()
		parentY.mu.Lock()
		donor.mu.Lock()
		go func(parentX, parentY, donor *individual) {
			parent, offspring := replace(parentX, parentY)
			fission(parent, offspring)
			parent.mu.Unlock()
			transfer(donor, offspring)
			donor.mu.Unlock()
			mutate(offspring)
			evaluate(inputs, outputs, total, offspring)
			mu.Lock()
			if offspring.fitness < lowest {
				lowest = offspring.fitness
				solutions <- solution{
					fitness:      100 - offspring.fitness / total * 100,
					constants:    append([]float64{}, offspring.constants...),
					instructions: append([]uint16{}, offspring.instructions...),
				}
			}
			mu.Unlock()
			offspring.mu.Unlock()
		}(parentX, parentY, donor)
	}
}
