package day1

import (
	"bufio"
	"fmt"
    "io"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
)

type Day1 event.Solvable

func New(path string, output io.Writer) event.Day {
    return Day1{Path: path, Output: output}
}

func (d Day1) Part1() {
	masses := scan(d.Path)
	var fuel int
	for _, m := range masses {
		fuel += calculateFuelForMass(m)
	}
	_, _ = fmt.Fprintf(d.Output, "Total fuel required: %d\n", fuel)
}

func (d Day1) Part2() {
	masses := scan(d.Path)
	var totalFuel int
	for _, m := range masses {
		totalFuel += calculateFuelForMassAndFuel(m)
	}
	_, _ = fmt.Fprintf(d.Output, "Total fuel required: %d\n", totalFuel)
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

func calculateFuelForMass(mass int) int {
	if mass <= 6 {
		return 0
	}
	massf := float64(mass)
	return int(math.Floor(massf/3))-2
}

func calculateFuelForMassAndFuel(mass int) int {
	if mass <= 6 {
		return 0
	}
	fuel := calculateFuelForMass(mass)
	return fuel + calculateFuelForMassAndFuel(fuel)
}

func calculateFuelForFuel(fuelMass int) int {
	if fuelMass < 6 {
		return 0
	}
	additionalFuel := calculateFuelForMass(fuelMass)
	return additionalFuel + calculateFuelForFuel(additionalFuel)
}