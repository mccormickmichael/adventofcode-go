package maze

import (
	"errors"
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
)

type Dir int

const (
	Up    Dir = 0
	Right Dir = 1
	Down  Dir = 2
	Left  Dir = 3
)

func (d Dir) Next() Dir {
	return (d + 1) % 4
}

type Coord struct {
	X, Y int
}

func (c Coord) Move(x, y int) Coord {
	return Coord{c.X + x, c.Y + y}
}

// Taxicab or Manhattan distance
func (c Coord) Distance(o Coord) int {
	return intmath.Abs(o.X - c.X) + intmath.Abs(o.Y - c.Y)
}

func New(top, left, bottom, right int) *Maze {
	xExtent := right-left
	yExtent := bottom-top
	maze := &Maze{xOffset:left, yOffset:top, xExtent:xExtent, yExtent:yExtent}
	maze.cells = make([][]*Cell, xExtent)
	for x := 0; x < len(maze.cells); x++ {
		maze.cells[x] = make([]*Cell, yExtent)
	}
	return maze
}

type Maze struct {
	xOffset, yOffset int
	xExtent, yExtent int
	cells [][]*Cell
}

func (m *Maze) At(x, y int) *Cell {
	if x < 0 || x > m.xExtent || y < 0 || y > m.yExtent {
		return nil
	}
	xo := x - m.xOffset
	yo := y - m.yOffset
	return m.cells[xo][yo]
}

func (m *Maze) Set(x, y int, c *Cell) error {
	if x < 0 || x > m.xExtent || y < 0 || y > m.yExtent {
		return errors.New(fmt.Sprintf("[%d, %d] not in maze extent [%d, %d, %d, %d]",
			x, y, m.xOffset, m.yOffset, m.xExtent-m.xOffset, m.yExtent-m.yOffset))
	}
	xo := x - m.xOffset
	yo := y - m.yOffset
	m.cells[xo][yo] = c

	return nil
}