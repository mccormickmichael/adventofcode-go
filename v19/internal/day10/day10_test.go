package day10

import (
	"math"
	"testing"
)

func TestSlope(t *testing.T) {
	cases := []struct {
		a, b asteroid
		expected float64
	}{
		{asteroid{0,0}, asteroid{1, 1}, 1.0},
		{asteroid{8, 7}, asteroid{0,2}, 0.625},
		{asteroid{0, 0}, asteroid{18, 0}, 0.0},
	}
	for _, c := range cases {
		actual := c.a.slope(c.b)
		if math.Abs(actual - c.expected) >= epsilon {
			t.Errorf("%v.slope(%v) == %.3f, expected %.3f", c.a, c.b, actual, c.expected)
		}
	}
}

func TestEqualSlopes(t *testing.T) {
	cases := []struct {
		a, b, mid asteroid
		expected  bool
	}{
		{asteroid{0,0}, asteroid{5, 0}, asteroid{2, 0}, true},
		{asteroid{0,0}, asteroid{3, 9}, asteroid{1, 3}, true},
		{asteroid{0,0}, asteroid{3, 10}, asteroid{1, 3}, false},
		{asteroid{0,0}, asteroid{0, 8}, asteroid{0,4}, true},
		{asteroid{0,0}, asteroid{0, 50}, asteroid{1, 49}, false},
	}
	for _, c := range cases {
		actual := c.mid.equalSlopes(c.a, c.b)
		if actual != c.expected {
			t.Errorf("%v.equalSlopes(%v, %v) == %t, expected %t", c.mid, c.a, c.b, actual, c.expected)
		}
	}
}

func TestPolar(t *testing.T) {
	cases := []struct {
		origin, a asteroid
		expectedThetaRep string
	}{
		{asteroid{0, 10},asteroid{0,0}, "000.000"},
		{asteroid{0, 0}, asteroid{10, 0}, "090.000"},
		{asteroid{0, 0}, asteroid{0, 10}, "180.000"},
		{asteroid{10, 0}, asteroid{0, 0}, "270.000"},
		{asteroid{0, 5}, asteroid{5, 0}, "045.000"},
		{asteroid{10, 0}, asteroid{5, 5}, "225.000"},
	}
	for _, c := range cases {
		actual := toPolar(c.origin, c.a).thetaRep
		if actual != c.expectedThetaRep {
			t.Errorf("expected Theta for %v -> %v == %s, expected %s", c.origin, c.a, actual, c.expectedThetaRep)
		}
	}
}