package ga

import (
	"fmt"
	"sync"
	"math/rand"
	"github.com/pehringer/mapper/internal/io"
	"github.com/pehringer/mapper/internal/vm"
	"github.com/pehringer/mapper/internal/types"
)

type (
	individual struct {
		mu sync.Mutex
		fitness float64
		types.Program
	}
)

func termination(mappings []types.Mapping, accuracy float64) float64 {
	target := 0.0
	for i := range mappings {
		target += mappings[i].AbsoluteOutput()
	}
	return target - target * accuracy
}

func Evolve(mappings []types.Mapping, accuracy float64, instructions, individuals int) types.Program {
	population := initialize(8, instructions, individuals)
	io.PrintStarting()
	wg := sync.WaitGroup{}
	for i := range (4194304 * 2) {
		parent1, parent2 := selectNeighbors(4, population)
		wg.Add(1)
		parent1.mu.Lock()
		parent2.mu.Lock()
		go func(parent1, parent2 *individual) {
			defer parent1.mu.Unlock()
			defer parent2.mu.Unlock()
			defer wg.Done()
			offspring := replaceDuel(parent1, parent2)
			crossoverSinglePoint(parent1, parent2, offspring)
			percent := rand.Float32()
			switch {
			case percent < 0.20:
				mutateBitFlips(1, offspring)
			case percent < 0.40:
				mutatePerturbation(-0.001, +0.001, offspring)
			case percent < 0.55:
				mutateScramble(1, offspring)
			case percent < 0.70:
				mutateSwap(1, offspring)
			case percent < 0.80:
				mutateBitFlips(3, offspring)
			case percent < 0.90:
				mutatePerturbation(-0.1, +0.1, offspring)
			case percent < 1.00:
				mutateScramble(3, offspring)
			}
			batch := make([]types.Mapping, 8)
			for i := range len(batch) {
				batch[i] = mappings[rand.Intn(len(mappings))]
			}
			evaluateFitness(batch, offspring)
		}(parent1, parent2)
		io.PrintProgress(float64(i) / (4194304 * 2))
	}
	wg.Wait()
	io.PrintComplete()
	target := termination(mappings, accuracy)
	solution := &population[0]
	for i := range population {
		evaluateFitness(mappings, &population[i])
		if population[i].fitness < solution.fitness {
			solution = &population[i]
		}
	}
	fmt.Println(target, "/", solution.fitness)
	simulation := vm.State{}
	fmt.Println(mappings[0].Input, "->", simulation.Run(mappings[0].Input, solution.Program))
	fmt.Println(mappings[1].Input, "->", simulation.Run(mappings[1].Input, solution.Program))
	fmt.Println(mappings[2].Input, "->", simulation.Run(mappings[2].Input, solution.Program))
	fmt.Println(mappings[3].Input, "->", simulation.Run(mappings[3].Input, solution.Program))
	fmt.Println(mappings[4].Input, "->", simulation.Run(mappings[4].Input, solution.Program))
	fmt.Println(mappings[5].Input, "->", simulation.Run(mappings[5].Input, solution.Program))
	fmt.Println(mappings[6].Input, "->", simulation.Run(mappings[6].Input, solution.Program))
	fmt.Println(mappings[7].Input, "->", simulation.Run(mappings[7].Input, solution.Program))
	fmt.Println(mappings[8].Input, "->", simulation.Run(mappings[8].Input, solution.Program))
	fmt.Println(mappings[9].Input, "->", simulation.Run(mappings[9].Input, solution.Program))
	return solution.Program
}

