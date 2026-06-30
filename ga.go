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
	x := rand.Intn(len(individuals))
	y := (x + 1) % len(individuals)
	d := rand.Intn(len(individuals))
	for d == x || d == y {
		d = rand.Intn(len(individuals))
	}
	return &individuals[x], &individuals[y], &individuals[d]
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

func (i *individual) fission(offspring *individual) {
	copy(offspring.constants, i.constants)
	offspring.instructions = append([]uint16(nil), i.instructions...)
}

func (i *individual) transfer(offspring *individual) {
	if rand.Intn(100) == 0 {
		n := rand.Intn(min(len(i.instructions), 10)) + 1
		t := rand.Intn(len(i.instructions) + 1 - n)
		body := i.instructions[t:t+n]
		o := rand.Intn(len(offspring.instructions) + 1)
		head := offspring.instructions[:o]
		tail := offspring.instructions[o:]
		offspring.instructions = append(head, append(body, tail...)...)
	}
}

func (i *individual) mutate() {
	switch rand.Intn(4) {
	case 0:
		m := rand.Intn(len(i.constants))
		i.constants[m] += rand.Float64()*0.002 - 0.001
	case 1:
		m := rand.Intn(len(i.instructions))
		i.instructions[m] = uint16(rand.Int())
	case 2:
		if len(i.instructions) <= 1 {
			break
		}
		m := rand.Intn(len(i.instructions))
		head := i.instructions[:m]
		tail := i.instructions[m+1:]
		i.instructions = append(head, tail...)
	case 3:
		m := rand.Intn(len(i.instructions))
		head := i.instructions[:m]
		body := []uint16{uint16(rand.Int())}
		tail := i.instructions[m:]
		i.instructions = append(head, append(body, tail...)...)
	}
}

func (i *individual) evaluate(inputs [][]float64, outputs []float64, penalty float64) {
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
			delta = penalty
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
			parent.fission(offspring)
			parent.mu.Unlock()
			donor.transfer(offspring)
			donor.mu.Unlock()
			offspring.mutate()
			offspring.evaluate(inputs, outputs, total)
			mu.Lock()
			if offspring.fitness < lowest {
				lowest = offspring.fitness
				solutions <- solution{
					fitness:      100 - offspring.fitness / total * 100,
					constants:    append([]float64(nil), offspring.constants...),
					instructions: append([]uint16(nil), offspring.instructions...),
				}
			}
			mu.Unlock()
			offspring.mu.Unlock()
		}(parentX, parentY, donor)
	}
}
