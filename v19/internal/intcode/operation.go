package intcode

import (
	"fmt"
)

type Operation interface {
	ex(*Intcode) (int, error)
}

func operation(memory []int, ip int) Operation {
	var op Operation
	switch opcode := memory[ip]; opcode {
	case 99:
		op = Halt{}
	case 1:
		op = Add{memory[ip+1], memory[ip+2], memory[ip+3]}
	case 2:
		op = Mul{memory[ip+1], memory[ip+2], memory[ip+3]}
	case 3:
		op = Input{memory[ip+1]}
	case 4:
		op = Output{memory[ip+1]}
	default:
		op = ErrOpcode{opcode}
	}
	return op
}

type ErrOpcode struct {
	opcode int
}
func (e ErrOpcode) ex(ic *Intcode) (int, error) {
	return 0, e
}
func (e ErrOpcode) Error() string {
	return fmt.Sprintf("Unexpected Opcode %d", e.opcode)
}

type Halt struct{}
func (h Halt) ex(ic *Intcode) (int, error) {
	ic.halt = true
	return 0, nil
}

type Binop struct {
	id1, id2 int
	result   int
}

type Add Binop
func (a Add) ex(ic *Intcode) (int, error) {
	if ok := ensureLength(ic.mem, ic.pc+3); ok {
		ic.mem[a.result] = ic.mem[a.id1] + ic.mem[a.id2]
		return 3, nil
	}
	return 0, ErrOutOfRange{ic, 3}
}

type Mul Binop
func (a Mul) ex(ic *Intcode) (int, error) {
	if ok := ensureLength(ic.mem, ic.pc+3); ok {
		ic.mem[a.result] = ic.mem[a.id1] * ic.mem[a.id2]
		return 3, nil
	}
	return 0, ErrOutOfRange{ic, 3}
}

type Ioop struct {
	index int
}

type Input Ioop
func (i Input) ex(ic *Intcode) (int, error) {
	if ok := ensureLength(ic.mem, ic.pc+1); ok {
		ic.mem[i.index] = ic.input
		return 1, nil
	}
	return 0, ErrOutOfRange{ic, 1}
}

type Output Ioop
func (o Output) ex(ic *Intcode) (int, error) {
	if ok := ensureLength(ic.mem, ic.pc+1); ok {
		ic.output = ic.mem[o.index]
		return 1, nil
	}
	return 0, ErrOutOfRange{ic, 1}
}

func ensureLength(mem []int, max int) bool {
	return len(mem) > max
}

type ErrOutOfRange struct {
	ic *Intcode
	need int
}
func (e ErrOutOfRange) Error() string {
	return fmt.Sprintf("Program Counter result of range: pc:%d, limit:%d, needed:%d", e.ic.Pc(), e.ic.Len(), e.need)
}