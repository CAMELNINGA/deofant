package main

import (
	"crypto/rand"
	"math"
	"math/big"
)

type Genome struct {
	Genome []int
}

func NewGenerateGenome(Result int) *Genome {
	var genome []int
	for i := 0; i < 5; i++ {

		big, _ := rand.Int(rand.Reader, big.NewInt(int64(Result)))
		genome = append(genome, int(big.Int64()))
	}
	g := &Genome{
		Genome: genome,
	}
	return g
}

func NewGenome(genome []int) *Genome {
	return &Genome{
		Genome: genome,
	}
}

func (g *Genome) Fitness(K1, K2, K3, K4, K5, Result int) float64 {
	per := float64(g.Genome[0]*K1 + g.Genome[1]*K2 + g.Genome[2]*K3 + g.Genome[3]*K4 + g.Genome[4]*K5 - Result)
	return math.Abs(per)
}
