package go_maze_generator

type Algorithm interface {
	init_for(*Grid)
	is_initialized() bool
	is_over() bool
	perform()
	perform_step()
}
