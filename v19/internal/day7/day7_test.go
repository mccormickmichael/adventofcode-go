package day7

import "testing"

func TestAmplify(t *testing.T) {
	cases := []struct {
		phase [5]int
		program []int
		expected int
	}{
		{
			[5]int{4,3,2,1,0},
			[]int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0},
			43210,
		},
		{
			[5]int{0,1,2,3,4},
			[]int{3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0},
			54321,
		},
		{
			[5]int{1,0,4,3,2},
			[]int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0},
			65210,
		},
	}
	for _, c := range cases {
		output := Amplify(c.phase, c.program)

		if output != c.expected {
			t.Errorf("Amplify(%v, %v) == %d, expected %d", c.phase, c.program, output, c.expected)
		}
	}
}