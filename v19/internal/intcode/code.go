package intcode

import (
	"fmt"
	"io"
	"log"
)

const ReallyBig = 1000000

type Intcode struct {
	mem    []int
	dump   io.Writer

	input  Inputter
	output Outputter
	halt   chan bool

	base   int
	pc     int
	count  int
	halted bool
	error  error
}

func New(values []int) *Intcode  {
	return Builder(values).Build()
}

func (ic *Intcode) SetInput(value int) {
	ic.input.Set(value)
}

func (ic *Intcode) PopInput() int {
	return ic.input.Input()
}

func (ic *Intcode) PushOutput(value int) {
	ic.output.Output(value)
}

func (ic *Intcode) Halt() {
	ic.halted = true
}

func (ic *Intcode) Error() error {
	return ic.error
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

func (ic *Intcode) MoveBase(offset int) {
	ic.base += offset
}

func (ic *Intcode) Peek(index int) int {
	return ic.mem[index]
}

func (ic *Intcode) Mpeek(index int, mode int) int {
	switch mode {
	case 0:
		if ic.hasCapacity(index) {
			return ic.mem[index]
		}
		return 0
	case 1:
		return index
	case 2:
		if ic.hasCapacity(index + ic.base) {
			return ic.mem[index + ic.base]
		}
		return 0
	default: ic.error = ModeError{"peek", ic, mode}
	}
	return 0
}

func (ic *Intcode) hasCapacity(need int) bool {
	return len(ic.mem) > need
}

func (ic *Intcode) resizeTo(cap int) {
	if cap > ReallyBig {
		log.Printf("WARN: capacity is pathologically large: %d", cap)
	}
	// TODO: deal with pathologically large capacities
	newMem := make([]int, cap)
	copy(newMem, ic.mem)
	ic.mem = newMem
}

func (ic *Intcode) Poke(index int, value int) {
	if !ic.hasCapacity(index) {
		ic.resizeTo(index+1)
	}
	ic.mem[index] = value
}

func (ic *Intcode) Mpoke(index int, mode int, value int) {
	switch mode {
	case 0: ic.Poke(index, value)
	case 2: ic.Poke(index + ic.base, value)
	case 1:
		ic.error = ModeError{"poke", ic,mode}
	}
}

func (ic *Intcode) GoRun() {
	err := ic.Run()
	ic.error = err
	ic.halt <- true
	ic.output.Close()
}

func (ic *Intcode) Run() error {
	if DumpFlag { Dump(ic) }
	for !ic.halted {
		ic.Step()
		if ic.error != nil {
			ic.halted = true
			return ic.error
		}
	}
	return nil
}

func (ic *Intcode) Step() {
	op, err := operation(ic.mem, ic.pc)
	if err != nil {
		ic.error = ic.StepError(err)
		return
	}
	ic.pc++
	adv := op.ex(ic)
	ic.count++
	ic.pc += adv
	if DumpFlag { Dump(ic) }
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

type ModeError struct {
	op   string
	ic   *Intcode
	mode int
}

func (e ModeError) Error() string {
	return fmt.Sprintf("Unknown mode %d for operation %s at %d on step %d", e.mode, e.op, e.ic.pc, e.ic.count)
}
