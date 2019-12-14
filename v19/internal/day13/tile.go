package day13

import (
	"fmt"
	"io"
	"strings"
)


const (
	Empty  = 0
	Wall   = 1
	Block  = 2
	Paddle = 3
	Ball   = 4
)

var glyphs = map[int]byte{
	Empty: ' ',
	Wall:  'X',
	Block: '=',
	Paddle:'-',
	Ball:  'o',
}
	
type coord struct {
	x, y int
}


type tile struct {
	coord
	id int
}


type board struct {
	extent coord
	tiles [][]int
	ball *tile
	paddle *tile
	score int
}


func newBoard(x, y int) *board {
	cols := make([][]int, x)
	for col := 0; col < len(cols); col++ {
		cols[col] = make([]int, y)
	}
	return &board{
		extent: coord{x, y},
		tiles:  cols,
	}
}
 
func (b *board) render(o io.Writer) {

	for y := 0; y < b.extent.y; y++ {
		buf := strings.Builder{}
		buf.Grow(b.extent.x)
		for x := 0; x < b.extent.x; x++ {
			buf.WriteByte(glyphs[b.tiles[x][y]])
		}
		fmt.Fprintln(o, buf.String())
	}
}

func (b *board) read(in chan int) {
	for t := readTile(in); t != nil; t = readTile(in) {
		if t.x < 0 && t.y == 0 {
			b.score = t.id
			continue
		}
		b.tiles[t.x][t.y] = t.id
		if t.id == Ball {
			b.ball = t
		}
		if t.id == Paddle {
			b.paddle = t
		}
	}
}
