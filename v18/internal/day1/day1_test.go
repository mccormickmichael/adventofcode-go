package day1

import (
    "bufio"
    "strings"
    "testing"
)

func TestSum(t *testing.T) {

    cases := []struct {
        values []int
        expected int
    }{
        {[]int{1, 2, 3}, 6},
        {[]int{+1, -2, +3, +1}, 3},
    }

    for _, c := range cases {
        actual := Sum(c.values)
        if actual != c.expected {
            t.Errorf("Sum(%d) == %d, expected %d", c.values, actual, c.expected)
        }
    }
}

func TestParse(t *testing.T) {
    expected := []int{1, -2, 3, 1}
    source := `+1
-2
+3
+1`
    scanner := bufio.NewScanner(strings.NewReader(source))

    actual := parse(scanner)
    if !Equal(actual, expected) {
        t.Errorf("input(%s) == %d, expected %d", source, actual, expected)
    }
}

func TestDup(t *testing.T) {
    cases := []struct {
        frequencies []int
        duplicate, index int
    }{
        {[]int{1, -2, 3, 1}, 2, 6},
        {[]int{1, -1}, 0, 2},
        {[]int{3, 3, 4, -2, -4}, 10, 7},
        {[]int{-6, 3, 8, 5, -6}, 5, 12},
        {[]int{7, 7, -2, -7, -4}, 14, 13},
    }

    for _, c := range cases {
        actualDuplicate, actualIndex := Dup(c.frequencies)
        if actualDuplicate != c.duplicate || actualIndex != c.index {
            t.Errorf("Dup(%d) == %d index %d, expected %d index %d",
                c.frequencies,
                actualDuplicate, actualIndex,
                c.duplicate, c.index)
        }
    }
}

func Equal(lhs, rhs []int) bool {
    if len(lhs) != len(rhs) {
        return false
    }
    for i, v := range lhs {
        if v != rhs[i] {
            return false
        }
    }
    return true
}