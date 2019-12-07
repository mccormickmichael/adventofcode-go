package intcode

import (
	"fmt"
	"io"
)

type Intcode struct {
	mem []int
	pc int
	count int
	halt bool
	input []int
	output int
	Dump io.Writer
}

func New(values []int) *Intcode  {
	v := make([]int, len(values))
	copy(v, values)
	return &Intcode{mem: v}
}

func (ic *Intcode) SetInput(input ...int) {
	ic.input = input
}

func (ic *Intcode) PopInput() int {
	val := ic.input[0]
	ic.input = ic.input[1:]
	return val
}

func (ic *Intcode) Output() int {
	return ic.output
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
	for !ic.halt {
		if err := ic.Step(); err != nil {
			ic.halt = true
			return err
		}
	}
	return nil
}

func (ic *Intcode) Step() error {
	op := operation(ic.mem, ic.pc)
	ic.pc++
	adv, err := op.ex(ic)
	if err != nil {
		return ic.StepError(err)
	}
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