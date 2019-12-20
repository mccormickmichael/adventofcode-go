package day15

import "github.com/mccormickmichael/adventofcode-go/v19/internal/maze"

type prober interface {
	probe(location maze.Coord, direction maze.Dir) *maze.Cell
}


type exploration struct {
	cell    *maze.Cell
	lastDir maze.Dir
}

func (e exploration) nextNilNeighbor() exploration {
	neighbors := e.cell.Neighbors()
	dir := e.lastDir
	for t := 0; t < 4; t++ {
		dir = dir.Next()
		c := neighbors[int(dir)]
		if c == nil {
			return exploration{c, dir}
		}
	}
	return exploration{nil, maze.Up}
}

type mapper struct {
	maze     *maze.Maze
	explores []exploration
	probe    prober
}

func (m *mapper) Step() {

	e := m.explores[len(m.explores)-1]
	nextCell := m.probe.probe(e.cell.Loc(), e.lastDir)

	if !nextCell.Traversable {
		exp := e.nextNilNeighbor()
		if exp.cell == nil {
			m.explores = m.explores[:len(m.explores)-1]
		} else {
			m.explores = append(m.explores, exp)
		}
	}
}