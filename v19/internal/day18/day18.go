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
	//_, _ = fmt.Fprintf(d.Output, "Entrance: %s\n", mapper.entrance)
	//for name, door := range mapper.doors() {
	//	_, _ = fmt.Fprintf(d.Output, "door[%s]: %s\n", name, door.loc)
	//}
	_, _ = fmt.Fprintln(d.Output)
	for _, key := range mapper.keys() {
		_, _ = fmt.Fprintf(d.Output, " key[%s]: %s\n", key.name, key.loc)
	}
	_, _ = fmt.Fprintln(d.Output)


	rootScenario := &scenario{maze:theMaze, loc:mapper.entrance, outstandingKeys:mapper.keys()}
	goals := newGoals(theMaze, rootScenario)
	goals.find()
	_, _ = fmt.Fprintln(d.Output, "Goals:")
	for _, g := range goals.found {
		_, _ = fmt.Fprintf(d.Output, "  %s%s -> %d\n", g.key.name, g.key.loc, g.dist)
	}
}

func (d day18) Part2() {

}