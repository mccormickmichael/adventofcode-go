package intcode

import (
	"io"
	"io/ioutil"
)

type BuildOptions struct {
	program []int
	input  Inputter
	output Outputter
	halt chan bool
	dump io.Writer
}


func Builder(program []int) BuildOptions {
	prog := make([]int, len(program))
	copy(prog, program)
	return BuildOptions{
		program: prog,
		input:  &nilio{},
		output: &nilio{},
		halt:   make(chan bool, 1),
		dump:   ioutil.Discard,
	}
}

func (b BuildOptions) WithInputValue(value int) BuildOptions {
	return b.WithInputter(&ValueInput{value})
}

func (b BuildOptions) WithInputChannel(input chan int) BuildOptions {
	return b.WithInputter(&ChannelInput{input})
}

func (b BuildOptions) WithInputter(input Inputter) BuildOptions {
	b.input = input
	return b
}

func (b BuildOptions) WithOutputChannel(output chan int) BuildOptions {
	return b.WithOutputter(&ChannelOutput{output})
}

func (b BuildOptions) WithOutputter(output Outputter) BuildOptions {
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
