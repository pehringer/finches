package ga

func replaceDuel(parent1, parent2 *individual) *individual {
	offspring := parent1
	if parent1.fitness < parent2.fitness {
		offspring = parent2
	}
	return offspring
}
