package main

import (
    "log"
    "os"
    "strconv"

    "github.com/mccormickmichael/adventofcode-go/v18/internal/day1"
)

func main() {
    part, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Fatalf("Expected numeric value for first argumnt, got %s", os.Args[1])
    }
    path := "input/day1.txt"
    if len(os.Args) > 2 {
        path = os.Args[2]
    }

    switch part {
    case 1:
        day1.Part1(path)
    case 2:
        day1.Part2(path)
    default:
        log.Fatalf("Argument 'Part' must be 1 or 2, got %d", part)
    }
}