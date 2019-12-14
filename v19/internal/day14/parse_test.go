package day14

import "testing"

func TestParseComponent(t *testing.T) {
	cases := []struct {
		text string
		expected componentDescriptor
	}{
		{"3 A", componentDescriptor{"A", 3}},
		{"123 BDABGD", componentDescriptor{"BDABGD", 123}},
	}
	for _, c := range cases {
		actual := parseComponent(c.text)
		if actual != c.expected {
			t.Errorf("parseComponent(%s) == %v, expected %v", c.text, actual, c.expected)
		}
	}
}

func TestParseInputs(t *testing.T) {
	cases := []struct {
		text string
		expected []componentDescriptor
	}{
		{"144 ORE", []componentDescriptor{{"ORE", 144}}},
		{"144 ORE, 7 BAGS", []componentDescriptor{{"ORE", 144}, {"BAGS", 7}}},
		{"7 A, 3 B, 9 C", []componentDescriptor{{"A", 7}, {"B", 3}, {"C", 9}}},
	}
	for _, c := range cases {
		actual := parseInputs(c.text)
		if !equalComponentDesc(actual, c.expected) {
			t.Errorf("parseComponent(%s) == %v, expected %v", c.text, actual, c.expected)
		}
	}
}

func equalComponentDesc(a, b []componentDescriptor) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if ai != b[i] {
			return false
		}
	}
	return true
}