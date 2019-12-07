package intcode

import (
	"fmt"
)

type Operation interface {
	ex(*Intcode) int
}

func operation(memory []int, ip int) (Operation, error) {
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
	case 5:
		op = JumpTrueOp{memory[ip+1], memory[ip+2], modes}
	case 6:
		op = JumpFalseOp{memory[ip+1], memory[ip+2], modes}
	case 7:
		op = LtCmpOp{memory[ip+1], memory[ip+2], memory[ip+3], modes}
	case 8:
		op = EqCmpOp{memory[ip+1], memory[ip+2], memory[ip+3], modes}
	default:
		return nil, ErrOpcode{opcode}
	}
	return op, nil
}

type Modes []int

func ParseModes(instruction int) Modes {
	modes := make(Modes, 0)
	v := instruction / 100
	for v != 0 {
		modes = append(modes, v%10)
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

func (e ErrOpcode) Error() string {
	return fmt.Sprintf("Unexpected Opcode %d", e.opcode)
}

type Halt struct{}

func (h Halt) ex(ic *Intcode) int {
	ic.Halt()
	return 0
}

type Binop struct {
	id0, id1 int
	result   int
	modes    Modes
}

type Add Binop

func (a Add) ex(ic *Intcode) int {
	val0 := ic.Mpeek(a.id0, a.modes.Mode(0))
	val1 := ic.Mpeek(a.id1, a.modes.Mode(1))
	ic.Poke(a.result, val0+val1)
	return 3
}

type Mul Binop

func (a Mul) ex(ic *Intcode) int {
	val0 := ic.Mpeek(a.id0, a.modes.Mode(0))
	val1 := ic.Mpeek(a.id1, a.modes.Mode(1))
	ic.Poke(a.result, val0*val1)
	return 3
}

type Ioop struct {
	index int
	modes Modes
}

type Input Ioop

func (i Input) ex(ic *Intcode) int {
	ic.Poke(i.index, ic.PopInput())
	return 1
}

type Output Ioop

func (o Output) ex(ic *Intcode) int {
	ic.PushOutput(ic.Mpeek(o.index, o.modes.Mode(0)))
	return 1
}

type JumpOp struct {
	value int
	dest  int
	modes Modes
}

type JumpTrueOp JumpOp

func (j JumpTrueOp) ex(ic *Intcode) int {
	val := ic.Mpeek(j.value, j.modes.Mode(0))
	if val != 0 {
		ic.SetPc(ic.Mpeek(j.dest, j.modes.Mode(1)))
		return 0
	}
	return 2
}

type JumpFalseOp JumpOp

func (j JumpFalseOp) ex(ic *Intcode) int {
	val := ic.Mpeek(j.value, j.modes.Mode(0))
	if val == 0 {
		ic.SetPc(ic.Mpeek(j.dest, j.modes.Mode(1)))
		return 0
	}
	return 2
}

type CmpOp struct {
	p0    int
	p1    int
	dest  int
	modes Modes
}

type LtCmpOp CmpOp

func (c LtCmpOp) ex(ic *Intcode) int {
	val0 := ic.Mpeek(c.p0, c.modes.Mode(0))
	val1 := ic.Mpeek(c.p1, c.modes.Mode(1))
	if val0 < val1 {
		ic.Poke(c.dest, 1)
	} else {
		ic.Poke(c.dest, 0)
	}
	return 3
}

type EqCmpOp CmpOp

func (c EqCmpOp) ex(ic *Intcode) int {
	val0 := ic.Mpeek(c.p0, c.modes.Mode(0))
	val1 := ic.Mpeek(c.p1, c.modes.Mode(1))
	if val0 == val1 {
		ic.Poke(c.dest, 1)
	} else {
		ic.Poke(c.dest, 0)
	}
	return 3
}

type ErrOutOfRange struct {
	ic   *Intcode
	need int
}

func (e ErrOutOfRange) Error() string {
	return fmt.Sprintf("Program Counter result of range: pc:%d, limit:%d, needed:%d", e.ic.Pc(), e.ic.Len(), e.need)
}
