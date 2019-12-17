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

	output := makeOutput(msg[1:])

	_, _ = fmt.Fprintf(d.Output, "input length: %d\n", len(digits))
	_, _ = fmt.Fprintf(d.Output, "elapsed time: %f\n", float64(elapsed/1000)/1000.0)
	_, _ = fmt.Fprintf(d.Output, "%s\n", output)
	
}

func (d day16) Part2() {

	line := input.SingleLineFile(d.Path)
	offset, _ := strconv.Atoi(line[:7])
	digits := input.Digits(strings.TrimSpace(line))

	start := time.Now()
	msg := make([]int, len(digits)*10000 + 1)
	for i := 0; i < 10000; i++ {
		copy(msg[i*len(digits)+1:], digits)
	}
	elapsed := time.Now().Sub(start)

	_, _ = fmt.Fprintf(d.Output, "created the message in %f ms\n", float64(elapsed/1000)/1000.0)

	start = time.Now()
	
	for i := 0; i < 100; i++ {
		fft2(msg[offset:])
	}

	elapsed = time.Now().Sub(start)

	_, _ = fmt.Fprintf(d.Output, "computed phases in %.3f ms\n", float64(elapsed/1000)/1000.0)

	output := makeOutput(msg[offset+1:])

	_, _ = fmt.Fprintf(d.Output, "message at offset %d is %s\n", offset, output)
}

func makeOutput(digits []int) string {
	b := strings.Builder{}
	for i := 0; i < 8; i++ {
		b.WriteString(strconv.Itoa(digits[i]))
	}
	return b.String()
}


// 6500000
