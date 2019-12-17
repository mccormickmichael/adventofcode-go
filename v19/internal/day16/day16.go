package day16

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
	"strconv"
	"strings"
)

type day16 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day16{Path:path, Output:output}
}

func (d day16) Part1() {
	digits := input.Digits(strings.TrimSpace(input.SingleLineFile(d.Path)))

	for i := 0; i < 100; i++ {
		digits = nextPhase(digits)
	}

	b := strings.Builder{}
	for i := 0; i < 8; i++ {
		b.WriteString(strconv.Itoa(digits[i]))
	}
	_, _ = fmt.Fprintf(d.Output, "input length: %d\n", len(digits))
	_, _ = fmt.Fprintf(d.Output, "%s\n", b.String())
	
}

func (d day16) Part2() {

	// notes on part 2: I can calculate the next phase in-place
	// because the first N terms are 0 for the N'th value

	// prepend '0' to the list of digits and you don't have to offset the series
	// add and subtract subsets of sequences. Multiplication not necessary
	// skip the zeroes.
}

6500000
