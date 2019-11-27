package day3

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v18/internal/event"
	"io"
)

type Day3 event.Solvable

func New(path string, output io.Writer) event.Day {
	return Day3{Path: path, Output: output}
}

func (d Day3) Part1() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!")
}

func (d Day3) Part2() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!")
}	
