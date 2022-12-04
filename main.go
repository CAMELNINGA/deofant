package main

import (
	"fmt"
	"math/rand"
)

const (
	Result = 30
	K1     = 1
	K2     = 2
	K3     = 3
	K4     = 4
	K5     = 5
)

func main() {

	genomes := make(map[int]*Genome, 50)
	for i := 0; i < 50; i++ {
		genomes[i] = NewGenerateGenome(Result)
	}
	G := Genomes{
		g: genomes,
	}
	Have := false
	for version := 1; version < 10000; version++ {
		fmt.Printf("version %d \n", version)
		var per []float32
		per, Have = G.Perchent(K1, K2, K3, K4, K5, Result)
		if Have {

			break
		}
		children := make(map[int]*Genome, 50)
		for i := range G.g {
			parent1, parent2 := ParentID(per)
			fmt.Printf("parent1 :%d \n parent2 : %d \n", parent1, parent2)
			children[i] = CrossOver(G.g[parent1], G.g[parent2])
		}
		children[rand.Intn(50)] = NewGenerateGenome(Result)
		G.g = children
	}

	if Have {
		fmt.Println("Find")
	} else {
		fmt.Println("Dont Find")
	}
}
