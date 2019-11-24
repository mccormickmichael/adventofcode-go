package day1

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func Part1(path string) {
    value := Sum(scan(path))
    fmt.Println(value)
}

func Part2(path string) {
    value, index := Dup(scan(path))
    fmt.Printf("Duplicate frequency [%d] at index %d\n", value, index)
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