package day2

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
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
	_, _ = fmt.Fprintf(d.Output, "Value count: %d", len(values))
	preset(values)
	count, err := run(values);
	if err != nil {
		_, _ = fmt.Fprintf(d.Output, "Unxepected error: %s", err)
		return
	}
	_, _ = fmt.Fprintf(d.Output, "Value: %d in %d instructions", values[0], count)
}	

func (d Day2) Part2() {
	_, _ = fmt.Fprintf(d.Output, "Unimplemented!\n")
}

func preset(values []int) {
	values[1] = 12
	values[2] = 2
}

func run(values []int) (int, error) {
	count := 0
	for i := 0; i < len(values); i += 4 {
		op := Op{values[i], values[i+1], values[i+2], values[i+3]}
		if op.code == 99 {
			log.Printf("values[%d]==99, halting", i)
			return count, nil
		}
		if err := op.execute(values); err != nil {
			log.Printf("values[%d] failed: %s", i, err)
			return count, err
		}
		count++
	}
	return count, nil
}	


func (op *Op) execute(values []int) error {
	switch op.code {
	case 1:
		values[op.out] = values[op.i1] + values[op.i2]
	case 2:
		values[op.out] = values[op.i1] * values[op.i2]
	default:
		return &ErrOpcode{op.code}
	}
	return nil
}

func (e *ErrOpcode) Error() string {
	return fmt.Sprintf("Unexpected Opcode %d", e.opcode)
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

