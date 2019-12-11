package day10

import (
	"fmt"
	"io"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
	"math"
	"strings"
)


type day10 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day10{Path:path, Output:output}
}


func (d day10) Part1() {
	asteroids := parse(input.Lines(d.Path))
	_, _ = fmt.Fprintf(d.Output, "There are %d asteroids\n", len(asteroids))
	unblocked := unblocked(asteroids)

	var maxVisible int
	var bestAsteroid asteroid

	for _, r := range asteroids {
		visible := countUnblockedFor(r, unblocked)
		if visible > maxVisible {
			maxVisible = visible
			bestAsteroid = r
		}
	}
	_, _ = fmt.Fprintf(d.Output, "Best asteroid is %v with %d visible\n", bestAsteroid, maxVisible)
}

func (d day10) Part2() {
	//origin := asteroid{22, 28}
	origin := asteroid{22,28}
	asteroids := parse(input.Lines(d.Path))
	polars := make([]polar, 0)
	for _, a := range asteroids {
		if !a.equals(origin) {
			polars = append(polars, toPolar(origin, a))
		}
	}
	targets := collect(polars)

	maxCount := 200
	var zappedAsteroid asteroid
	zappedCount := 0
	index := 0
	for zappedCount < len(polars) && zappedCount < maxCount {
		targetIndex := index % len(targets)
		radialTarget := targets[targetIndex]
		if len(radialTarget) > 0 {
			zappedCount++
			p := radialTarget[0]
			zappedAsteroid = p.asteroid
			_, _ = fmt.Fprintf(d.Output, "%3d Zapped %v at angle %s\n", zappedCount, zappedAsteroid, p.thetaRep)
			targets[targetIndex] = radialTarget[1:]
		}
		index++
	}
}


func parse(lines []string) []asteroid {
	asteroids := make([]asteroid, 0)
	for y, line := range lines {
		for x, b := range []byte(strings.TrimSpace(line)) {
			if b == '#' {
				asteroids = append(asteroids, asteroid{x, y})
			}
		}
	}
	return asteroids
}

func countUnblockedFor(roid asteroid, unblocked []lineOfSight) int {
	count := 0
	for _, los := range unblocked {
		if roid.equals(los.a) || roid.equals(los.b)  {
			count++
		}
	}
	return count
}

func unblocked(roids []asteroid) []lineOfSight {

	unblocked := make([]lineOfSight, 0)

	for i := 0; i < len(roids)-1; i++ {
		for j := i+1; j < len(roids); j++ {
			los := lineOfSight{roids[i], roids[j]}
			if !los.blocked(roids) {
				unblocked = append(unblocked, los)
			}
		}
	}
	return unblocked
}

type asteroid struct {
	x, y int
}

func (c asteroid) equals(o asteroid) bool {
	return c.x == o.x && c.y == o.y
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
	return intmath.Abs(a.x - c.x) + intmath.Abs(a.y - c.y)
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
	return float64(c.y-o.y)/float64(c.x-o.x)
}

type lineOfSight struct {
	a, b asteroid
}


func (los lineOfSight) blocked(roids []asteroid) bool {
	for _, r := range roids {
		if r.equals(los.a) || r.equals(los.b) {
			continue
		}
		if r.blocks(los.a, los.b) {
			return true
		}
	}
	return false
}
