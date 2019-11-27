package day1

import (
    "bufio"
    "fmt"
    "github.com/mccormickmichael/adventofcode-go/v18/internal/event"
    "io"
    "log"
    "os"
    "strconv"
)

type Day1 struct {
    event.DayThing
}


func New(path string, output io.Writer) *Day1 {
    return &Day1{event.DayThing{Path: path, Output: output}}
}
func (d Day1) Part1() {
    value := Sum(scan(d.Path))
    _, _ = fmt.Fprintf(d.Output, "Resulting frequency [%d]\n", value)
}

func (d Day1) Part2() {
    value, index := Dup(scan(d.Path))
    _, _ = fmt.Fprintf(d.Output, "Duplicate frequency [%d] at index %d\n", value, index)
}

func scan(path string) []int {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    return parse(scanner)
}

func parse(values *bufio.Scanner) []int {
    var s []int

    for values.Scan() {
        v, err := strconv.Atoi(values.Text())
        if err != nil {
            log.Fatal(err)
        }
        s = append(s, v)
    }
    return s
}

func Sum(values []int) int {
    value := 0
    for _, v := range values {
        value += v
    }
    return value
}

func Dup(frequencies []int) (frequency, index int) {
    frequency = 0
    index = 0
    knownFrequencies := map[int]bool{0: true}
    for index < 1000000 {
        f := frequencies[index % len(frequencies)]
        frequency = frequency + f
        index++
        if _, ok := knownFrequencies[frequency]; ok {
            return
        }
        knownFrequencies[frequency] = true
    }

    return -1, -1
}