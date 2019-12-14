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
	out := make(chan int, 3)

	ic := intcode.Builder(program).WithOutputChannel(out).Build()

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
