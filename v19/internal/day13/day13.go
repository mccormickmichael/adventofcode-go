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
	// program := input.ParseInts(input.SingleLineFile(d.Path))
	// out := make(chan int, 3)
	// in  := make(chan int)

	// ic := intcode.Builder(program).WithOutput(out).WithInput(in).Build()
	// ic.Poke(0, 2)
	
	// go ic.GoRun()

	// b := newBoard(38, 20)

	// b.read(out)
	// b.render(d.Output)
	// _, _ = fmt.Fprintf(d.Output, "Score: %d\n", b.score)

	// in <- intmath.Cmp(b.ball.x, b.paddle.x)

	// for x := 0; x < b.extent.x; x++ {
	// 	for y := 0; y < b.extent.y; y++ {
	// 		_, _ = fmt.Fprintf(d.Output, "%v ", readTile(out))
	// 	}
	// }
	// 	b.read(out)
	// b.render(d.Output)
	// _, _ = fmt.Fprintf(d.Output, "Score: %d\n", b.score)

}

