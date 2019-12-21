package day15

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
	"log"
)

type Prober interface {
	probe(location maze.Coord, direction maze.Dir) (*maze.Cell, error)
}

type directionFinder interface {
	nextDirection() maze.Dir
}


type exploration struct {
	cell    *maze.Cell
	dir     maze.Dir
	reverse maze.Dir
}

func (e *exploration) nextDirection() maze.Dir {
	neighbors := e.cell.Neighbors()
	nextDir := e.dir
	for t := 0; t < 4; t++ {
		nextDir = nextDir.Next()
		if neighbors[int(nextDir)] == nil {
			return nextDir
		}
	}
	return maze.None
}

type Mapper struct {
	maze     *maze.Maze
	explores []*exploration
	probe    Prober
	steps    int
	maxDepth int
	oxygenDistance int
}

func NewMapper(maze *maze.Maze, probe Prober) *Mapper{
	return &Mapper{maze:maze, probe:probe}
}

func (m *Mapper) Start(start maze.Coord) {
	startCell := maze.NewCell("S", m.maze, start, true)
	err := m.maze.Set(start.X, start.Y, startCell)
	if err != nil {
		log.Printf("Error setting cell %v: %s", startCell, err)
	}
	m.explores = []*exploration{{startCell, maze.Up, maze.None}}
}

func (m *Mapper) Map() error {
	for len(m.explores) > 0 {
		if err := m.Step(); err != nil {
			return fmt.Errorf("step %d: %s", m.steps, err)
		}
		m.steps++
	}
	return nil
}

func (m *Mapper) Step() error {
	lastIndex := len(m.explores)-1
	e := m.explores[lastIndex]

	if e.dir == maze.None {
		e.cell.Explored = true
		m.explores = m.explores[:lastIndex]
		return m.reverse(e.cell.Loc(), e.reverse)
	}

	nextCell, err := m.probe.probe(e.cell.Loc(), e.dir)
	if err != nil {
		return err
	}
	if nextCell.Id() == "O" {
		m.oxygenDistance = len(m.explores)
	}

	if !nextCell.Traversable || nextCell.Explored {
		e.dir = e.nextDirection()
		return nil
	}
	m.explores = append(m.explores, &exploration{nextCell, e.dir, e.dir.Reverse()})
	e.dir = e.nextDirection()
	if len(m.explores) > m.maxDepth {
		m.maxDepth = len(m.explores)
	}
	return nil
}

func (m *Mapper) reverse(loc maze.Coord, reverseDir maze.Dir) error {
	if len(m.explores) == 0 {
		return nil
	}
	last := m.explores[len(m.explores)-1]
	result, err := m.probe.probe(loc, reverseDir)
	if err != nil {
		return fmt.Errorf("reversing %v, %s: %s", loc, reverseDir, err)
	}
	if result.Loc() != last.cell.Loc() {
		return fmt.Errorf("expected reverse to %v but was %v", last.cell.Loc(), result.Loc())
	}
	return nil
}