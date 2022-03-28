package main

import mazegen "go-maze-generator"
import "fmt"
import "math/rand"
import "time"
import "flag"

func main() {
	stepping := flag.Bool("stepping", false, "If provided, displays each steps of the generation")
	interval := flag.Int("interval", 250, "Sets the interval between two displayed steps")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	grid := mazegen.New_grid(16, 16)
	grid.Initialize(false)
	
	var bt = mazegen.New_Backtracking()
	
	if *stepping {
		fmt.Println(grid.To_s())
		grid.Apply_x_steps(bt, 0)
		for !grid.Is_fully_generated(bt) {
			time.Sleep(time.Duration(*interval) * time.Millisecond)
			grid.Apply_step(bt)
			fmt.Println(grid.To_s())
		}
	} else {
		grid.Apply(bt)
		fmt.Println(grid.To_s())
	}
}
