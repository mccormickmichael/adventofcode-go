package day5

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"io"
)

type Day5 event.Solvable

func New(path string, output io.Writer) event.Day {
	return Day5{Path:path, Output:output}
}

func (d Day5) Part1() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	o := make(chan int)
	ic := intcode.Builder(program).WithOutput(o).Build()
	ic.SetInput(1)

	var output int
	go ic.GoRun()
	for output = range o {}

	if err := ic.Error(); err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected error: %s\n", err)
		return
	}
	_, _ = fmt.Fprintf(d.Output, "Diagnostic Code: %d in %d instructions\n", output, ic.Count())
}

func (d Day5) Part2() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	o := make(chan int)
	ic := intcode.Builder(program).WithOutput(o).Build()
	ic.SetInput(5)

	var output int
	go ic.GoRun()
	for output = range o {}

	if err := ic.Error(); err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected error: %s\n", err)
		return
	}
	_, _ = fmt.Fprintf(d.Output, "Diagnostic Code: %d in %d instructions\n", output, ic.Count())
}