package day18

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
)

type day18 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day18{Path:path, Output:output}
}

func (d day18) Part1() {
	mapper := newMapper()
	theMaze := mapper.build(input.Lines(d.Path))

	theMaze.Render(d.Output)
	_, _ = fmt.Fprintf(d.Output, "Entrance: %s\n", mapper.entrance)
	for name, door := range mapper.doors {
		_, _ = fmt.Fprintf(d.Output, "door[%s]: %s\n", name, door.loc)
	}
	for name, key := range mapper.keys {
		_, _ = fmt.Fprintf(d.Output, " key[%s]: %s\n", name, key.loc)
	}

}

func (d day18) Part2() {

}