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

type CurveFitGene struct {
	A float64
	B float64
	C float64
}

type CurveFitGenes []CurveFitGene

// Sorting interfac
func (s CurveFitGenes) Len() int {
	return len(s)
}

// Swap
func (s CurveFitGenes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s CurveFitGenes) Less(i, j int) bool {
	a := s[i].Score()
	b := s[j].Score()
	r := a < b
	return r
}

// Return values of a function
func (s *CurveFitGene) F(t float64) float64 {
	return s.A*t*t + s.B*t + s.C
}

// ...
func (s *CurveFitGene) Hash() string {
	return "000"
}

// Make Real Functions easier to write
type RealFunc func(t float64) float64

// ...
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
		-4.0, -3.0, -2.0, 1.0, 0.0,
		4.0, 3.0, 2.0, 1.0, 0.0,
	}

	sort.Float64s(samples)

	totalError := 0.0

	for _, t := range samples {
		totalError += math.Abs(f(t)/1000. - g(t)/1000.)
	}

	return totalError
}

// How well does it fit the problem?
func (self *CurveFitGene) Score() float64 {
	f := func(t float64) float64 {
		return 81*t*t - 200*t - 133.
		return 81*t*t*t - 200*t*t + t - 133.
	}
	_ = f
	return Delta(self.F, f)
}

// Give me a rand answer based on this one
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

// Give me a rand answer based on this one
func (s *CurveFitGene) Mutate() CurveFitGene {
	s.A += rand.NormFloat64()
	s.B += rand.NormFloat64()
	s.C += rand.NormFloat64()
	return *s
}

// Blend 2 solutions together
func RandomCurveFitGene() CurveFitGene {
	return CurveFitGene{
		rand.Float64()*1000. - 500.,
		rand.Float64()*1000. - 500.,
		rand.Float64()*1000. - 500.,
	}
}
