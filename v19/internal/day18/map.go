package day18

import "github.com/mccormickmichael/adventofcode-go/v19/internal/maze"

type Door struct {
	name string
	loc  maze.Coord
}

type Key struct {
	name string
	loc  maze.Coord
	// door Door  ?? pointer?
}

func makeDoor(name string, loc maze.Coord) Door {
	return Door{name, loc}
}

func makeKey(name string, loc maze.Coord) Key {
	return Key{name, loc}
}