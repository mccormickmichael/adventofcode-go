package day13

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
	"io"
)

type day13 event.Solvable


func New(path string, output io.Writer) event.Day {
	return day13{Path:path, Output:output}
}

func (d day13) Part1() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	out := make(chan int, 3)

	ic := intcode.Builder(program).WithOutput(out).Build()

	go ic.GoRun()

	blocks := 0
	b := newBoard(38, 20)

	b.read(out)
	for x := 0; x < b.extent.x; x++ {
		for y := 0; y < b.extent.y; y++ {
			if b.tiles[x][y] == Block {
				blocks++
			}
		}
	}

	b.render(d.Output)
	_, _ = fmt.Fprintf(d.Output, "Blocks: %d\n", blocks)
}

// expected extents: x: 38 y: 20

func (d day13) Part2() {

}

func readTile(in chan int) *tile {

	buf := [3]int{}
	count := 0
	for count < 3 {
		select {
		case buf[count] = <- in:
			count ++
		default:
			return nil
		}
	}
	return &tile{coord{buf[0], buf[1]}, buf[2]}
}
