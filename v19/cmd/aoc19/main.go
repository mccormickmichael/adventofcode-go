package main

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day3"
	"log"
	"os"
	"strconv"

	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day1"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day2"
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

	d := event.NilDay()

	switch day {
	case 1: d = day1.New(path, os.Stdout)
	case 2: d = day2.New(path, os.Stdout)
	case 3: d = day3.New(path, os.Stdout)
	default:
		bail(fmt.Sprintf("Day %d is not implemented.", day))
	}
	fmt.Printf("Day %2d of Advent of Code 2018:\n", day)
	fmt.Println("------------------------------")

	fmt.Println("Part 1:")
	d.Part1()

	fmt.Println()

	fmt.Println("Part 2:")
	d.Part2()
}

func bail(message string) {
	_, _ = fmt.Fprintf(os.Stderr, message)
	os.Exit(1)
}

func usage() {
	bail("Usage: aoc18 <day> [input file path]")
}
