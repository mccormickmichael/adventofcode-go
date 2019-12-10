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

	ic := intcode.Builder(program).WithOutput(oc).Build()
	ic.SetInput(1)
	go ic.GoRun()

	for o := range oc {
		_, _ = fmt.Fprintf(d.Output, "Keycode: %d\n", o)
	}
	if err := ic.Error(); err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected err %v\n", err)
	}
}

func (d Day9) Part2() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	oc := make(chan int)

	ic := intcode.Builder(program).WithOutput(oc).Build()
	ic.SetInput(2)
	go ic.GoRun()

	for o := range oc {
		_, _ = fmt.Fprintf(d.Output, "Coordinate: %d\n", o)
	}
	if err := ic.Error(); err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected err %v\n", err)
	}
	_, _ = fmt.Fprintf(d.Output, "done?\n")
}