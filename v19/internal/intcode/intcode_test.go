package intcode

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/test"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		program[] int
		id1, id2, result int
		modes int
		expected int
	}{
		{[]int{2}, 0, 0, 0, 0, 4},
		{[]int{0}, 3, 6, 0, 1100, 9},
		{[]int{3}, 0, 2, 0, 1000, 5},
		{[]int{3}, 2, 0, 0, 100, 5},
	}
	for _, c := range cases {
		ic := New(c.program)
		a := Add{c.id1, c.id2, c.result, ParseModes(c.modes)}
		_, _ = a.ex(ic)
		actual := ic.Peek(0)
		if actual != c.expected {
			t.Errorf("Add(%d, %d, %d, %v) == %d, expected %d %v", c.id1, c.id2, c.result, c.modes, actual, c.expected, ic.mem)
		}
	}
}

func TestInputOutput(t *testing.T) {
	cases := []struct {
		program []int
		input int
		expectedOutput int
	}{
		{[]int{3,0,4,0,99}, 5, 5},
		{[]int{3,0,4,0,99}, 7, 7},
		{[]int{3,0,104,110,99}, 0, 110},
	}
	for _, c := range cases {
		ic := New(c.program)
		ic.SetInput(c.input)
		err := ic.Run()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		actual := ic.Output()
		if actual != c.expectedOutput {
			t.Errorf("Run(%v) yielded %d, expected %d [%v]", c.program, actual, c.expectedOutput, ic.mem)
		}
	}
}

func TestRun(t *testing.T) {
	cases := []struct {
		values []int
		index int
		expected int
	}{
		{[]int{1,0,0,0,99}, 0, 2},
		{[]int{1,2,3,3,99}, 3, 6},
		{[]int{1,0,1,4,99,5,6,0,99},0, 11},
		{[]int{1,1,1,4,99,5,6,0,99}, 0, 30},
	}
	for _, c := range cases {
		ic := New(c.values)
		err := ic.Run()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		actual := ic.Peek(c.index)
		if actual != c.expected {
			t.Errorf("Run(%v) yielded %d, expected %d [%v]", c.values, actual, c.expected, ic.mem)
		}
	}
}

func TestParseModes(t *testing.T) {
	cases := []struct {
		instruction int
		expectedModes []int
	}{
		{10102, []int{1, 0, 1}},
		{10099, []int{0, 0, 1}},
		{101, []int{1}},
	}
	for _, c := range cases {
		actualModes := ParseModes(c.instruction)
		if !test.EqualIntSlice(actualModes, c.expectedModes) {
			t.Errorf("ParseModes(%d) == %v, expected %v", c.instruction, actualModes, c.expectedModes)
		}
	}
}