package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Genomes struct {
	g map[int]*Genome
}

func (genomes Genomes) Perchent(K1, K2, K3, K4, K5, Result int) ([]float32, bool) {
	var per []float32
	var sum float32
	for _, v := range genomes.g {
		unic := float32(v.Fitness(K1, K2, K3, K4, K5, Result))
		if unic == 0 {
			fmt.Println("we find")
			fmt.Println(v.Genome)
			return per, true
		}
		per = append(per, 1/unic)
		sum += 1 / unic

	}

	for i := range per {
		per[i] = per[i] / sum * 100
	}
	return per, false
}

func CrossOver(g1, g2 *Genome) *Genome {
	line, _ := rand.Int(rand.Reader, big.NewInt(int64(5)))
	var rnk []int
	rnk = append(rnk, g1.Genome[:line.Int64()]...)
	rnk = append(rnk, g2.Genome[line.Int64():]...)

	return NewGenome(rnk)
}

func ParentID(per []float32) (int, int) {
	var sum float32
	g1ID, _ := rand.Int(rand.Reader, big.NewInt(int64(100)))
	g2ID, _ := rand.Int(rand.Reader, big.NewInt(int64(100)))

	var parent1, parent2 int
	for i, v := range per {
		sum += v
		if float32(g1ID.Int64()) >= sum {

			parent1 = i
		}
		if float32(g2ID.Int64()) >= sum {
			parent2 = i
		}
	}
	return parent1, parent2
}
