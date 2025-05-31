package ga

import (
	"math/rand"
)

func mutateBitFlips(bits int, offspring *individual) {
	i := rand.Intn(len(offspring.Instructions))
	for j := 0; j < bits; j++ {
		offspring.Instructions[i] ^= uint16(1) << rand.Intn(16)
	}
}

func mutatePerturbation(min, max float64, offspring *individual) {
	i := rand.Intn(len(offspring.Data))
	offspring.Data[i] += rand.Float64() * (max - min) + min
}

func mutateSwap(offspring *individual) {
	i := rand.Intn(len(offspring.Instructions))
	j := rand.Intn(len(offspring.Instructions))
	temporary := offspring.Instructions[i]
	offspring.Instructions[i] = offspring.Instructions[j]
	offspring.Instructions[j] = temporary
}
