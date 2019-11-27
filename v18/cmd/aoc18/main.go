package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mccormickmichael/adventofcode-go/v18/internal/day1"
	"github.com/mccormickmichael/adventofcode-go/v18/internal/event"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Expected numeric value for first argument, got %s", os.Args[1])
	}
	path := fmt.Sprintf("input/day%d.txt", day)
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	var d event.Day

	switch day {
	case 1:
		d = day1.New(path, os.Stdout)
	default:
		log.Fatalf("Day %d is not implemented.", day)
	}
	fmt.Printf("Day %d of Advent of Code 2018:\n", day)
	d.Part1()
	d.Part2()
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage: aoc18 <day> [input file path]")
	os.Exit(1)
}