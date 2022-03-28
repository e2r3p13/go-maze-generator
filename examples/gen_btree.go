package main

import mazegen "go-maze-generator"
import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UnixNano())

	grid := mazegen.New_grid(16, 16)
	grid.Initialize(false)
	
	var bt = mazegen.New_BinaryTree()
	grid.Apply(bt)

	fmt.Println(grid.To_s())
}
