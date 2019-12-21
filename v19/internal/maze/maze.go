package maze

import (
	"errors"
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
	"io"
	"strings"
)

type Dir int

const (
	None  Dir = -1
	Up    Dir = 0
	Right Dir = 1
	Down  Dir = 2
	Left  Dir = 3
)

func (d Dir) Next() Dir {
	if d == None {
		return None
	}
	return (d + 1) % 4
}

func (d Dir) Reverse() Dir {
	if d == None {
		return None
	}
	return (d + 2) % 4
}

func (d Dir) String() string {
	switch d {
	case None: return "None"
	case Up:   return "Up"
	case Right: return "Right"
	case Down: return "Down"
	case Left: return "Left"
	}
	return fmt.Sprintf("%d", d)
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

type Celler interface {
	At(x, y int) *Cell
}

func NewMaze(top, left, bottom, right int) *Maze {
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
	xo := x - m.xOffset
	yo := y - m.yOffset
	if xo < 0 || xo > m.xExtent || yo < 0 || yo > m.yExtent {
		return nil
	}
	return m.cells[xo][yo]
}

func (m *Maze) Set(x, y int, c *Cell) error {
	xo := x - m.xOffset
	yo := y - m.yOffset
	if xo < 0 || xo > m.xExtent || yo < 0 || yo > m.yExtent {
		return errors.New(fmt.Sprintf("[%d, %d] not in maze extent [%d, %d, %d, %d]",
			x, y, m.xOffset, m.yOffset, m.xExtent-m.xOffset, m.yExtent-m.yOffset))
	}
	m.cells[xo][yo] = c

	return nil
}

func (m *Maze) Render(o io.Writer) {
	_, _ = fmt.Fprintf(o, "top: %d left %d\n", m.xOffset, m.yOffset)

	for y := 0; y < m.yExtent; y++ {
		b := strings.Builder{}
		for x := 0; x < m.xExtent; x++ {
			b.WriteString(m.cells[x][y].String())
		}
		_, _ = fmt.Fprintln(o, b.String())
	}
}