package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func xxx() {
	fmt.Println("x_x")
}

// Gene encoding of the curve-fitter
type CurveFitGene struct {
	A, B, C float64
}

// Just a simpler way to say a slice of genes
type CurveFitGenes []CurveFitGene

/**
 * Sortable
 * Sortable
 * Sortable
 */

// Return length of
func (s CurveFitGenes) Len() int {
	return len(s)
}

// Swap two elements
func (s CurveFitGenes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Return whether one element is less than another
func (s CurveFitGenes) Less(i, j int) bool {
	return s[i].Score() < s[j].Score()
}

/**
 * Gene
 * Gene
 * Gene
 */

// Return
func (self *CurveFitGene) Score() float64 {
	return Delta(
		self.F,
		func(t float64) float64 {
			return 2010*t*t - 200*t - 133.
			return 8100*t*t*t - 200*t*t + t - 133.
		},
	)
}

// Mix two genes
func (s *CurveFitGene) Mix(c *CurveFitGene) CurveFitGene {
	gene := CurveFitGene{s.A, s.B, s.C}

	switch rand.Int() % 3 {
	case 0:
		gene.A = c.A
	case 1:
		gene.B = c.B
	case 2:
		gene.C = c.C
	default:
	}

	return gene
}

// Mutate a single gene
func (s *CurveFitGene) Mutate() CurveFitGene {
	s.A += rand.NormFloat64() * 10
	s.B += rand.NormFloat64() * 10
	s.C += rand.NormFloat64() * 10
	return *s
}

// Return a random gene
func RandomCurveFitGene() CurveFitGene {
	LOW := -1500.0
	HIGH := 1500.0
	return CurveFitGene{
		rand.Float64()*(HIGH-LOW) + LOW,
		rand.Float64()*(HIGH-LOW) + LOW,
		rand.Float64()*(HIGH-LOW) + LOW,
	}
}

/**
 * Tertiary shit
 * Tertiary shit
 * Tertiary shit
 */

// Make Real Functions easier to write
type RealFunc func(t float64) float64

// Approximation
func Delta(f RealFunc, g RealFunc) float64 {

	SAMPLE_SIZE := 10
	LEFT := -50.
	RIGHT := 50.

	samples := make([]float64, SAMPLE_SIZE)

	for i := 0; i < SAMPLE_SIZE; i++ {
		samples[i] = rand.Float64()*(RIGHT-LEFT) + LEFT
	}

	// Omit random sampling for this problem

	samples = []float64{
		-4.0, -3.0, -2.0, 1.0,
		0.0,
		1.0, 2.0, 3.0, 4.0,
	}

	sort.Float64s(samples)

	totalError := 0.0

	for _, t := range samples {
		totalError += math.Abs(f(t)/1000. - g(t)/1000.)
	}

	return totalError
}

// Return values of a function
func (s *CurveFitGene) F(t float64) float64 {
	return s.A*t*t + s.B*t + s.C
}
