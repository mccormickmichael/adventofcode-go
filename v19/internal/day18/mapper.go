package day18

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
)

type mapper struct {
	entrance   maze.Coord
	foundDoors map[string]Door
	foundKeys  map[string]Key
}

func newMapper() mapper {
	return mapper{foundDoors: make(map[string]Door), foundKeys: make(map[string]Key)}
}

func (m *mapper) keys() []Key {
	keys := make([]Key, 0, len(m.foundKeys))
	for _, k := range m.foundKeys {
		keys = append(keys, k)
	}
	return keys
}

func (m *mapper) doors() []Door {
	doors := make([]Door, 0, len(m.foundDoors))
	for _, k := range m.foundDoors {
		doors = append(doors, k)
	}
	return doors
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
		m.foundDoors[name] = makeDoor(name, loc)
		return maze.NewCell(string(b), mz, loc, false), nil
	case b >= 97 && b <= 122:
		name := string(b)
		m.foundKeys[name] = makeKey(name, loc)
		return maze.NewCell(name, mz, loc, true), nil
	default:
		return nil, fmt.Errorf("unknown cell type %b", b)
	}
}