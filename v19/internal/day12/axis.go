package day12

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
	"io"
)

type monitor struct {
	pos   [4]int
	vel   [4]int
	count int
}

func (m *monitor) step() {
	p := m.pos
	m.vel[0] += intmath.Cmp(p[1], p[0]) + intmath.Cmp(p[2], p[0]) + intmath.Cmp(p[3], p[0])
	m.vel[1] += intmath.Cmp(p[0], p[1]) + intmath.Cmp(p[2], p[1]) + intmath.Cmp(p[3], p[1])
	m.vel[2] += intmath.Cmp(p[0], p[2]) + intmath.Cmp(p[1], p[2]) + intmath.Cmp(p[3], p[2])
	m.vel[3] += intmath.Cmp(p[0], p[3]) + intmath.Cmp(p[1], p[3]) + intmath.Cmp(p[2], p[3])

	m.pos[0] += m.vel[0]
	m.pos[1] += m.vel[1]
	m.pos[2] += m.vel[2]
	m.pos[3] += m.vel[3]

	m.count++
}

func findCycle(pos [4]int, vel [4]int, o io.Writer) {
	monitor := monitor{pos:pos, vel:vel}
	for i := 1; i <= 1000000; i++ {
		monitor.step()
		if monitor.pos == pos {
			if monitor.vel == vel {
				_, _ = fmt.Fprintf(o, "Matching position AND velocity in %d cycles\n", monitor.count)
				_, _ = fmt.Fprintf(o, "P: %v  V: %v \n", monitor.pos, monitor.vel)
			} else {
				_, _ = fmt.Fprintf(o, "Matching position in %d cycles\n", monitor.count)
				_, _ = fmt.Fprintf(o, "P: %v  V: %v \n", monitor.pos, monitor.vel)
			}
			monitor.count = 0
		}
	}
}