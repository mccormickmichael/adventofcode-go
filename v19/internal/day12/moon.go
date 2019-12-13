package day12

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
)

type vector struct {
	x, y, z int
}

type position vector

func (p position) energy() int {
	return intmath.Abs(p.x) + intmath.Abs(p.y) + intmath.Abs(p.z)
}

func (p position) String() string {
	return fmt.Sprintf("[%5d,%5d,%5d]", p.x, p.y, p.z)
}

type velocity vector

func (v velocity) energy() int {
	return intmath.Abs(v.x) + intmath.Abs(v.y) + intmath.Abs(v.z)
}

func (v velocity) String() string {
	return fmt.Sprintf("[%5d,%5d,%5d]", v.x, v.y, v.z)
}

type moon struct {
	name string
	pos  position
	vel  velocity
}

func (m *moon) String() string {
	return fmt.Sprintf("%s P: %s,   V: %s", m.name, m.pos, m.vel)
}

func (m *moon) energy() int {
	return m.pos.energy() * m.vel.energy()
}

func (m *moon) accelerate(moons []*moon) {
	for _, o := range moons {
		
		// parameters passed are opposite the obvious order so as to induce attraction
		m.vel.x += intmath.Cmp(o.pos.x, m.pos.x)
		m.vel.y += intmath.Cmp(o.pos.y, m.pos.y)
		m.vel.z += intmath.Cmp(o.pos.z, m.pos.z)
	}
}

func (m *moon) velocitate() {
	m.pos.x += m.vel.x
	m.pos.y += m.vel.y
	m.pos.z += m.vel.z
}
