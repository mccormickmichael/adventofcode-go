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
	r.init(1000000000000)
	for _, rx := range input.Lines(d.Path) {
		r.makeReaction(rx)
	}

	amount := 1

	err := r.refine(amount)
	if err != nil {
		_, _ = fmt.Fprintf(d.Output, "Out of resources! %v\n", err)
	}
	o := r.find("ORE")
	_, _ = fmt.Fprintf(d.Output, "ORE required for %d FUEL: %d\n", amount, o.consumed)

	for name, c := range r {
		_, _ = fmt.Fprintf(d.Output, "%7s Produced: %10d, Consumed: %10d, Remain: %10d\n", name, c.produced, c.consumed, c.silo)
	}
}

func (d day14) Part2() {
	r := refinery{}
	r.init(1000000000000)
	for _, rx := range input.Lines(d.Path) {
		r.makeReaction(rx)
	}

	amount := 3000000
	count := 0
	for count < amount {
		err := r.refine(1)
		if err != nil {
			_, _ = fmt.Fprintf(d.Output, "Out of resources! %s\n", err)
			break
		}
		count++
	}

	o := r.find("ORE")
	_, _ = fmt.Fprintf(d.Output, "ORE required for %d FUEL: %d\n", count, o.consumed)

	for name, c := range r {
		_, _ = fmt.Fprintf(d.Output, "%5s Produced: %12d, Consumed: %12d, Remain: %3d\n", name, c.produced, c.consumed, c.silo)
	}
}


