package intcode

import (
	"io"
	"io/ioutil"
)

type BuildOptions struct {
	program []int
	input, output chan int
	halt chan bool
	dump io.Writer
}

func Builder(program []int) BuildOptions {
	prog := make([]int, len(program))
	copy(prog, program)
	return BuildOptions{
		program: prog,
		input:  make(chan int, 2),
		output: make(chan int, 10),
		halt:   make(chan bool, 1),
		dump:   ioutil.Discard,
	}
}

func (b BuildOptions) WithInput(input chan int) BuildOptions {
	b.input = input
	return b
}

func (b BuildOptions) WithOutput(output chan int) BuildOptions {
	b.output = output
	return b
}
func (b BuildOptions) WithHalt(halt chan bool) BuildOptions {
	b.halt = halt
	return b
}

func (b BuildOptions) WithDump(dump io.Writer) BuildOptions {
	b.dump = dump
	return b
}

func (b BuildOptions) Build() *Intcode {
	return &Intcode{mem:b.program, input:b.input, output:b.output, halt:b.halt, dump:b.dump}
}