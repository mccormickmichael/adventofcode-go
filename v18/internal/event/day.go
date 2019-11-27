package event

import (
	"io"
)

type Day interface {
	Part1()
	Part2()
}

type nilDay struct {
	string
}
func (n nilDay) Part1() {}
func (n nilDay) Part2() {}

func NilDay() Day{
	return nilDay{"bogus"}
}

type Solvable struct {
	Path string
	Output io.Writer
}
