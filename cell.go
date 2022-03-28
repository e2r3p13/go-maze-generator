package go_maze_generator

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
		if l.linked == linked && l.dest != nil {
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
