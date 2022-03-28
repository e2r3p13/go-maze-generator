package main

import mazegen "go-maze-generator"
import "fmt"
import "math/rand"
import "time"
import "flag"

func djb2(str string) uint64 {
	var hash uint64 = 5381
	for c := range str {
		hash = hash * 33 + uint64(c)
	}
	return hash % 1000
}

func algorithmFromString(alg string) mazegen.Algorithm {
	value := djb2(alg)
	switch value {
		case 39: return mazegen.New_Backtracking()
		case 866: return mazegen.New_BinaryTree()
	}
	return mazegen.New_Backtracking()
}

func main() {
	stepping := flag.Bool("stepping", false, "If provided, displays each steps of the generation")
	interval := flag.Int("interval", 250, "Sets the interval between two displayed steps")
	algorithm := flag.String("algorithm", "Backtracking", "Choose the generation algorithm'")
	size := flag.Int("size", 16, "side of the grid")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	grid := mazegen.New_grid(*size, *size)
	alg := algorithmFromString(*algorithm)
	grid.Initialize(alg)
	
	if *stepping {
		fmt.Println(grid.To_s())
		for !grid.Is_fully_generated() {
			time.Sleep(time.Duration(*interval) * time.Millisecond)
			grid.Generate_step()
			fmt.Println(grid.To_s())
		}
	} else {
		grid.Generate()
		fmt.Println(grid.To_s())
	}
}
