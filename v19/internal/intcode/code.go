package intcode

import (
	"fmt"
	"io"
)

type Intcode struct {
	mem    []int
	pc     int
	count  int
	halted bool
	input  chan int
	output chan int
	halt   chan bool
	Dump io.Writer
}

func New(values []int) *Intcode  {
	v := make([]int, len(values))
	copy(v, values)
	input := make(chan int, 3)
	output := make(chan int, 10)
	halt := make(chan bool)
	return &Intcode{mem: v, input: input, output: output, halt: halt}
}

func (ic *Intcode) SetInput(input ...int) {
	for _, i := range input {
		ic.input <- i
	}
}

func (ic *Intcode) PopInput() int {
	return <- ic.input
}

func (ic *Intcode) PushOutput(value int) {
	ic.output <- value
}

func (ic *Intcode) PopOutput() int {
	return <- ic.output
}

func (ic *Intcode) Len() int {
	return len(ic.mem)
}

func (ic *Intcode) Pc() int {
	return ic.pc
}

func (ic *Intcode) SetPc(newPc int) {
	ic.pc = newPc
}

func (ic *Intcode) Count() int {
	return ic.count
}

func (ic *Intcode) Peek(index int) int {
	return ic.mem[index]
}

func (ic *Intcode) Mpeek(index int, mode int) int {
	switch mode {
	case 1:  return index
	default: return ic.mem[index]
	}
}

func (ic *Intcode) Poke(index int, value int) {
	ic.mem[index] = value
}

func (ic *Intcode) Run() error {
	if DumpFlag { Dump(ic) }
	for !ic.halted {
		if err := ic.Step(); err != nil {
			ic.halted = true
			close(ic.output)
			return err
		}
	}
	close(ic.output)
	return nil
}

func (ic *Intcode) Step() error {
	op, err := operation(ic.mem, ic.pc)
	if err != nil {
		return ic.StepError(err)
	}
	ic.pc++
	adv := op.ex(ic)
	ic.count++
	ic.pc += adv
	if DumpFlag { Dump(ic) }
	return nil
}

func (ic *Intcode) StepError(cause error) error {
	c:= make([]int, len(ic.mem))
	copy(c, ic.mem)
	return ErrStep{ic.pc, ic.count, c, cause}
}
type ErrStep struct {
	pc, count int
	mem []int
	cause error
}
func (e ErrStep) Error() string {
	return fmt.Sprintf("Error executing step %d: %s\nDump: PC:%d Heap:%v", e.count, e.cause, e.pc, e.mem)
}