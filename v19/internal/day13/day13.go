package day13

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	//	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
	"io"
)

type day13 event.Solvable


func New(path string, output io.Writer) event.Day {
	return day13{Path:path, Output:output}
}

func (d day13) Part1() {
	program := input.ParseInts(input.SingleLineFile(d.Path))

	b := newBoard(38, 20)
	b.renderer = d.Output // ick.

	ic := intcode.Builder(program).WithOutputter(b).Build()

	blocks := 0
	ic.Run()

	for x := 0; x < b.extent.x; x++ {
		for y := 0; y < b.extent.y; y++ {
			if b.tiles[x][y] == Block {
				blocks++
			}
		}
	}

	_, _ = fmt.Fprintf(d.Output, "Blocks: %d\n", blocks)
}

// expected extents: x: 38 y: 20

func (d day13) Part2() {
	program := input.ParseInts(input.SingleLineFile(d.Path))

	b := newBoard(38, 20)
	b.renderer = d.Output

	ic := intcode.Builder(program).WithOutputter(b).WithInputter(b).Build()
	ic.Poke(0, 2)
	
	ic.Run()

	if ic.Error() != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected error on turn %d: %s", b.turn, ic.Error())
		b.render()
	}
}

