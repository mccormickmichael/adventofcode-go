package day12

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
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
	maxSteps := 1
	for i := 1; i <= maxSteps; i++ {
		accelerate(moons)
		velocitate(moons)
	}
	printstate(d.Output, maxSteps, moons)
}

func (d day12) Part2() {
	moons := parse(input.Lines(d.Path))
	_,_ = fmt.Fprint(d.Output, "X Axis:\n")
	x := [4]int{moons[0].pos.x, moons[1].pos.x, moons[2].pos.x, moons[3].pos.x}
	findCycle(x, [4]int{0,0,0,0}, d.Output)

	_,_ = fmt.Fprint(d.Output, "\nY Axis:\n")
	y := [4]int{moons[0].pos.y, moons[1].pos.y, moons[2].pos.y, moons[3].pos.y}
	findCycle(y, [4]int{0,0,0,0}, d.Output)

	_,_ = fmt.Fprint(d.Output, "\nZ Axis:\n")
	z := [4]int{moons[0].pos.z, moons[1].pos.z, moons[2].pos.z, moons[3].pos.z}
	findCycle(z, [4]int{0,0,0,0}, d.Output)

	a := 286332
	b := 161428
	c := 102356

	ab := intmath.Lcm(a,b)
	abc := intmath.Lcm(ab, c)
	_, _ = fmt.Fprintf(d.Output, "LCM of %d, %d, %d : %d\n",	a, b, c, abc)
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
	_, _ = fmt.Fprintf(o, "Step %d:\n", step)
	energy := 0
	for _, m := range moons {
		_, _ = fmt.Fprintf(o, "%s energy: %d\n", m, m.energy())
		energy += m.energy()
	}
	_, _ = fmt.Fprintf(o, "total energy: %d\n", energy)
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
