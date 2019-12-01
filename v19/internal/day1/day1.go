package day1

import (
    "fmt"
    "io"

    "github.com/mccormickmichael/adventofcode-go/v19/internal/event"
)

type Day1 event.Solvable

func New(path string, output io.Writer) event.Day {
    return Day1{Path: path, Output: output}
}

func (d Day1) Part1() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!")
}

func (d Day1) Part2() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!")
}
