package day17

import (
	"fmt"
	"io"
	"strings"
)

type feed struct {
	o io.Writer
	b strings.Builder
	dust int
}

func (f *feed) Output(o int) {

	if o > 128 {
		f.dust = o
		return
	}
	
	b := byte(o)

	if b == '\n' {
		_, _ = fmt.Fprintln(f.o, f.b.String())
		f.b = strings.Builder{}
		return
	}
	f.b.WriteByte(b)
}

func (f *feed) Close() {
}

type logic struct {
	instructions []byte
}

func (l *logic) Input() int {
	v := int(l.instructions[0])
	l.instructions = l.instructions[1:]
	return v
}
