package go_maze_generator

import "math/rand"

type backtracking struct {
	grid *Grid
	cell_stack []*Cell
}

func New_Backtracking() *backtracking {
	b := &backtracking {}
	return b
}

func (b *backtracking) init_for(grid *Grid) {
	b.grid = grid
	b.cell_stack = make([]*Cell, 0, grid.Size())
	b.cell_stack = append(b.cell_stack, grid.At_index(0))
}

func (b *backtracking) is_initialized() bool {
	return b.grid != nil
}

func (b *backtracking) is_over() bool {
	return len(b.cell_stack) == 0
}

func (b *backtracking) perform() {
	for !b.is_over() {
		b.perform_step()
	}
}

func (b *backtracking) perform_step() {
	var current_cell = b.cell_stack[len(b.cell_stack) - 1]
	
	// Get the list of unvisited neighbour cells
	var dests = current_cell.Unlinked_cells()
	for i := 0; i < len(dests); i++ {
		cell := dests[i]
		if len(cell.Linked_cells()) > 0 {
			// A cell has already been visited if it's linked, so we remove them from the list
			dests[i] = dests[len(dests) - 1]
			dests = dests[:len(dests) - 1]
			i--
		}
	}

	if len(dests) > 0 {
		// Now we choose a random dest, carve a way into it and push it on the stack
		next_cell := dests[rand.Intn(len(dests))]
		current_cell.Link(next_cell, true)
		b.cell_stack = append(b.cell_stack, next_cell)
	} else {
		// We can't go anywhere, thus we backtrack
		b.cell_stack = b.cell_stack[:len(b.cell_stack) - 1]
	}
}
