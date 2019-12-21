package day18

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
)

type mapper struct {
	entrance maze.Coord
	doors    map[string]Door
	keys     map[string]Key
}

func newMapper() mapper {
	return mapper{doors: make(map[string]Door), keys: make(map[string]Key)}
}


func (m *mapper) build(layout []string) *maze.Maze {

	theMaze := maze.NewMaze(0, 0, len(layout), len(layout[0]))
	for y, row := range layout {
		for x, b := range []byte(row) {
			loc := maze.Coord{x, y}
			cell, err := m.parseCell(theMaze, loc, b)
			if err != nil {
				// TODO: HANDLE
			}
			_ = theMaze.Set(loc, cell)
		}
	}
	return theMaze
}

func (m *mapper) parseCell(mz *maze.Maze, loc maze.Coord, b byte) (*maze.Cell, error) {
	switch {
	case b == '@':
		m.entrance = loc
		return maze.NewCell("@", mz, loc, true), nil
	case b == '#':
		return maze.NewCell("#", mz, loc, false), nil
	case b == '.':
		return maze.NewCell(".", mz, loc, true), nil
	case b >= 65 && b <= 90:
		name := string(b)
		m.doors[name] = makeDoor(name, loc)
		return maze.NewCell(string(b), mz, loc, false), nil
	case b >= 97 && b <= 122:
		name := string(b)
		m.keys[name] = makeKey(name, loc)
		return maze.NewCell(name, mz, loc, true), nil
	default:
		return nil, fmt.Errorf("unknown cell type %b", b)
	}
}