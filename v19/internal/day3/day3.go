package day3

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day3 event.Solvable

type direction string
const (
	Up direction = "U"
	Down direction = "D"
	Left direction = "L"
	Right direction = "R"
)
var dirFrom = map[string]direction{  // TODO: there must be a more idiomatic way of doing this
	"U": Up,
	"D": Down,
	"L": Left,
	"R": Right,
}

type outlay struct {
	direction direction
	distance int
}

var outlayRe = regexp.MustCompile(`([UDLR])(\d+)`)

func parseOutlay(text string) outlay {
	if match := outlayRe.FindStringSubmatch(text); match != nil {
		distance, _ := strconv.Atoi(match[2])
		direction := dirFrom[match[1]]
		return outlay{direction, distance}
	}
	return outlay{}
}

func offset(p event.Point, o outlay) event.Point {
	switch o.direction {
	case Up:
		return p.Offset(0, o.distance)
	case Down:
		return p.Offset(0, -o.distance)
	case Right:
		return p.Offset(o.distance, 0)
	case Left:
		return p.Offset(-o.distance, 0)
	}
	return p
}

func New(path string, output io.Writer) event.Day {
	return Day3{Path: path, Output: output}
}

func (d Day3) Part1() {
	lines := input.Lines(d.Path)
	wire1 := makeSegments(parse(lines[0]))
	wire2 := makeSegments(parse(lines[1]))

	intersections := intersections(wire1, wire2)[1:]
	points := ToPoints(intersections)
	sort.Sort(event.PointSlice(points))
	closest := points[0]

	_, _ = fmt.Fprintf(d.Output, "Closest intersection: %v distance: %d", closest, closest.Magnitude())
}

func (d Day3) Part2() {
	lines := input.Lines(d.Path)
	wire1 := makeSegments(parse(lines[0]))
	wire2 := makeSegments(parse(lines[1]))

	intersections := intersections(wire1, wire2)[1:]
	wireLengths := make([]int, len(intersections))
	for i, x := range intersections {
		length1 := wireLength(wire1, x.wire1Index, x.p)
		length2 := wireLength(wire2, x.wire2Index, x.p)
		wireLengths[i] = length1 + length2
	}
	sort.Sort(sort.IntSlice(wireLengths))

	_, _ = fmt.Fprintf(d.Output, "Shortest wire length: %d\n", wireLengths[0])
}

func parse(input string) []outlay {
	tokens := strings.Split(input, ",")
	outlays := make([]outlay, len(tokens))
	for i, t := range tokens {
		outlays[i] = parseOutlay(t)
	}
	return outlays
}

func makeSegments(outlays []outlay) []event.Segment {
	segments := make([]event.Segment, len(outlays))
	lastPoint := event.Point{}
	for i, o := range outlays {
		nextPoint := offset(lastPoint, o)
		segments[i] = event.Segment{Head: lastPoint, Tail: nextPoint}
		lastPoint = nextPoint
	}
	return segments
}

type intersection struct {
	p event.Point
	wire1Index, wire2Index int
}

func ToPoints(intersections []intersection) []event.Point {
	points := make([]event.Point, len(intersections))
	for i, intersection := range intersections {
		points[i] = intersection.p
	}
	return points
}

func intersections(wire1, wire2 []event.Segment) []intersection {
	intersections := make([]intersection, 0)
	for i := 0; i < len(wire1); i++ {
		for j := 0; j < len(wire2); j++ {
			more := event.Intersections(wire1[i], wire2[j])
			for _, m := range more {
				intersections = append(intersections, intersection{m, i, j})
			}
		}
	}
	return intersections
}

type Wire []event.Segment

func wireLength(w Wire, index int, p event.Point) int {
	length := 0
	for i := 0; i < index; i++ {
		length += w[i].Length()
	}
	length += w[index].Head.Distance(p)
	return length
}