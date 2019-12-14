package day11

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intcode"
)

type position struct {
	x, y int
}

func (p position) move(d direction) position {
	return position{p.x + d.dx, p.y + d.dy}
}

type robot struct {
	pos position
	dir *direction
}

func (r *robot) move() {
	r.pos = r.pos.move(*(r.dir))
}

func (r *robot) turn(d int) {
	r.dir = r.dir.next[d]
}

type direction struct {
	dx, dy int
	next []*direction
}

var Up    = &direction{dx: 0, dy: 1}
var Down  = &direction{dx: 0, dy:-1}
var Left  = &direction{dx:-1, dy: 0}
var Right = &direction{dx: 1, dy: 0}

func init() {
	Up.next    = []*direction{Left,  Right}
	Down.next  = []*direction{Right, Left}
	Left.next  = []*direction{Down,  Up}
	Right.next = []*direction{Up,    Down}
}

type panels map[position]int

func paint(program []int, startColor int) panels {

	input := make(chan int, 1)
	output := make(chan int, 2)
	ic := intcode.Builder(program).WithInputChannel(input).WithOutputChannel(output).Build()

	go ic.GoRun()

	robot := &robot{dir: Up}
	input <- startColor
	panels := make(panels)

	var dir int
	var color int
	var running = true
	for running {
		color, running = <- output
		if running {
			dir, running = <- output
			if running {

				panels[robot.pos] = color
				robot.turn(dir)
				robot.move()
				input <- panels[robot.pos]
			}
		}
	}
	
	return panels
}
