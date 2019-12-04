package event

import "testing"

func TestNonIntersection(t *testing.T) {
	cases := []struct {
		a, b Segment
	}{
		{Segment{Point{}, Point{5, 0}}, Segment{Point{3, 4}, Point{3, 1}}},
	}
	for _, c := range cases {
		intersections := Intersections(c.a, c.b)
		if intersections != nil {
			t.Errorf("%v, %v should not intersect, but got %v", c.a, c.b, intersections)
		}
	}
}

func TestOrthogonalIntersection(t *testing.T) {
	cases := []struct {
		a, b Segment
		intersection Point
	}{
		{seg(0, 0, 5, 0), seg(2, 2, 2, -2), pt(2,0)},
		{seg(5, 0, 2, 0), seg(2, 2, 2, 0), pt(2, 0)},
		{seg(0, 0, 0, 4), seg(2, 2, -1, 2), pt(0, 2)},
	}
	for _, c := range cases {
		intersections := Intersections(c.a, c.b)
		if len(intersections) != 1 {
			t.Errorf("%v, %v, intersect at %v, expected %v", c.a, c.b, intersections, c.intersection)
		}
	}
}

func TestHorizontalIntersection(t *testing.T) {
	cases := []struct {
		a, b Segment
		expected []Point
	}{
		{seg(0, 0, 2, 0), seg(0, 1, 2, 1), nil},
		{seg(0, 0, 2, 0), seg(3, 0, 4, 0), nil},
		{seg(0, 0, 2, 0), seg(1, 0, 3, 0),
			[]Point{pt(1, 0), pt(2, 0)}},
			{seg(2, 0, 3, 0), seg(0, 0, 2, 0),
				[]Point{pt(2, 0)}},
		{seg(1, 0, 3, 0), seg(0, 0, 2, 0),
			[]Point{pt(1, 0), pt(2, 0)}},
	}
	for _, c := range cases {
		actual := hIntersects(c.a, c.b)
		if !EqualPoints(actual, c.expected) {
			t.Errorf("%v, %v intersect at %v, expected %v", c.a, c.b, actual, c.expected)
		}
	}
}

func TestVerticalIntersection(t *testing.T) {
	cases := []struct {
		a, b Segment
		expected []Point
	}{
		{seg(0, 0, 0,2), seg(1, 0, 1, 2), nil},
		{seg(0, 0, 0, 2), seg(0, 3, 0, 4), nil},
		{seg(0, 0, 0, 2), seg(0, 1, 0, 3),
			[]Point{pt(0, 1), pt(0, 2)}},
		{seg( 0, 2, 0, 3), seg(0, 0, 0, 2),
			[]Point{pt(0, 2)}},
		{seg(0, 1, 0, 3), seg(0, 0, 0, 2),
			[]Point{pt(0, 1), pt(0, 2)}},
	}
	for _, c := range cases {
		actual := vIntersects(c.a, c.b)
		if !EqualPoints(actual, c.expected) {
			t.Errorf("%v, %v intersect at %v, expected %v", c.a, c.b, actual, c.expected)
		}
	}
}


func pt(x, y int) Point {
	return Point{x, y}
}

func seg(xHead, yHead, xTail, yTail int) Segment {
	return Segment{
		Point{xHead, yHead},
		Point{xTail, yTail},
	}
}