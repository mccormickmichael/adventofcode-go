package day15

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
)

func Tiny(m *maze.Maze) *TestMap {
	return NewTestMap(m, []string{
		"#####",
		"#...#",
		"##.##",
		"##.##",
		"#####",
	}, maze.Coord{2, 3})
}

type TestMap struct {
	maze *maze.Maze
	cells [][]*maze.Cell
	Start maze.Coord
}

func NewTestMap(m *maze.Maze, layout []string, start maze.Coord) *TestMap {
	cells := make([][]*maze.Cell, len(layout[0]))
	for c := 0; c < len(cells); c++ {
		cells[c] = make([]*maze.Cell, len(layout))
	}
	for y, row := range layout {
		for x, b := range []byte(row) {
			traversable := true
			if b == '#' {
				traversable = false
			}
			cells[x][y] = maze.NewCell(string(b), m, maze.Coord{x, y}, traversable)
		}
	}
	return &TestMap{m,cells, start}
}

func (m *TestMap) probe(loc maze.Coord, dir maze.Dir) (*maze.Cell, error) {
	switch dir {
	case maze.Up: loc.Y-=1
	case maze.Down: loc.Y +=1
	case maze.Left: loc.X-=1
	case maze.Right: loc.X+=1
	}
	c := m.cells[loc.X][loc.Y]
	err := m.maze.Set(loc, c)
	if err != nil {
		return nil, fmt.Errorf("error setting cell %v: %s", c, err)
	}
	return c, nil
}