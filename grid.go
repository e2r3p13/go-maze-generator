package go_maze_generator

import "math/rand"
import "time"

type Grid struct {
	W int
	H int
	generator Algorithm
	cells []Cell
}

// Creates and returns a grid
// /!\ The grid needs to be initialized before use /!\
func New_grid(w, h int) Grid {
	return Grid {
		W: w,
		H: h,
		cells: make([]Cell, w * h),
	}
}

// Returns the cell at the @x, @y coordinates in the grid
func (g *Grid) At(x, y int) *Cell {
	if (x >= 0 && x < g.W && y >= 0 && y < g.H) {
		return &g.cells[g.W * y + x]
	} else {
		return nil
	}
}

// Returns the @i nth cell of the grid
func (g *Grid) At_index(i int) *Cell {
	if i >= 0 && i < g.Size() {
		return &g.cells[i]
	} else {
		return nil
	}
}

// Pick and returns a random cell from the grid
func (g *Grid) Random_cell() *Cell {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(g.Size())
	return g.At_index(index)
}

// Returns the number of cells the grid contains
func (g *Grid) Size() int {
	return g.W * g.H
}

// Initializes the grid by connecting the cells.
// @linked determines weither the cells are linked by default or not
func (g *Grid) Initialize(generator Algorithm) {
	for y := 0; y < g.H; y++ {
		for x := 0; x < g.W; x++ {
			cell := g.At(x, y)
			cell.n = Link {g.At(x, y - 1), false}
			cell.s = Link {g.At(x, y + 1), false}
			cell.e = Link {g.At(x + 1, y), false}
			cell.w = Link {g.At(x - 1, y), false}
		}
	}
	g.generator = generator
	generator.init_for(g)
}

// Returns a string represening the grid with ascii characters
func (g *Grid) To_s() string {
	str := "+"
	for i := 0; i < g.H; i++ {
		str += "---+"
	}
	str += "\n"
	for y := 0; y < g.H; y++ {
		top := "|"
		bottom := "+"
		for x := 0; x < g.W; x++ {
			if g.At(x, y).e.linked {
				top += "    "
			} else {
				top += "   |"
			}
			if g.At(x, y).s.linked {
				bottom += "   +"
			} else {
				bottom += "---+"
			}
		}
		str += top + "\n" + bottom + "\n"
	}
	return str
}

func (g *Grid) Is_fully_generated() bool {
	return g.generator.is_over()
}

func (g *Grid)Generate() {
	g.generator.perform()
}

func (g *Grid)Generate_step() {
	g.Apply_x_steps(1)
}

func (g *Grid)Apply_x_steps(x int) {
	for i := 0; i < x && !g.generator.is_over(); i++ {
		g.generator.perform_step()
	}
}
