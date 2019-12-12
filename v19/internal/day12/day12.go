package day12

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
	"regexp"
	"strconv"
)

type day12 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day12{Path:path, Output:output}
}

func (d day12) Part1() {
	moons := parse(input.Lines(d.Path))

	printstate(d.Output, 0, moons)
	
	maxSteps := 1000

	for i := 1; i <= maxSteps; i++ {
		accelerate(moons)
		velocitate(moons)
		//printstate(d.Output, i, moons)
	}
	printstate(d.Output, maxSteps, moons)
}

func (d day12) Part2() {
}

func accelerate(moons []*moon) {
	for _, m := range moons {
		m.accelerate(moons)
	}
}

func velocitate(moons []*moon) {
	for _, m := range moons {
		m.velocitate()
	}
}

func printstate(o io.Writer, step int, moons []*moon) {
	fmt.Fprintf(o, "Step %d:\n", step)
	energy := 0
	for _, m := range moons {
		fmt.Fprintf(o, "%s energy: %d\n", m, m.energy())
		energy += m.energy()
	}
	fmt.Fprintf(o, "total energy: %d\n", energy)
}

var posRe = regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)

func parse(lines []string) []*moon {

	moons := make([]*moon, len(lines))
	for i, line := range lines {

		if match := posRe.FindStringSubmatch(line); match != nil {
			// TODO: capture errors and return them! Remember Day 5!
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			z, _ := strconv.Atoi(match[3])
			pos := position{x, y, z}
			moons[i] = &moon{pos:pos}
		}
	}
	return moons
}
