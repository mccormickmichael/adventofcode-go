package day2

import "testing"

func TestHasTwo(t *testing.T) {
	cases := []struct {
		boxid    string
		expected bool
	}{
		{"abcdef", false},
		{"bababe", true},
		{"abbcde", true},
		{"abcccd", false},
		{"aabcdd", true},
		{"abcdee", true},
		{"ababab", false},
	}
	for _, c := range cases {
		actual := hasTwo(c.boxid)
		if actual != c.expected {
			t.Errorf("hasTwo(%s) == %t, expected %t", c.boxid, actual, c.expected)
		}
	}
}

func TestHasThree(t *testing.T) {
	cases := []struct {
		boxid    string
		expected bool
	}{
		{"abcdef", false},
		{"bababe", true},
		{"abbcde", false},
		{"abcccd", true},
		{"aabcdd", false},
		{"abcdee", false},
		{"ababab", true},
	}
	for _, c := range cases {
		actual := hasThree(c.boxid)
		if actual != c.expected {
			t.Errorf("hasTwo(%s) == %t, expected %t", c.boxid, actual, c.expected)
		}
	}
}

func TestDiff(t *testing.T) {
	cases := []struct {
		lhs, rhs string
		expectedCount, expectedIndex int
	}{
		{"axcye", "abcde", 2, 1},
		{"abcde", "abcde", 0, -1},
		{"fghij", "fguij", 1, 2},
	}
	for _, c := range cases {
		actualCount, actualIndex := diff(c.lhs, c.rhs)
		if actualCount != c.expectedCount {
			t.Errorf("diff(%s, %s) count == %d, expected %d", c.lhs, c.rhs, actualCount, c.expectedCount)
		}
		if actualIndex != c.expectedIndex {
			t.Errorf("diff(%s, %s) index == %d, expected %d", c.lhs, c.rhs, actualIndex, c.expectedIndex)
		}
	}
}

func TestClean(t *testing.T) {
	cases := []struct {
		id string
		index int
		expected string
	}{
		{"abcdef", 1, "acdef"},
		{"first", 0, "irst"},
		{"outofrange", 10, "outofrange"},
		{"last", 3, "las"},
	}
	for _, c := range cases {
		actual := clean(c.id, c.index)
		if actual != c.expected {
			t.Errorf("clean(%s, %d) == %s, expected %s", c.id, c.index, actual, c.expected)
		}
	}
}