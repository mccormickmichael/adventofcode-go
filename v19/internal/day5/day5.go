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
	// TODO: This deadlocks Intcode by pushing too much output.
	// TODO: Move this to a goroutine and report on the last output
	program := input.ParseInts(input.SingleLineFile(d.Path))
	ic := intcode.New(program)
	ic.SetInput(1)
	err := ic.Run()
	if err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected error: %s\n", err)
		return
	}
	_, _ = fmt.Fprintf(d.Output, "Diagnostic Code: %d in %d instructions\n", ic.PopOutput(), ic.Count())
}

func (d Day5) Part2() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	ic := intcode.New(program)
	ic.SetInput(5)
	err := ic.Run()
	if err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected error: %s\n", err)
		return
	}
	_, _ = fmt.Fprintf(d.Output, "Diagnostic Code: %d in %d instructions\n", ic.PopOutput(), ic.Count())
}