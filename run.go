package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Create a bunch of random genes
func Seed(n int) CurveFitGenes {
	results := make(CurveFitGenes, n)
	for i := 0; i < n; i++ {
		results[i] = RandomCurveFitGene()
	}
	return results
}

// Score everything
func Rank(genes CurveFitGenes) {
	sort.Sort(genes)
}

// Arrange in order of success,
// and make a probability function for this
func Select(genes CurveFitGenes) {
	sort.Sort(genes)

	best := genes[0:30]
	seed := Seed(30)

	_ = best
	_ = seed
}

// CROSS-OVER!!
func CrossOver(genes CurveFitGenes) CurveFitGenes {
	sort.Sort(genes)

	results := make(CurveFitGenes, 100)

	for i := 0; i < 20; i++ {
		results[i] = genes[i]
	}

	for i := 20; i < 50; i++ {
		h := results[rand.Int()%20]
		results[i] = genes[i].Mix(&h)
	}

	for i := 50; i < 80; i++ {
		results[i] = genes[i-50]
		results[i].Mutate()
	}

	for i := 80; i < 100; i++ {
		results[i] = RandomCurveFitGene()
	}

	return results
}

// MUTATE!!
func Mutate(genes CurveFitGenes) {
	for i, v := range genes {
		if i%3 == 0 {
			v.Mutate()
		}
	}
}

// ...

// MAIN
func main() {
	genePool := Seed(100)

	for i := 0; i < 1000; i++ {
		genePool = CrossOver(genePool)
		Mutate(genePool)
	}

	fmt.Println(genePool[0].Score())
	fmt.Println(genePool[0])
}

func x_x() {
	fmt.Println("x_x")
}
