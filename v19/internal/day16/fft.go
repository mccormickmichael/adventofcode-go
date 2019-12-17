package day16

import (
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
)

func fft(msg []int) {
	for i := 1; i < len(msg); i++ {
		applyOne(msg, i)
	}
}


func applyOne(msg []int, index int) {
	pos := 0
	neg := 0
	for p := index; p < len(msg); p += index * 4 {
		end := intmath.Min(p + index, len(msg))
		for i := p; i < end; i++ {
			pos += msg[i]
		}
	}
	
	for p := index * 3; p < len(msg); p += index * 4 {
		end := intmath.Min(p + index, len(msg))
		for i := p; i < end; i++ {
			neg += msg[i]
		}
	}

	msg[index] = intmath.Abs(pos - neg) % 10
}
