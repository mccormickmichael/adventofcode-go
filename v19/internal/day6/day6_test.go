package day6

import (
	"strings"
	"testing"
)

func TestBuildTree(t *testing.T) {

	entries := []Entry{{"A", "B"}, {"B", "C"}, {"B", "D"}}
	tree := make(OrbitTree)

	tree.buildOrbitTree(entries)

	if tree["A"].parent != nil {
		t.Errorf("A.parent was %s, expected nil", tree["A"].parent)
	}
	if tree["B"].parent.id != "A" {
		t.Errorf("B.parent was %s, expected A", tree["B"].parent)
	}
	if tree["C"].parent.id != "B" {
		t.Errorf("C.parent was %s, expected B", tree["C"].parent)
	}
	if tree["D"].parent.id != "B" {
		t.Errorf("D.parent was %s, expected B", tree["D"].parent)
	}
}

func TestPath(t *testing.T) {
	entries := []Entry{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"},
		{"E", "F"}, {"B", "G"}, {"G", "H"}, {"D", "I"},
		{"E", "J"}, {"J", "K"}, {"K", "L"}}

	tree := make(OrbitTree)
	tree.buildOrbitTree(entries)

	cases := []struct {
		id   string
		path string
		len  int
	}{
		{"A", "A", 0},
		{"B", "A.B", 1},
		{"C", "A.B.C", 2},
		{"D", "A.B.C.D", 3},
		{"E", "A.B.C.D.E", 4},
		{"F", "A.B.C.D.E.F", 5},
		{"G", "A.B.G", 2},
		{"H", "A.B.G.H", 3},
		{"I", "A.B.C.D.I", 4},
		{"J", "A.B.C.D.E.J", 5},
		{"K", "A.B.C.D.E.J.K", 6},
		{"L", "A.B.C.D.E.J.K.L", 7},
	}

	for _, c := range cases {
		actualPath := strings.Join(tree[c.id].Path(), ".")
		if actualPath != c.path {
			t.Errorf("Path(%s) was %s, expected %s", c.id, actualPath, c.path)
		}
		actualLen := tree[c.id].PathLen()
		if actualLen != c.len {
			t.Errorf("PathLen(%s) was %d, expected %d", c.id, actualLen, c.len)
		}
	}
}

func TestUnCommonPath(t *testing.T) {
	entries := []Entry{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"},
		{"E", "F"}, {"B", "G"}, {"G", "H"}, {"D", "I"},
		{"E", "J"}, {"J", "K"}, {"K", "L"},
		{"K", "YOU"}, {"I", "SAN"}}
	tree := make(OrbitTree)
	tree.buildOrbitTree(entries)

	expectedA := "D.E.J.K.YOU"
	expectedB := "D.I.SAN"

	actual := UnCommonPaths(tree["YOU"], tree["SAN"])
	actualA := strings.Join(actual[0], ".")
	actualB := strings.Join(actual[1], ".")
	if actualA != expectedA {
		t.Errorf("uncommon YOU path %s, expected %s", actualA, expectedA)
	}
	if actualB != expectedB {
		t.Errorf("uncommon SAN path %s, expected %s", actualB, expectedB)
	}
}

func TestTransferCount(t *testing.T) {
	cases := []struct {
		a, b     []string
		expected int
	}{
		{[]string{"D", "E", "J", "K", "Y"}, []string{"D", "I", "S"}, 4},
		{[]string{"K", "Y"}, []string{"K", "S"}, 0},
		{[]string{"K", "Y"}, []string{"K", "L", "S"}, 1},
	}
	for _, c := range cases {
		tc := TransferCount(c.a, c.b)
		if tc != c.expected {
			t.Errorf("TransferCount(%v, %v) == %d, expected %d", c.a, c.b, tc, c.expected)
		}
	}
}