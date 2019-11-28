package day3

import (
	"testing"
)

func TestParseRecord(t *testing.T) {
	cases := []struct {
		line string
		expectedID, expectedX, expectedY, expectedWidth, expectedHeight uint16
	}{
		{"#1 @ 1,3: 4x4", 1, 1, 3, 4, 4},
		{"#2 @ 3,1: 4x4", 2, 3, 1, 4, 4},
		{"#3 @ 5,5: 2x2", 3, 5, 5, 2, 2},
	}

	for _, c := range cases {
		actual := parseClaim(c.line)
		if actual.id != c.expectedID {
			t.Errorf("parseClaim(%s).id == %d, expected %d", c.line, actual.id, c.expectedID)
		}
	}
}
