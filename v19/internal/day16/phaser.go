package day16

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
)



var phase = [4]int{ 0, 1, 0, -1 }

type phaser interface {
	next() int
}

type phaseGenerator struct {
	period int
	index  int
}

func (p *phaseGenerator) next() int {
	p.index++
	return phase[(p.index/p.period) % 4]
}


// "period" is 0-based
func newPhaser(period int) phaser {
	return &phaseGenerator{period:period+1}
}

func nextPhase(src []int) []int {
	next := make([]int, len(src))

	for i := 0; i < len(src); i++ {
		p := newPhaser(i)
		next[i] = apply(src, p)
	}
	return next
}


func apply(src []int, p phaser) int {
	result := 0
	for i := 0; i < len(src); i++ {
		result += src[i] * p.next()
	}
	return intmath.Abs(result) % 10
}
