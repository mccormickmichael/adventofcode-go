package day18

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
	"math"
)

type Door struct {
	name string
	loc  maze.Coord
}

type Key struct {
	name string
	loc  maze.Coord
	door Door
}

func makeDoor(name string, loc maze.Coord) Door {
	return Door{name, loc}
}

func makeKey(name string, loc maze.Coord) Key {
	return Key{name:name, loc:loc}
}

type Vault struct {
	maze *maze.Maze
	uncollectedKeys []Key
	doors []Door
}

func (v *Vault) closeDoors() {
	for _, d := range v.doors {
		v.maze.At(d.loc).Traversable = false
	}
}

func (v *Vault) openDoor(d Door) {
	v.maze.At(d.loc).Traversable = true
}

type Elf struct {
	loc maze.Coord
	keys []Key
}

type PathSolver struct {
	scenarioStack  []*scenario
	leastSteps     int
	leastStepPaths []string
}

func newPathSolver(root *scenario) *PathSolver {
	return &PathSolver{
		scenarioStack:[]*scenario{root},
		leastSteps:math.MaxInt64,
		leastStepPaths:[]string{},
	}
}

func (ps *PathSolver) remaining() int {
	return len(ps.scenarioStack)
}

func (ps *PathSolver) push(s *scenario) {
	ps.scenarioStack = append(ps.scenarioStack, s)
}

func (ps *PathSolver) pop() *scenario {
	lastIndex := len(ps.scenarioStack)-1
	last := ps.scenarioStack[lastIndex]
	ps.scenarioStack[lastIndex] = nil
	ps.scenarioStack = ps.scenarioStack[:lastIndex]
	return last
}

func (ps *PathSolver) minSteps() int {
	return ps.leastSteps
}

func (ps *PathSolver) setPathSteps(path string, steps int) {
	if steps < ps.leastSteps {
		ps.leastSteps = steps
		ps.leastStepPaths = []string{path}
	} else if steps == ps.leastSteps {
		ps.leastStepPaths = append(ps.leastStepPaths, path)
	}
}
