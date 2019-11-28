package day3

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v18/internal/event"
	"io"
	"regexp"
	"strconv"
)

type Day3 event.Solvable

var claimRe = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

type claim struct {
	id,	x, y, width, height uint16
}

func New(path string, output io.Writer) event.Day {
	return Day3{Path: path, Output: output}
}

func (d Day3) Part1() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!")
}

func (d Day3) Part2() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!")
}

func parseClaim(line string) claim {
	matches := claimRe.FindStringSubmatch(line)

	id, _     := strconv.Atoi(matches[1])
	x, _      := strconv.Atoi(matches[2])
	y, _      := strconv.Atoi(matches[3])
	width, _  := strconv.Atoi(matches[4])
	height, _ := strconv.Atoi(matches[5])
		
	return claim{
		uint16(id),
		uint16(x),
		uint16(y),
		uint16(width),
		uint16(height),
	}
}
