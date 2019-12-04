package day4

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"io"
	"strconv"
)

type Day4 struct {
	Min, Max int
	Output io.Writer
}

func New(min, max int, output io.Writer) event.Day {
	return Day4{Min: min, Max: max, Output: output}
}

func (d Day4) Part1() {
	count := 0
	for i := d.Min; i <= d.Max; i++ {
		sv := []byte(strconv.Itoa(i))
		if hasRepeats(sv) && increases(sv) {
			count++
		}
	}

	_, _ = fmt.Fprintf(d.Output, "Possible passwords: %d\n", count)
}

func (d Day4) Part2() {
	count := 0
	for i := d.Min; i <= d.Max; i++ {
		sv := []byte(strconv.Itoa(i))
		if hasRepeats(sv) && increases(sv) && hasDouble(sv) {
			count++
		}
	}
	_, _ = fmt.Fprintf(d.Output, "Possible passwords! %d\n", count)
}

func hasRepeats(digits []byte) bool {
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			return true
		}
	}
	return false
}

func increases(digits []byte) bool {
	max := digits[0]
	for i := 1; i < len(digits); i++ {
		if digits[i] < max {
			return false
		}
		max = digits[i]
	}
	return true
}

func hasDouble(digits []byte) bool {
	histo := make(map[byte]int)
	for i := 0; i < len(digits); i++ {
		histo[digits[i]] = histo[digits[i]] + 1
	}
	for _, count := range histo {
		if count == 2 {
			return true
		}
	}
	return false
}