package day2

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"io"
)

type Day2 event.Solvable

func New(path string, output io.Writer) event.Day {
    return Day2{Path: path, Output: output}
}

func (d Day2) Part1() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!\n")
}

func (d Day2) Part2() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!\n")
}
