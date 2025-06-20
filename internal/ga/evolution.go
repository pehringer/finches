package ga

import (
	"fmt"
	"math"
	"sync"
)

func Evolution(inputs [][]float64, outputs []float64, generations, individuals int) ([]float64, []uint16) {
	penalty := 0.0
	for i := range outputs {
		penalty += math.Abs(outputs[i])
	}
	fmt.Println(penalty)
	population := initialization(16, 1, individuals)
	solution := finalist(population)
	prior := solution.fitness
	count := 0
	for range generations {
		wg := sync.WaitGroup{}
		for range population {
			parent1, parent2 := selection(population, 8)
			parent1.mu.Lock()
			parent2.mu.Lock()
			wg.Add(1)
			go func(parent1, parent2 *individual) {
				defer parent1.mu.Unlock()
				defer parent2.mu.Unlock()
				defer wg.Done()
				offspring := replacement(parent1, parent2)
				crossover(offspring, parent1, parent2)
				mutation(offspring)
				evaluation(offspring, penalty, inputs, outputs)
			}(parent1, parent2)
		}
		wg.Wait()
		solution := finalist(population)
		count++
		if count == 256 {
			if prior - solution.fitness < 1.0 {
				incrementation(population)
			}
			prior = solution.fitness
			count = 0
		}
		fmt.Println(len(solution.instructions), ":", solution.fitness)
	}
	return solution.constants, solution.instructions
}
