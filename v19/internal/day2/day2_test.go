package day2

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/test"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		input string
		expected []int
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}},
	}

	for _, c := range cases {
		actual := parse(c.input)
		if !test.EqualIntSlice(actual, c.expected) {
			t.Errorf("TestParse(%s) == %c, expected %c", c.input, actual, c.expected)
		}
	}
}