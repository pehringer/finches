package ga

import (
	"math/rand"
)

func mutateBitFlips(bits int, offspring *individual) {
	for i := 0; i < bits; i++ {
		j := rand.Intn(len(offspring.Instructions))
		offspring.Instructions[j] ^= uint16(1) << rand.Intn(16)
	}
}

func mutatePerturbation(min, max float64, offspring *individual) {
	i := rand.Intn(len(offspring.Data))
	offspring.Data[i] += rand.Float64() * (max - min) + min
}

func mutateQuantization(scale float64, offspring *individual) {
	i := rand.Intn(len(offspring.Data))
	scaled := int64(offspring.Data[i] * scale)
	offspring.Data[i] = float64(scaled) / scale
}

func mutateSwap(offspring *individual) {
	i := rand.Intn(len(offspring.Instructions))
	j := rand.Intn(len(offspring.Instructions))
	temporary := offspring.Instructions[i]
	offspring.Instructions[i] = offspring.Instructions[j]
	offspring.Instructions[j] = temporary
}
