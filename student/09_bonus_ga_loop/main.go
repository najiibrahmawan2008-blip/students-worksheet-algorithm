package main

import (
	"fmt"
	"math/rand"
)

const (
	populationSize = 6
	chromosomeSize = 4
	maxGenerations = 10
	mutationRate   = 0.1
)

func fitness(candidate string) int {
	count := 0
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == '1' {
			count++
		}
	}
	return count
}

func randomCandidate() string {
	candidate := make([]byte, chromosomeSize)

	for i := 0; i < chromosomeSize; i++ {
		if rand.Intn(2) == 1 {
			candidate[i] = '1'
		} else {
			candidate[i] = '0'
		}
	}

	return string(candidate)
}

func selectParent(population []string) string {
	a := population[rand.Intn(len(population))]
	b := population[rand.Intn(len(population))]

	if fitness(a) >= fitness(b) { // prefer the fitter candidate
		return a
	}
	return b
}

func crossover(parent1, parent2 string) string {
	return parent1[:2] + parent2[2:] // keep part of each parent
}

func mutate(candidate string) string {
	bits := []byte(candidate)

	for i := 0; i < len(bits); i++ {
		if rand.Float64() < mutationRate { // mutate only sometimes
			if bits[i] == '0' {
				bits[i] = '1'
			} else {
				bits[i] = '0'
			}
		}
	}

	return string(bits)
}

func main() {

	population := make([]string, populationSize)

	for i := 0; i < populationSize; i++ {
		population[i] = randomCandidate()
	}

	for generation := 1; generation <= maxGenerations; generation++ {
		bestFitness := 0
		bestCandidate := ""

		for _, candidate := range population {
			score := fitness(candidate)

			if score > bestFitness {
				bestFitness = score
				bestCandidate = candidate
			}
		}

		fmt.Printf(
			"Generation %d: Best = %s, Fitness = %d\n",
			generation,
			bestCandidate,
			bestFitness,
		)

		if bestFitness == chromosomeSize { // stop when every bit is 1
			fmt.Println("Optimal solution found!")
			break
		}

		newPopulation := make([]string, populationSize)

		for i := 0; i < populationSize; i++ {
			parent1 := selectParent(population)
			parent2 := selectParent(population)

			child := crossover(parent1, parent2)
			child = mutate(child)

			newPopulation[i] = child
		}

		population = newPopulation
	}
}
