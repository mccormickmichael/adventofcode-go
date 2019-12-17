package day16

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/test"
	"testing"
)

func TestPhaseGenerator(t *testing.T) {
	cases := []struct {
		period int
		expected []int
	}{
		{0, []int{1, 0, -1, 0}},
		{1, []int{0, 1, 1, 0, 0, -1, -1, 0}},
		{2, []int{0, 0, 1, 1, 1, 0, 0, 0, -1}},
		{4, []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, -1, -1, -1, -1, -1, 0, 0, 0, 0, 0}},
	}

	for _, c := range cases {
		ph := newPhaser(c.period)

		for i := 0; i < len(c.expected); i++ {
			actual := ph.next()
			if actual != c.expected[i] {
				t.Errorf("phaser(0)[%d] == %d, expected %d", i, actual, c.expected[i])
			}
		}
	}
}

func TestNextPhase(t *testing.T) {
	cases := []struct {
		input []int
		phases int
		expected[]int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 1, []int{4, 8, 2, 2, 6, 1, 5, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, []int{3, 4, 0, 4, 0, 4, 3, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, []int{0, 1, 0, 2, 9, 4, 9, 8}},
	}

	for _, c := range cases {
		actual := c.input
		for i := 0; i < c.phases; i++ {
			actual = nextPhase(actual)
		}
		if !test.EqualIntSlice(actual, c.expected) {
			t.Errorf("%v after %d phases was %v, expected %v", c.input, c.phases, actual, c.expected)
		}
	}
}

func TestNextPhaseBig(t *testing.T) {
	cases := []struct {
		input    string
		phases   int
		expected string
	}{
		{"80871224585914546619083218645595", 100, "24176176"},
		{"19617804207202209144916044189917", 100, "73745418"},
		{"69317163492948606335995924319873", 100, "52432133"},
	}

	for _, c := range cases {
		expected := input.Digits(c.expected)
		actual := input.Digits(c.input)
		for i := 0; i < c.phases; i++ {
			actual = nextPhase(actual)
		}
		if !test.EqualIntSlice(actual[:8], expected) {
			t.Errorf("%v after %d phases was %v, expected %v", c.input, c.phases, actual, c.expected)
		}
	}
}

func TestFFTSmall(t *testing.T) {
	cases := []struct {
		input []int
		phases int
		expected[]int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 1, []int{4, 8, 2, 2, 6, 1, 5, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, []int{3, 4, 0, 4, 0, 4, 3, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, []int{0, 1, 0, 2, 9, 4, 9, 8}},
	}

	for _, c := range cases {
		actual := make([]int, len(c.input)+1)
		copy(actual[1:], c.input)
		for i := 0; i < c.phases; i++ {
			fft(actual)
		}
		if !test.EqualIntSlice(actual[1:], c.expected) {
			t.Errorf("%v after %d phases was %v, expected %v", c.input, c.phases, actual[1:], c.expected)
		}
	}
}

func TestFftBig(t *testing.T) {
	cases := []struct {
		input    string
		phases   int
		expected string
	}{
		{"80871224585914546619083218645595", 100, "24176176"},
		{"19617804207202209144916044189917", 100, "73745418"},
		{"69317163492948606335995924319873", 100, "52432133"},
	}

	for _, c := range cases {
		actual := make([]int, len(c.input)+1)
		expected := input.Digits(c.expected)
		ad := input.Digits(c.input)
		copy(actual[1:], ad)
		for i := 0; i < c.phases; i++ {
			fft(actual)
		}
		if !test.EqualIntSlice(actual[1:9], expected) {
			t.Errorf("%v after %d phases was %v, expected %v", c.input, c.phases, actual, c.expected)
		}
	}
}
