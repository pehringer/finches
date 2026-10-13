package main

import (
	"math"
	"math/rand"
	"sync"
)

type solution struct {
	fitness      float64
	constants    []float64
	instructions []int
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
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
			rand.Float64() * 200.0 - 100.0,
		}
		individuals[i].instructions = []int{
			rand.Int(),
		}
	}
	return individuals
}

func seleCt(individuals []individual) (*individual, *individual, *individual) {
	x := rand.Intn(len(individuals))
	y := (x + 1) % len(individuals)
	d := rand.Intn(len(individuals))
	for d == x || d == y {
		d = rand.Intn(len(individuals))
	}
	return &individuals[x], &individuals[y], &individuals[d]
}

func replace(parentX, parentY *individual) (*individual, *individual) {
	if parentX.fitness - parentY.fitness > 0.001 {
		return parentY, parentX
	}
	if parentY.fitness - parentX.fitness > 0.001 {
		return parentX, parentY
	}
	if len(parentX.instructions) >= len(parentY.instructions) {
		return parentY, parentX
	}
	return parentX, parentY
}

func (i *individual) fission(parent *individual) {
	copy(i.constants, parent.constants)
	i.instructions = append([]int(nil), parent.instructions...)
}

func (i *individual) mutate(donor *individual) {
	switch rand.Intn(6) {
	case 0:
		m := rand.Intn(len(i.constants))
		i.constants[m] += rand.Float64()* 0.002 - 0.001
	case 1:
		m := rand.Intn(len(i.constants))
		i.constants[m] = rand.Float64()* 200.0 - 100.0
	case 2:
		m := rand.Intn(len(i.instructions))
		i.instructions[m] ^= 1 << rand.Intn(15)
	case 3:
		body := []int{rand.Int()}
		m := rand.Intn(len(i.instructions))
		head := i.instructions[:m]
		tail := i.instructions[m:]
		i.instructions = append(head, append(body, tail...)...)
	case 4:
		n := rand.Intn(len(donor.instructions))
		body := []int{donor.instructions[n]}
		m := rand.Intn(len(i.instructions))
		head := i.instructions[:m]
		tail := i.instructions[m:]
		i.instructions = append(head, append(body, tail...)...)
	case 5:
		if len(i.instructions) <= 1 {
			break
		}
		m := rand.Intn(len(i.instructions))
		head := i.instructions[:m]
		tail := i.instructions[m+1:]
		i.instructions = append(head, tail...)
	}
	if rand.Intn(2) == 0 {
		i.mutate(donor)
	}
}

func (i *individual) evaluate(inputs [][]float64, outputs []float64) {
	i.fitness = 0
	machine := setup(i.constants)
	for s := range min(len(inputs), len(outputs)) {
		machine.process(inputs[s], i.instructions)
		if math.IsNaN(outputs[s]) {
			continue
		}
		output := machine.reset(i.constants)
		delta := math.Abs(output - outputs[s])
		if math.IsNaN(delta) || math.IsInf(delta, 0) {
			i.fitness = math.MaxFloat64
			break
		}
		i.fitness += delta
	}
}

func evolve(population int, inputs [][]float64, outputs []float64, solutions chan<- solution) {
	total := 0.0
	for i := range outputs {
		if math.IsNaN(outputs[i]) {
			continue
		}
		total += math.Abs(outputs[i])
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
			offspring.fission(parent)
			parent.mu.Unlock()
			offspring.mutate(donor)
			donor.mu.Unlock()
			offspring.evaluate(inputs, outputs)
			mu.Lock()
			if offspring.fitness <= lowest {
				lowest = offspring.fitness
				solutions <- solution{
					fitness:      100 - offspring.fitness / total * 100,
					constants:    append([]float64(nil), offspring.constants...),
					instructions: append([]int(nil), offspring.instructions...),
				}
			}
			mu.Unlock()
			offspring.mu.Unlock()
		}(parentX, parentY, donor)
	}
}
