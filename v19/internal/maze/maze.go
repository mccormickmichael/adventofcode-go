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

func (c Coord) String() string {
	return fmt.Sprintf("[%d, %d]", c.X, c.Y)
}

func (c Coord) Up() Coord {
	return c.Move(0, -1)
}

func (c Coord) Right() Coord {
	return c.Move(1, 0)
}

func (c Coord) Down() Coord {
	return c.Move(0, 1)
}

func (c Coord) Left() Coord {
	return c.Move(-1, 0)
}

func (c Coord) Move(x, y int) Coord {
	return Coord{c.X + x, c.Y + y}
}

// Taxicab or Manhattan distance
func (c Coord) Distance(o Coord) int {
	return intmath.Abs(o.X - c.X) + intmath.Abs(o.Y - c.Y)
}

type Locator interface {
	At(loc Coord) *Cell
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

func (m *Maze) At(loc Coord) *Cell {
	xo := loc.X - m.xOffset
	yo := loc.Y - m.yOffset
	if xo < 0 || xo > m.xExtent || yo < 0 || yo > m.yExtent {
		return nil
	}
	return m.cells[xo][yo]
}

func (m *Maze) Set(loc Coord, c *Cell) error {
	xo := loc.X - m.xOffset
	yo := loc.Y - m.yOffset
	if xo < 0 || xo > m.xExtent || yo < 0 || yo > m.yExtent {
		return errors.New(fmt.Sprintf("%v not in maze extent [%d, %d, %d, %d]",
			loc, m.xOffset, m.yOffset, m.xExtent-m.xOffset, m.yExtent-m.yOffset))
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