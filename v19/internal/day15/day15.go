package day15

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
	"log"
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
	inc  := make(chan int, 2)
	outc := make(chan int, 2)
	probe := Droid{theMaze, maze.Coord{0,0},inc, outc}

	ic := intcode.Builder(program).WithInputChannel(inc).WithOutputChannel(outc).Build()
	go ic.GoRun()

	mapper := NewMapper(theMaze, probe)
	mapper.Start(probe.loc)
	for i := 0; i < 200; i++ {
		mapper.Step()
		//theMaze.Render(d.Output)
		head := mapper.explores[len(mapper.explores)-1]
		_, _ = fmt.Fprintf(d.Output, "%v %v\n", head.cell.Loc(), head.dir)
	}
	//mapper.Map()
	theMaze.Render(d.Output)
}

func (d day15) Part2() {
}

type Droid struct {
	maze *maze.Maze
	loc maze.Coord
	input chan int
	output chan int
}

func (p Droid) probe(loc maze.Coord, dir maze.Dir) *maze.Cell {
	p.input <- dirToCommand[dir]
	result := <- p.output
	println(result)
	switch dir {
	case maze.Up:
		loc.Y -= 1
	case maze.Down:
		loc.Y += 1
	case maze.Left:
		loc.X -= 1
	case maze.Right:
		loc.X += 1
	}
	var cellId string
	switch result {
	case 0:
		cellId = "#"
	case 2:
		cellId = "O"
	default:
		cellId = "."
	}
	if current := p.maze.At(loc.X, loc.Y); current != nil {
		return current
	}
	c := maze.NewCell(cellId, p.maze, loc, result != 0)
	if err := p.maze.Set(loc.X, loc.Y, c); err != nil {
		log.Printf("Error setting cell %v at %v: %s\n", c, loc, err)
	}
	return c
}

var dirToCommand = map[maze.Dir]int {
	maze.Up: 1,
	maze.Right: 4,
	maze.Down: 2,
	maze.Left: 3,
}
