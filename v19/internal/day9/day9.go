package day9

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"io"
)

type Day9 event.Solvable

func New(path string, output io.Writer) event.Day {
	return Day9{Path:path, Output:output}
}

func (d Day9) Part1() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	oc := make(chan int)

	ic := intcode.Builder(program).WithOutputChannel(oc).WithInputValue(1).Build()
	go ic.GoRun()

	for o := range oc {
		_, _ = fmt.Fprintf(d.Output, "Keycode: %d in %d instructions with memory %d\n", o, ic.Count(), ic.Len())
	}
	if err := ic.Error(); err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected err %v\n", err)
	}
}

func (d Day9) Part2() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	oc := make(chan int)

	ic := intcode.Builder(program).WithOutputChannel(oc).WithInputValue(2).Build()
	go ic.GoRun()

	for o := range oc {
		_, _ = fmt.Fprintf(d.Output, "Coordinates: %d in %d instructions with memory %d\n", o, ic.Count(), ic.Len())
	}
	if err := ic.Error(); err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected err %v\n", err)
	}
	_, _ = fmt.Fprintf(d.Output, "done?\n")
}
