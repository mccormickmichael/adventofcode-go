package day2

import (
	"bufio"
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v18/internal/event"
	"io"
	"log"
	"os"
)

type Day2 event.Solvable

func New(path string, output io.Writer) event.Day {
	return Day2{Path: path, Output: output}
}

func (d Day2) Part1() {
	var twos, threes int
	boxIds := scan(d.Path)

	for _, b := range boxIds {
		if hasTwo(b) { twos++ }
		if hasThree(b) { threes++ }
	}
	checksum := checksum(twos, threes)

	_, _ = fmt.Fprintf(d.Output, "Checksum [%d] for %d twos and %d threes of %d ids\n", checksum, twos, threes, len(boxIds))
}

func (d Day2) Part2() {
	boxIds := scan(d.Path)
	for i, lhs := range boxIds {
		for _, rhs := range boxIds[i+1:] {
			if count, index := diff(lhs, rhs); count == 1 {
				_, _ = fmt.Fprintf(d.Output, "Common letters [%s] in box ids %s and %s\n", clean(lhs, index), lhs, rhs)
				return
			}
		}
	}
	_, _ = fmt.Fprintf(d.Output, "No sufficiently common box ids found!\n")
}

func scan(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	boxIds := make([]string, 100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boxIds = append(boxIds, scanner.Text())
	}
	return boxIds
}

func hasN(boxid string, n int) bool {
	histo := map[byte]int{}

	for i := 0; i < len(boxid); i++ {
		histo[boxid[i]] = histo[boxid[i]] + 1
	}
	for _, count := range histo {
		if count == n {
			return true
		}
	}
	return false
}

func checksum(twos, threes int) int {
	return twos * threes
}

func hasTwo(boxid string) bool {
	return hasN(boxid, 2)
}

func hasThree(boxid string) bool {
	return hasN(boxid, 3)
}

func clean(boxId string, index int) string {
	if index < 0 || index >= len(boxId) { return boxId }
	return boxId[:index] + boxId[index+1:]
}

func diff(lhs string, rhs string) (count, index int) {
	count = 0
	index = -1
	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			count++
			if index < 0 {
				index = i
			}
		}
	}
	return
}