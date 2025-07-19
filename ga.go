package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type individual struct {
	mu           sync.Mutex
	fitness      float64
	constants    []float64
	instructions []uint16
}

func initialize(population int, outputs []float64) []individual {
	individuals := make([]individual, population)
	for i := range individuals {
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
	if rand.Float64() < 0.01 {
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

func evaluate(inputs [][]float64, outputs []float64, penalty float64, offspring *individual) {
	offspring.fitness = 0
	for i := range min(len(inputs), len(outputs)) {
		output := simulateProgram(inputs[i], offspring.constants, offspring.instructions)
		delta := math.Abs(output - outputs[i])
		if math.IsNaN(delta) || math.IsInf(delta, 0) {
			offspring.fitness += penalty
			continue
		}
		offspring.fitness += delta
	}
}

func terminate(individuals []individual) *individual {
	alpha := &individuals[0]
	for i := range individuals {
		if individuals[i].fitness < alpha.fitness {
			alpha = &individuals[i]
		}
	}
	return alpha
}

func evolve(generations, population int, inputs [][]float64, outputs []float64) ([]float64, []uint16) {
	total := 0.0
	for i := range outputs {
		total += math.Abs(outputs[i])
	}
	individuals := initialize(population, outputs)
	for i := range generations {
		wg := sync.WaitGroup{}
		for range population {
			parentX, parentY, donor := seleCt(individuals)
			parentX.mu.Lock()
			parentY.mu.Lock()
			donor.mu.Lock()
			wg.Add(1)
			go func(parentX, parentY, donor *individual) {
				defer parentX.mu.Unlock()
				defer parentY.mu.Unlock()
				defer donor.mu.Unlock()
				defer wg.Done()
				evaluate(inputs, outputs, total, transfer(donor, mutate(fission(replace(parentX, parentY)))))
			}(parentX, parentY, donor)
		}
		wg.Wait()
		fmt.Printf("\r%.2f%%", float64(i)/float64(generations)*100)
	}
	alpha := terminate(individuals)
	fmt.Printf("\rInstructions: %d Error: %f%%\n", len(alpha.instructions), alpha.fitness/total*100)
	return alpha.constants, alpha.instructions
}
