package day15

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
	"strings"
)

var commands = map[string]int{
	"N": 1,
	"S": 2,
	"W": 3,
	"E": 4,
	"n": 1,
	"s": 2,
	"w": 3,
	"e": 4,
}

type droid struct {
	x, y int
}

func (d *droid) move(cmd string) {
	switch cmd {
	case "N":
		d.y--
	case "S":
		d.y++
	case "W":
		d.x--
	case "E":
		d.x++
	}
}

func (d day15) Manual() {
	program := input.ParseInts(input.SingleLineFile(d.Path))

	inc  := make(chan int, 2)
	outc := make(chan int, 2)

	ic := intcode.Builder(program).WithInputChannel(inc).WithOutputChannel(outc).Build()

	droid := &droid{}
	go ic.GoRun()
	for true {
		var cmd string
		_, _ = fmt.Scanln(&cmd)
		cmd = strings.TrimSpace(cmd)
		c, ok := commands[cmd]
		if !ok {
			_, _ = fmt.Fprintf(d.Output, "Bad command.\n")
			continue
		}
		inc <- c
		r := <- outc

		switch r {
		case 0:
			_, _ = fmt.Fprintf(d.Output, "Droid hit a wall Droid is at %v.\n", droid)
		case 1:
			droid.move(cmd)
			_, _ = fmt.Fprintf(d.Output, "Droid moved %s. Droid is at %v\n", cmd, droid)
		case 2:
			droid.move(cmd)
			_, _ = fmt.Fprintf(d.Output, "Droid found the oxygen sensor at %v!\n", droid)
		}
	}
}