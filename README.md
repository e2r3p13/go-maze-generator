# go-maze-generator

A Go package to generate mazes, inspired by Jamis Buck, from the book *"Mazes for Programmers"*.

Currently learning Go, any code related advice is welcome.

## Installation

```
git clone https://github.com/lfalkau/go-maze-generator.git
cd go-maze-generator
sudo go install
```

## Basic usage example
```
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
```
More examples are coming

## 🔨 Work In Progress

### General
- [x] Create basic data structures
- [x] Add methods to use those structures
- [x] Add a method to print ascii grid
- [x] Split code in separate packages
- [ ] Handle errors in a Go fashion way
- [ ] Implement generations algorithms and methods
- [ ] Add a method to export the grid as an image
- [ ] Write examples

### Algorithms
- [x] Binary tree
- [ ] Growing tree
- [ ] Sidewinder
- [ ] Aldous-Broder
- [ ] Wilson
- [ ] Hunt-and-Kill
- [ ] Recursive Backtracker
- [ ] Recursive Division
- [ ] Krusal
- [ ] Prim (simplified)
- [ ] Prim (true)
- [ ] Eller
