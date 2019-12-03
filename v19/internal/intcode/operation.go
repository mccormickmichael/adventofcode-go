package intcode

import (
	"fmt"
)

type Operation interface {
	ex(*Intcode) (int, error)
}

func operation(memory []int, instructionPointer int) Operation {
	var op Operation
	switch opcode := memory[instructionPointer]; opcode {
	case 99:
		op = Halt{}
	case 1:
		op = Add{}
	case 2:
		op = Mul{}
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

type Add struct{}
func (a Add) ex(ic *Intcode) (int, error) {
	if ok := ensureLength(ic.mem, ic.pc+2); ok {
		ic.mem[ic.pc+2] = ic.mem[ic.pc] + ic.mem[ic.pc+1]
		return 3, nil
	}
	return 0, ErrOutOfRange{len(ic.mem), ic.pc, 3}
}

type Mul struct{}
func (a Mul) ex(ic *Intcode) (int, error) {
	if ok := ensureLength(ic.mem, ic.pc+2); ok {
		ic.mem[ic.pc+2] = ic.mem[ic.pc] * ic.mem[ic.pc+1]
		return 3, nil
	}
	return 0, ErrOutOfRange{len(ic.mem), ic.pc, 3}
}

func ensureLength(mem []int, max int) bool {
	return len(mem) > max
}

type ErrOutOfRange struct {
	max, pc, need int
}
func (e ErrOutOfRange) Error() string {
	return fmt.Sprintf("Program Counter out of range: pc:%d, limit:%d, needed:%d", e.pc, e.max, e.need)
}