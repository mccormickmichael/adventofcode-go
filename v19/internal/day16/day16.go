package day16

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
	"strconv"
	"strings"
	"time"
)

type day16 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day16{Path:path, Output:output}
}

func (d day16) Part1() {
	digits := input.Digits(strings.TrimSpace(input.SingleLineFile(d.Path)))

	msg := make([]int, len(digits)+1)
	copy(msg[1:], digits)
	start := time.Now()

	for i := 0; i < 100; i++ {
		fft(msg)
	}

	elapsed := time.Now().Sub(start)

	b := strings.Builder{}
	for i := 1; i < 9; i++ {
		b.WriteString(strconv.Itoa(msg[i]))
	}
	_, _ = fmt.Fprintf(d.Output, "input length: %d\n", len(digits))
	_, _ = fmt.Fprintf(d.Output, "elapsed time: %f\n", float64(elapsed/1000)/1000.0)
	_, _ = fmt.Fprintf(d.Output, "%s\n", b.String())
	
}

func (d day16) Part2() {

	digits := input.Digits(strings.TrimSpace(input.SingleLineFile(d.Path)))

	start := time.Now()
	msg := make([]int, len(digits)*10000 + 1)
	for i := 0; i < 10000; i++ {
		copy(msg[i*len(digits)+1:], digits)
	}
	elapsed := time.Now().Sub(start)

	_, _ = fmt.Fprintf(d.Output, "created the message in %f ms\n", float64(elapsed/1000)/1000.0)

	start = time.Now()
	
	for i := 0; i < 100; i++ {
		fft(msg)
		_, _ = fmt.Fprintf(d.Output, "phase %d\n", i+1)
	}

	elapsed = time.Now().Sub(start)

	_, _ = fmt.Fprintf(d.Output, "computed phases in %f ms\n", float64(elapsed/1000)/1000.0)

	// notes on part 2: I can calculate the next phase in-place
	// because the first N terms are 0 for the N'th value

	// prepend '0' to the list of digits and you don't have to offset the series
	// add and subtract subsets of sequences. Multiplication not necessary
	// skip the zeroes.
}

// 6500000
