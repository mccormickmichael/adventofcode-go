package day3

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"sort"
	"testing"
)

func TestOutlay(t *testing.T) {
	cases := []struct {
		input string
		expected outlay
	}{
		{"R10", outlay{Right, 10}},
		{"U1", outlay{Up, 1}},
		{"L777", outlay{Left, 777}},
		{"D012", outlay{Down, 12}},
		{"blah", outlay{}},
	}
	for _, c := range cases {
		actual := parseOutlay(c.input)
		if actual != c.expected {
			t.Errorf("parseOutlay(%s) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}

func TestOffset(t *testing.T) {
	origin := event.Point{0, 0}
	cases := []struct {
		outlay outlay
		expected event.Point
	}{
		{ outlay{Right, 10}, event.Point{10, 0} },
		{ outlay{Left, 10}, event.Point{-10, 0} },
		{ outlay{Up, 19}, event.Point{0, 19} },
		{ outlay{Down, 10},  event.Point{0, -10} },
	}
	for _, c := range cases {
		actual := offset(origin, c.outlay)
		if actual != c.expected {
			t.Errorf("offset(%v, %v) == %v, expected %v", origin, c.outlay, actual, c.expected)
		}
	}
}

func TestMakeSegments(t *testing.T) {
	cases := []struct {
		outlays          []outlay
		expectedEndpoint event.Point
	}{
		{[]outlay{outlay{Right, 8}, outlay{Up, 5}, outlay{Left, 5}, outlay{Down, 3}},
			event.Point{3, 2}},
		{[]outlay{outlay{Up, 7}, outlay{Right, 6}, outlay{Down, 4}, outlay{Left, 4}},
			event.Point{2, 3}},
	}
	for _, c := range cases {
		segments := makeSegments(c.outlays)
		lastPoint := segments[len(segments)-1].Tail
		if lastPoint != c.expectedEndpoint {
			t.Errorf("makeSegments(%v) == %v, expected %v", c.outlays, lastPoint, c.expectedEndpoint)
		}
	}
}

func TestIntersections(t *testing.T) {
	cases := []struct {
		p1, p2   string
		expected []event.Point
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", []event.Point{pt(3, 3),pt(6,5)}},
	}
	for _, c := range cases {
		wire1 := makeSegments(parse(c.p1))
		wire2 := makeSegments(parse(c.p2))
		intersections := intersections(wire1, wire2)[1:]
		actualPoints := ToPoints(intersections)
		sort.Sort(event.PointSlice(actualPoints))
		if !event.EqualPoints(actualPoints, c.expected) {
			t.Errorf("%s, %s yielded %v, expected %v", c.p1, c.p2, actualPoints, c.expected)
		}
	}
}

func TestClosest(t *testing.T) {
	cases := []struct {
		p1, p2   string
		expected int
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for _, c := range cases {
		wire1 := makeSegments(parse(c.p1))
		wire2 := makeSegments(parse(c.p2))
		intersections := intersections(wire1, wire2)[1:]
		actualPoints := ToPoints(intersections)
		sort.Sort(event.PointSlice(actualPoints))
		if actualPoints[0].Magnitude() != c.expected {
			t.Errorf("%s, %s yielded %v, expected %v", c.p1, c.p2, actualPoints, c.expected)
		}
	}
}

func TestShortestWires(t *testing.T) {
	cases := []struct {
		p1, p2   string
		expected int
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 30},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}
	for _, c := range cases {
		wire1 := makeSegments(parse(c.p1))
		wire2 := makeSegments(parse(c.p2))
		intersections := intersections(wire1, wire2)[1:]

		wireLengths := make([]int, len(intersections))
		for i, x := range intersections {
			length1 := wireLength(wire1, x.wire1Index, x.p)
			length2 := wireLength(wire2, x.wire2Index, x.p)
			wireLengths[i] = length1 + length2
		}
		sort.Sort(sort.IntSlice(wireLengths))
		if wireLengths[0] != c.expected {
			t.Errorf("%s, %s yielded %v, expected %v", c.p1, c.p2, wireLengths, c.expected)
		}
	}

}

func pt(x, y int) event.Point {
	return event.Point{x, y}
}