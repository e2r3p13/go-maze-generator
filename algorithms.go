package go_maze_generator

import "math/rand"

type Algorithm interface {
	init_for(*Grid)
	is_initialized() bool
	is_over() bool
	perform()
	perform_step()
}

// ----------------

type Sidewinder struct {

}

func (s Sidewinder) on(grid *Grid) {

}

func (s Sidewinder) step_on(grid *Grid) {

}

// -----------------

type binaryTree struct {
	grid *Grid
	current_cell_index int
}

func New_BinaryTree() *binaryTree {
	bt := &binaryTree {}
	return bt
}

func (b *binaryTree) init_for(grid *Grid) {
	b.grid = grid
	b.current_cell_index = 0 
}

func (b *binaryTree) is_initialized() bool {
	return b.grid != nil
}

func (b *binaryTree) is_over() bool {
	return b.current_cell_index == b.grid.Size()
}

func (b *binaryTree) perform() {
	for !b.is_over() {
		b.perform_step()
	}
}

func (b *binaryTree) perform_step() {
	cell := b.grid.At_index(b.current_cell_index)	
	dests := make([]*Cell, 0, 2)
	
	if cell.n.dest != nil { dests = append(dests, cell.n.dest) }
	if cell.e.dest != nil { dests = append(dests, cell.e.dest) }

	if len(dests) > 0 {
		choice := rand.Intn(len(dests))	
		cell.Link(dests[choice], true)
	}

	b.current_cell_index += 1
}
