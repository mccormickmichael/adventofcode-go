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

func TestOpcode(t *testing.T) {
	cases := []struct {
		values []int
		op Op
		expected int
	}{
		{[]int{1}, Op{1, 0, 0, 0}, 2},
	}

	for _, c := range cases {
		c.op.execute(c.values)
		actual := c.values[0]
		if actual != c.expected {
			t.Errorf("%c.execute(%c) yielded %d, expected %d", c.op, c.values, actual, c.expected)
		}
	}
}

func testRun(t *testing.T) {
	cases := []struct {
		values []int
		index, expected int
	}{
		{[]int{1,0,0,0,99}, 0, 2},
		{[]int{2,3,0,3,99}, 3, 6},
		{[]int{1,1,1,4,99,5,6,0,99}, 0, 30},
		{[]int{1,1,1,4,99,5,6,0,99}, 4, 2},
	}

	for _, c := range cases {
		_, e := run(c.values)
		if e != nil {
			t.Errorf("Unexpected error: %s", e)
		}
		actual := c.values[c.index]
		if actual != c.expected {
			t.Errorf("run(%c) == %d at %d, expected %d", c.values, actual, c.index, c.expected)
		}
	}
}
