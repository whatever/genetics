package main

type Gene interface {
	Score() float32
	Mix(*Gene) Gene
	Swap(*Gene) Gene
	Mutate() Gene
}

type Genes []Gene

type GenePool interface {
	Generate() Genes
	Rate()
	Compete()
}
