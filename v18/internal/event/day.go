package event

import (
	"io"
)

type Day interface {
	Part1()
	Part2()
}

type DayThing struct {
	Path string
	Output io.Writer
}
