package day8

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/event"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/input"
	"io"
	"strings"
)

type Day8 event.Solvable

func New(path string, output io.Writer) event.Day {
	return Day8{Path:path, Output:output}
}

const (
	ImageWidth  = 25
	ImageHeight = 6
)

func (d Day8) Part1() {
	s := strings.TrimSpace(input.SingleLineFile(d.Path))
	digits := toDigits(s)
	layers := makeLayers(digits, ImageHeight*ImageWidth)
	layerIndex := layerIndexWithLeastZeroes(layers)
	ones := count(layers[layerIndex], 1)
	twos := count(layers[layerIndex], 2)
	_, _ = fmt.Fprintf(d.Output, "Checksum for layer %d: %d\n", layerIndex, ones*twos)
}

func (d Day8) Part2() {
	s := strings.TrimSpace(input.SingleLineFile(d.Path))
	digits := toDigits(s)
	layers := makeLayers(digits, ImageHeight*ImageWidth)

	pixels := make([]byte, ImageHeight*ImageWidth)
	for i := 0; i < ImageWidth*ImageHeight; i++ {
		for li := 0; li < len(layers); li++ {
			if layers[li][i] == 0 {
				pixels[i] = ' '
				break
			} else if layers[li][i] == 1 {
				pixels[i] = '0'
				break
			}
		}
	}
	image := makeLayers(pixels, ImageWidth)
	for _, row := range image {
		_, _ = fmt.Fprintf(d.Output, string(row) + "\n")
	}
}

func makeLayers(digits []byte, layerSize int) [][]byte {
	layers := make([][]byte, len(digits)/layerSize)
	for i := 0; i < len(layers); i++ {
		layers[i] = digits[:layerSize]
		digits = digits[layerSize:]
	}
	return layers
}

func toDigits(s string) []byte {
	digits := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = s[i]-48 // this makes me slightly uncomfortable.
	}
	return digits
}

func layerIndexWithLeastZeroes(layers [][]byte) int {
	zeroCount := len(layers[0])
	layerIndex := 0
	for i, layer := range layers {
		zeroes := count(layer, 0)
		if zeroes < zeroCount {
			zeroCount = zeroes
			layerIndex = i
		}
	}
	return layerIndex
}

func count(layer []byte, value byte) int {
	count := 0
	for _, d := range layer {
		if d == value {
			count ++
		}
	}
	return count
}