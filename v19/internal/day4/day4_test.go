package day4

import (
	"strconv"
	"testing"
)

func TestHasRepeats(t *testing.T) {
	cases := []struct {
		n int
		expected bool
	}{
		{111111, true},
		{123899, true},
		{987654, false},
	}
	for _, c := range cases {
		sv := strconv.Itoa(c.n)
		actual := hasRepeats([]byte(sv))
		if actual != c.expected {
			t.Errorf("hasRepeats(%s) == %t, expected %t", sv, actual, c.expected)
		}
	}
}

func TestHasDouble(t *testing.T) {
	cases := []struct {
		n int
		expected bool
	}{
		{111111, false},
		{112233, true},
		{123444, false},
		{111122, true},
	}
	for _, c := range cases {
		sv := strconv.Itoa(c.n)
		actual := hasDouble([]byte(sv))
		if actual != c.expected {
			t.Errorf("hasDouble(%s) == %t, expected %t", sv, actual, c.expected)
		}
	}
}

func TestIncreases(t *testing.T) {
	cases := []struct {
		n int
		expected bool
	}{
		{111111, true},
		{123899, true},
		{987654, false},
		{123560, false},
		{000001, true},
	}
	for _, c := range cases {
		sv := strconv.Itoa(c.n)
		actual := increases([]byte(sv))
		if actual != c.expected {
			t.Errorf("increases(%s) == %t, expected %t", sv, actual, c.expected)
		}
	}
}