package day18

import mz "github.com/mccormickmichael/adventofcode-go/v19/internal/maze"

type scenario struct {
	maze   *mz.Maze
	loc    mz.Coord
	path   string
	parent *scenario
	outstandingKeys []Key
}
