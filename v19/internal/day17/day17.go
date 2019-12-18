package day17

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"io"
	"strings"
)

type day17 event.Solvable

func New(path string, output io.Writer) event.Day {
	return day17{Path:path, Output:output}
}

// 33 rows, 51 columns

func (d day17) Part1() {
	program := input.ParseInts(input.SingleLineFile(d.Path))

	robot    := &robot{}
	scaffold := newScaffold(robot)
	builder  := &builder{scaffold:scaffold}
	
	ic := intcode.Builder(program).WithOutputter(builder).Build()

	ic.Run()

	scaffold.square()

	cols, rows := scaffold.extent()
	_, _ = fmt.Fprintf(d.Output, "rows: %d cols:%d\n", rows, cols)
	for r := 0; r < rows; r++ {
		_, _ = fmt.Fprintf(d.Output, "%2d %s\n", r, string(scaffold.cells[r]))
	}
	intersections := scaffold.intersections()

	sum := 0
	for _, i := range intersections {
		param := i.x * i.y
		_, _ = fmt.Fprintf(d.Output, "%v, %d\n", i, param)
		sum += param
	}
	_, _ = fmt.Fprintf(d.Output, "sum: %d", sum)
}

// L,6,R,12,R,8,R,8,R,12,L,12,R,8,R,12,L,12,L,6,R,12,R,8,R,12,L,12,L,4,L,4,L,6,R,12,R,8,R,12,L,12,L,4,L,4,L,6,R,12,R,8,R,12,L,12,L,4,L,4,R,8,R,12,L,12,
//
// A: L,6,R,12,R,8,
//
// B: R,8,R,12,L,12
// 
// C: R,12,L,12,L,4,L,4
//
// A,B,B,A,C,A,C,A,C,B

func (d day17) Part2() {
	program := input.ParseInts(input.SingleLineFile(d.Path))

	p := []int{'A',',','B',',','B',',','A',',','C',',','A',',','C',',','A',',','C',',','B','\n'}
	a := []int{'L',',','6',',','R',',','1','2',',','R',',','8','\n'}
	b := []int{'R',',','8',',','R',',','1','2',',','L',',','1','2','\n'}
	c := []int{'R',',','1','2',',','L',',','1','2',',','L',',','4',',','L',',','4','\n'}
	o := []int{'n','\n'}

	instructions := append(p, a...)
	instructions = append(instructions, b...)
	instructions = append(instructions, c...)
	instructions = append(instructions, o...)

	feed  := &feed{d.Output, strings.Builder{}, -1}
	logic := &intcode.SliceInput{instructions}

	ic := intcode.Builder(program).WithOutputter(feed).WithInputter(logic).Build()
	ic.Poke(0, 2)

	ic.Run()
	if ic.Error() != nil {
		_, _ = fmt.Fprintf(d.Output, "Unexpected error %s\n", ic.Error())
	} else {
		_, _ = fmt.Fprintf(d.Output, "Dust collected: %d\n", feed.dust)
	}
	
}
