package day11

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
)

type day11 event.Solvable


func New(path string, output io.Writer) event.Day {
	return day11{Path:path, Output:output}
}

func (d day11) Part1() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	panels := paint(program, 0)
	_, _ = fmt.Fprintf(d.Output, "Painted %d panels\n", len(panels))
}

func (d day11) Part2() {
	program := input.ParseInts(input.SingleLineFile(d.Path))
	panels := paint(program, 1)

	var minx, maxx, miny, maxy int
	
	for p := range panels {
		if p.x < minx {
			minx = p.x
		}
		if p.y < miny {
			miny = p.y
		}
		if p.x > maxx {
			maxx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
	}

	for y := maxy; y >= miny; y-- {
		for x := minx; x <= maxx; x++ {
			pos := position{x, y}
			if panels[pos] == 1 {
				fmt.Fprint(d.Output, "#")
			} else {
				fmt.Fprint(d.Output, ".")
			}
		}
		fmt.Fprint(d.Output, "\n")
	}
}
