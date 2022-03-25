package go_maze_generator

type Algorithm interface {
	on(*Grid)
	step_on(*Grid)
}

// ----------------

type Sidewinder struct {

}

func (s Sidewinder) on(grid *Grid) {

}

func (s Sidewinder) step_on(grid *Grid) {

}

// -----------------

type BinaryTree struct {
	current_cell Cell
}

func (s BinaryTree) on(grid *Grid) {

}

func (s BinaryTree) step_on(grid *Grid) {

}
