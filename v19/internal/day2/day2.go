package day2

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Day2 event.Solvable

type Op struct {
	code, i1, i2, out int
}

type ErrOpcode struct {
	opcode int
}

func New(path string, output io.Writer) event.Day {
    return Day2{Path: path, Output: output}
}

func (d Day2) Part1() {

	values := scan(d.Path)
	ic := intcode.New(values)
	ic.Poke(1, 12)
	ic.Poke(2, 2)

	err := ic.Run()
	if err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unxepected error: %s", err)
		return
	}
	_, _ = fmt.Fprintf(d.Output, "Value: %d in %d instructions", ic.Peek(0), ic.Count())
}	

func (d Day2) Part2() {
	originalValues := scan(d.Path)
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			ic := intcode.Builder(originalValues).Build()
			ic.Poke(1, noun)
			ic.Poke(2, verb)

			if err := ic.Run(); err != nil {
				log.Fatalf("Unexpected error at noun:%d, verb:%d", noun, verb)
			}
			if ic.Peek(0) == 19690720 {
				_, _ = fmt.Fprintf(d.Output, "noun: %d verb: %d answer %d\n", noun, verb, noun * 100 + verb)
				return
			}
		}
	}
	_, _ = fmt.Fprintf(d.Output, "Did not find any input pair that produced 19690720!\n")
}

func scan(path string) []int {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return parse(string(bytes))
}

func parse(s string) []int {
	tokens := strings.Split(s, ",")
	values := make([]int, len(tokens))
	for i, t := range tokens {
		v, _ := strconv.Atoi(t)
		values[i] = v
	}
	return values
}

