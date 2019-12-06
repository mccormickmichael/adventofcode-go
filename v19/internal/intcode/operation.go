package intcode

import (
	"fmt"
)

type Operation interface {
	ex(*Intcode) (int, error)
}

func operation(memory []int, ip int) Operation {
	var op Operation
	value := memory[ip]
	opcode := value % 100
	modes := ParseModes(value)
	switch opcode {
	case 99:
		op = Halt{}
	case 1:
		op = Add{memory[ip+1], memory[ip+2], memory[ip+3], modes}
	case 2:
		op = Mul{memory[ip+1], memory[ip+2], memory[ip+3], modes}
	case 3:
		op = Input{memory[ip+1], modes}
	case 4:
		op = Output{memory[ip+1], modes}
	default:
		op = ErrOpcode{opcode}
	}
	return op
}

type Modes []int
func ParseModes(instruction int) Modes {
	modes := make(Modes, 0)
	v := instruction/ 100
	for v != 0 {
		modes = append(modes, v % 10)
		v = v / 10
	}
	return modes
}
func (m Modes) Mode(i int) int {
	if i < len(m) {
		return m[i]
	}
	return 0
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
	id0, id1 int
	result   int
	modes    Modes
}

type Add Binop
func (a Add) ex(ic *Intcode) (int, error) {
	val0 := ic.Mpeek(a.id0, a.modes.Mode(0))
	val1 := ic.Mpeek(a.id1, a.modes.Mode(1))
	ic.Poke(a.result, val0 + val1)
	return 3, nil
}

type Mul Binop
func (a Mul) ex(ic *Intcode) (int, error) {
	val0 := ic.Mpeek(a.id0, a.modes.Mode(0))
	val1 := ic.Mpeek(a.id1, a.modes.Mode(1))
	ic.Poke(a.result, val0 * val1)
	return 3, nil
}

type Ioop struct {
	index int
	modes  Modes
}

type Input Ioop
func (i Input) ex(ic *Intcode) (int, error) {
		ic.Poke(i.index, ic.input)
		return 1, nil
}

type Output Ioop
func (o Output) ex(ic *Intcode) (int, error) {
		ic.output = ic.Mpeek(o.index, o.modes.Mode(0))
		return 1, nil
}

type ErrOutOfRange struct {
	ic *Intcode
	need int
}
func (e ErrOutOfRange) Error() string {
	return fmt.Sprintf("Program Counter result of range: pc:%d, limit:%d, needed:%d", e.ic.Pc(), e.ic.Len(), e.need)
}