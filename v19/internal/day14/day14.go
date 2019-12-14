package day14

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
)

type day14 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day14{Path:path, Output:output}
}

func (d day14) Part1() {
	r := refinery{}
	r.init()
	for _, rx := range input.Lines(d.Path) {
		r.makeReaction(rx)
	}
	//
	//for _, c := range r {
	//	_, _ = fmt.Fprintf(d.Output, "%6s %t\n", c.name, c.producer != nil)
	//}

	r.refine()
	o := r.find("ORE")
	_, _ = fmt.Fprintf(d.Output, "ORE required for 1 FUEL: %d\n", o.consumed)

	for name, c := range r {
		_, _ = fmt.Fprintf(d.Output, "%8s Produced: %8d, Consumed: %8d, Remain: %8d\n", name, c.produced, c.consumed, c.silo)
	}
}

func (d day14) Part2() {

}


