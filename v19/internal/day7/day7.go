package day7

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"io"
	"log"
)

type Day7 event.Solvable

func New(path string, output io.Writer) event.Day {
	return Day7{Path: path, Output: output}
}

func (d Day7) Part1() {
	var maxOutput int
	var maxOutputPhase [5]int

	phases := EnumeratePhases([]int{0, 1, 2, 3, 4})
	program := input.ParseInts(input.SingleLineFile(d.Path))
	for _, p := range phases {
		output := Amplify(p, program)
		if output > maxOutput {
			maxOutput = output
			maxOutputPhase = p
		}
	}
	_, _ = fmt.Fprintf(d.Output, "Max Output %d at Phase %v\n", maxOutput, maxOutputPhase)
}

func (d Day7) Part2() {
	var maxOutput int
	var maxOutputPhase [5]int

	phases := EnumeratePhases([]int{5, 6, 7, 8, 9})
	program := input.ParseInts(input.SingleLineFile(d.Path))
	for _, p := range phases {
		output := Amplify(p, program)
		if output > maxOutput {
			maxOutput = output
			maxOutputPhase = p
		}
	}
	_, _ = fmt.Fprintf(d.Output, "Max Output %d at Phase %v\n", maxOutput, maxOutputPhase)
}

// This is a terrible way to do this.
func EnumeratePhases(r0 []int) [][5]int {
	phases := make([][5]int, 0)
	for index0, a := range r0 {
		r1 := make([]int, len(r0))
		copy(r1, r0)
		r1 = append(r1[:index0], r1[index0+1:]...)
		for index1, b := range r1 {
			r2 := make([]int, len(r1))
			copy(r2, r1)
			r2 = append(r2[:index1], r2[index1+1:]...)
			for index2, c := range r2 {
				r3 := make([]int, len(r2))
				copy(r3, r2)
				r3 = append(r3[:index2], r3[index2+1:]...)
				for index3, d := range r3 {
					r4 := make([]int, len(r3))
					copy(r4, r3)
					r4 = append(r4[:index3], r4[index3+1:]...)
					phase := [5]int{a, b, c, d, r4[0]}
					phases = append(phases, phase)
				}
			}
		}
	}
	return phases
}

func Amplify(phase [5]int, program []int) int {
	ab := make(chan int, 2)
	bc := make(chan int, 2)
	cd := make(chan int, 2)
	de := make(chan int, 2)
	ea := make(chan int, 2)

	ehalt := make(chan bool)

	ics := []*intcode.Intcode{
		intcode.Builder(program).WithInput(ea).WithOutput(ab).Build(),
		intcode.Builder(program).WithInput(ab).WithOutput(bc).Build(),
		intcode.Builder(program).WithInput(bc).WithOutput(cd).Build(),
		intcode.Builder(program).WithInput(cd).WithOutput(de).Build(),
		intcode.Builder(program).WithInput(de).WithOutput(ea).WithHalt(ehalt).Build(),
	}
	for i, ic := range ics {
		go ic.GoRun()
		ic.SetInput(phase[i])
	}
	ics[0].SetInput(0)

	<- ehalt

	var output int
	for output = range ea {}

	for amp, ic := range ics {
		if err := ic.Error(); err != nil {
			log.Fatalf("Unexpected error at phase %v amp %d: %v", phase, amp, err)
		}
	}

	return output
}