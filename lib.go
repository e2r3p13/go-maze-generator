package main

import "fmt"
import "time"
import "math/rand"

type Grid struct {
	W int
	H int
	cells []Cell
}

type Link struct {
	dest *Cell 
	linked bool
}

type Cell struct {
	n Link
	s Link
	e Link
	w Link
}


func (c *Cell) links() []*Link {
	return []*Link{&c.n, &c.s, &c.e, &c.w}
}

func (c *Cell) adj_by_link(linked bool) []*Cell {
	cells := make([]*Cell, 0, 4)
	for _, l := range c.links() {
		if l.linked == linked {
			cells = append(cells, l.dest)
		}
	}
	return cells
}

// Returns the list of adjacent linked cells
func (c *Cell) Linked_cells() []*Cell {
	return c.adj_by_link(true)
}

// Returns the list of adjacent unlinked cells
func (c *Cell) Unlinked_cells() []*Cell {
	return c.adj_by_link(false)	
}

func (c *Cell) set_link(dest *Cell, bidi bool, state bool) {
	for _, l := range c.links() {
		if dest == l.dest {
			l.linked = state
			if bidi {
				dest.set_link(c, false, state)
			}
			break
		}
	}
}

// Links the cell with the @dest one.
// if bidi is set to true, also links @dest with the cell
func (c *Cell) Link(dest *Cell, bidi bool) {
	c.set_link(dest, bidi, true)
}

// Unlinks the cell from the @dest one.
// if bidi is set to true, also unlinks @dest from the cell
func (c *Cell) Unlink(dest *Cell, bidi bool) {
	c.set_link(dest, bidi, false)
}

// Returns true if the cell and @dest are linked,
// false otherwise
func (c *Cell) Linked(dest *Cell) bool {
	for _, v := range c.Linked_cells() {
		if v == dest {
			return true
		}
	}
	return false
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
func (g *Grid) Initialize(linked bool) {
	for y := 0; y < g.H; y++ {
		for x := 0; x < g.W; x++ {
			cell := g.At(x, y)
			cell.n = Link {g.At(x, y - 1), linked && g.At(x, y - 1) != nil}
			cell.s = Link {g.At(x, y + 1), linked && g.At(x, y + 1) != nil}
			cell.e = Link {g.At(x + 1, y), linked && g.At(x + 1, y) != nil}
			cell.w = Link {g.At(x - 1, y), linked && g.At(x - 1, y) != nil}
		}
	}
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

func main() {
	grid := New_grid(8, 8)
	grid.Initialize(true)
	fmt.Println(grid.To_s())
}
