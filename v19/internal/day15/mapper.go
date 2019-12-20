package day15

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
	"log"
)

type Prober interface {
	probe(location maze.Coord, direction maze.Dir) *maze.Cell
}

type directionFinder interface {
	nextDirection() maze.Dir
}


type exploration struct {
	cell *maze.Cell
	dir  maze.Dir
}

func (e exploration) nextDirection() maze.Dir {
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
	explores []exploration
	probe    Prober
}

func NewMapper(maze *maze.Maze, probe Prober) *Mapper{
	return &Mapper{maze:maze, probe:probe}
}

func (m *Mapper) Start(start maze.Coord) {
	startCell := maze.NewCell("O", m.maze, start, true)
	err := m.maze.Set(start.X, start.Y, startCell)
	if err != nil {
		log.Printf("Error setting cell %v: %s", startCell, err)
	}
	m.explores = []exploration{{startCell, maze.Up}}
}

func (m *Mapper) Map() {

	for len(m.explores) > 0 {
		m.Step()
	}
}

func (m *Mapper) Step() {
	lastIndex := len(m.explores)-1
	e := m.explores[lastIndex]

	nextCell := m.probe.probe(e.cell.Loc(), e.dir)

	if !nextCell.Traversable || nextCell.Explored {
		nextDir := e.nextDirection()
		if nextDir == maze.None {
			m.explores[lastIndex].cell.Explored = true
			m.explores[lastIndex].cell = nil
			m.explores = m.explores[:lastIndex]
			return
		}
		m.explores[lastIndex].dir = nextDir
		return
	}
	m.explores = append(m.explores, exploration{nextCell, e.dir})
}