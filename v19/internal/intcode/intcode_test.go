package intcode

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		values[] int
		index int
		expected int
	}{
		{[]int{1,0,0,0,99}, 0, 2},
		{[]int{1,2,3,3,99}, 3, 6},
		{[]int{1,0,1,4,99,5,6,0,99},0, 11},
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

func TestRun(t *testing.T) {
	cases := []struct {
		values []int
		index int
		expected int
	}{
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