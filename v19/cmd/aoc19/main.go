package main

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day14"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day3"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day4"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day5"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day6"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day7"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day8"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day9"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day10"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day11"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day12"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day13"
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
	path := fmt.Sprintf("day%d.input.txt", day)
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	d := event.NilDay()

	switch day {
	case 1: d =  day1.New(path, os.Stdout)
	case 2: d =  day2.New(path, os.Stdout)
	case 3: d =  day3.New(path, os.Stdout)
	case 4: d =  day4.New(152085, 670283, os.Stdout)
	case 5: d =  day5.New(path, os.Stdout)
	case 6: d =  day6.New(path, os.Stdout)
	case 7: d =  day7.New(path, os.Stdout)
	case 8: d =  day8.New(path, os.Stdout)
	case 9: d =  day9.New(path, os.Stdout)
	case 10: d = day10.New(path, os.Stdout)
	case 11: d = day11.New(path, os.Stdout)
	case 12: d = day12.New(path, os.Stdout)
	case 13: d = day13.New(path, os.Stdout)
	case 14: d = day14.New(path, os.Stdout)
	default:
		bail(fmt.Sprintf("Day %d is not implemented.\n", day))
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
