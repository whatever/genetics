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

type CurveFitGenePool struct {
	genes CurveFitGenes
}

// SEED!!
func (s *CurveFitGenePool) Generate(n int) {
	s.genes = make(CurveFitGenes, n)
	for i := 0; i < n; i++ {
		s.genes[i] = RandomCurveFitGene()
	}
}

func ScaleDown(n int, percent float64) int {
	return int(float64(n) * percent)
}

// CROSSOVER!!
func (s *CurveFitGenePool) CrossOver() {
	sort.Sort(s.genes)

	results := make(CurveFitGenes, 0)

	size := len(s.genes)

	TOP_SIZE := ScaleDown(size, 0.2)
	MIX_SIZE := ScaleDown(size, 0.3)
	RND_SIZE := ScaleDown(size, 0.3)

	// Keep some number of top
	for i := 0; i < TOP_SIZE; i++ {
		results = append(results, s.genes[i])
	}

	// Mix top with top
	for i := 0; i < MIX_SIZE; i++ {
		h := results[rand.Int()%TOP_SIZE]
		results = append(
			results,
			s.genes[i].Mix(&h),
		)
	}

	// Mutate some of the bottom
	for i := 0; i < size-TOP_SIZE-MIX_SIZE-RND_SIZE; i++ {
		h := s.genes[i]
		results = append(
			results,
			h.Mutate(),
		)
	}

	// Add 20 random ones in
	for i := 0; i < RND_SIZE; i++ {
		results = append(
			results,
			RandomCurveFitGene(),
		)
	}

	// Replace old generation with the current one
	s.genes = results
}

// MUTATE!!
func (s *CurveFitGenePool) Mutate() {
	for i, v := range s.genes {
		if i%2 == 0 {
			v.Mutate()
		}
	}
}

// Return Best Gene
func (s *CurveFitGenePool) Best() CurveFitGene {
	return s.genes[0]
}

// MAIN
func main() {
	genes := CurveFitGenePool{[]CurveFitGene{}}
	genes.Generate(667)

	for i := 0; i < 1000; i++ {
		genes.CrossOver()
		genes.Mutate()
	}

	b := genes.Best()
	fmt.Println(b, b.Score())
}
