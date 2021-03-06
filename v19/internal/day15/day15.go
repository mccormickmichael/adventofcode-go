package day15

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
	//	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
	"io"
)

type day15 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day15{Path:path, Output:output}
}

func (d day15) Part1() {
	program := input.ParseInts(input.SingleLineFile(d.Path))

	theMaze := maze.NewMaze(-25, -25, 25, 25)
	inc := make(chan int, 2)
	outc := make(chan int, 2)
	probe := &Droid{theMaze, maze.Coord{0, 0}, maze.Coord{0, 0}, inc, outc}

	ic := intcode.Builder(program).WithInputChannel(inc).WithOutputChannel(outc).Build()
	go ic.GoRun()

	mapper := NewMapper(theMaze, probe)
	mapper.Start(probe.loc)
	err := mapper.Map()
	if err != nil {
		_, _ = fmt.Fprint(d.Output, err)
	}
	theMaze.Render(d.Output)
	_, _ = fmt.Fprintf(d.Output, "Oxygen system at %v\n", probe.oxygen)
	_, _ = fmt.Fprintf(d.Output, "Steps: %d, max depth: %d\n", mapper.steps, mapper.maxDepth)
	_, _ = fmt.Fprintf(d.Output, "Distance to the oxygen system: %d", mapper.oxygenDistance)
}

func (d day15) Part2() {
	program := input.ParseInts(input.SingleLineFile(d.Path))

	theMaze := maze.NewMaze(-25, -25, 25, 25)
	inc := make(chan int, 2)
	outc := make(chan int, 2)
	probe := &Droid{theMaze, maze.Coord{0, 0}, maze.Coord{0, 0}, inc, outc}

	ic := intcode.Builder(program).WithInputChannel(inc).WithOutputChannel(outc).Build()
	go ic.GoRun()

	mapper := NewMapper(theMaze, probe)
	mapper.Start(probe.loc)
	if err := mapper.Map(); err != nil {
		_, _ = fmt.Fprint(d.Output, err)
	}
	ticks, err := flood(theMaze, maze.Coord{18,-18})
	if err != nil {
		_, _ = fmt.Fprint(d.Output, err)
	}
	_, _ = fmt.Fprintf(d.Output, "Ticks to flood with oxygen: %d\n", ticks)
}

type Droid struct {
	maze *maze.Maze
	loc maze.Coord
	oxygen maze.Coord
	input chan int
	output chan int
}

func (p *Droid) probe(loc maze.Coord, dir maze.Dir) (*maze.Cell, error) {

	if loc != p.loc {
		return nil, fmt.Errorf("probe command %s issued for location %v but droid is at %v", dir, loc, p.loc)
	}

	p.input <- dirToCommand[dir]
	result := <- p.output

	switch dir {
	case maze.Up:
		loc.Y -= 1
	case maze.Down:
		loc.Y += 1
	case maze.Left:
		loc.X -= 1
	case maze.Right:
		loc.X += 1
	default:
		return nil, fmt.Errorf("unexpected direction %s at %v", dir, loc)
	}
	var cellId string
	switch result {
	case 0:
		cellId = "#"
	case 2:
		cellId = "O"
		p.loc = loc
		p.oxygen = loc
	default:
		cellId = "."
		p.loc = loc
	}
	if current := p.maze.At(loc); current != nil {
		return current, nil
	}
	c := maze.NewCell(cellId, p.maze, loc, result != 0)
	if err := p.maze.Set(loc, c); err != nil {
		return nil, fmt.Errorf("Error setting cell %v at %v: %s\n", c, loc, err)
	}
	return c, nil
}

var dirToCommand = map[maze.Dir]int {
	maze.Up: 1,
	maze.Right: 4,
	maze.Down: 2,
	maze.Left: 3,
}

func flood(m *maze.Maze, loc maze.Coord) (int, error) {
	start := m.At(loc)
	start.Traversable = false
	front := []*maze.Cell{start}
	ticks := 0
	for len(front) > 0 {
		newFront := make([]*maze.Cell, 0)
		for _, c := range front {
			neighbors := c.Neighbors()
			for _, n := range neighbors {
				if n.Traversable {
					n.Traversable = false
					newFront = append(newFront, n)
				}
			}
		}
		front = newFront
		if len(front) > 0 {
			ticks++
		}
	}
	return ticks, nil
}
