package day10

import (
	"fmt"
	"io"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
)


type day10 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day10{Path:path, Output:output}
}


func (d day10) Part1() {
	asteroids := parse(input.Lines(d.Path))
	_, _ = fmt.Fprintf(d.Output, "There are %d asteroids\n", len(asteroids))
}

func (d day10) Part2() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!\n")
}


func parse(lines []string) []asteroid {
	asteroids := make([]asteroid, 0)
	for y, line := range lines {
		for x, b := range []byte(line) {
			if b == '#' {
				asteroids = append(asteroids, asteroid{x, y})
			}
		}
	}
	return asteroids
}

func unblocked(roids []asteroid) []lineOfSight {

	unblocked := make([]lineOfSight, 0)

	for i := 0; i < len(roids)-1; i++ {
		for j := i+1; j < len(roids); j++ {
			los := lineOfSight{roids[i], roids[j]}
			if !los.blocked(roids) {
				unblocked := append(unblocked, los)
			}
		}
	}
	return unblocked
}

type asteroid struct {
	x, y int
}

func (c asteroid) blocks(a, b asteroid) bool {
	abdist := a.dist(b)
	acdist := a.dist(c)
	bcdist := b.dist(c)

	if abdist != acdist + bcdist {
		return false
	}
	return a.equalSlopes(b, c)
}

func (c asteroid) dist(a asteroid) int {
	return intmath.Abs(a.x - b.x) + intmath.Abs(a.y - b.y)
}

const epsilon = 0.001

func (c asteroid) equalSlopes(a, b asteroid) bool {
	// vertically aligned
	if a.x-b.x == 0 {
		return a.x-c.x == 0
	}
	// horizontally aligned
	if a.y-b.y == 0 {
		return a.y-c.y == 0
	}
	return math.Abs(a.slope(b) - a.slope(c)) < epsilon
}

func (c asteroid) slope(o asteroid) float64 {
	return (c.y-o.y)/float64(c.x-o.x)
}

type lineOfSight struct {
	a, b asteroid
}


func (los lineOfSight) blocked(roids []asteroid) bool {
	for _, r := range roids {
		if r.blocks(los.a, los.b) {
			return true
		}
	}
	return false
}
