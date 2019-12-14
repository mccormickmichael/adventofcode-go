package day13

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/intmath"
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
	extent   coord
	tiles    [][]int
	ball     *tile
	paddle   *tile
	score    int
	gameOver bool
	renderer io.Writer
	buf      [2]int
	bufsize  int
}

func (b *board) Input() int {
	if b.paddle != nil && b.ball != nil {
		return intmath.Cmp(b.ball.x, b.paddle.x)
	}
	return 0
}

func (b *board) Output(val int) {

	switch b.bufsize {
	case 0:
		b.buf[0] = val
		b.bufsize = 1
	case 1:
		b.buf[1] = val
		b.bufsize = 2
	case 2:
		t := tile{coord{b.buf[0], b.buf[1]}, val}
		b.bufsize = 0
		b.setTile(&t)
	}
}

func (b *board) Close() {
	b.gameOver = true
	b.render()
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

func (b *board) setTile(t *tile) {
	if t.x < 0 && t.y == 0 {
		b.score = t.id
		// TODO: consider rendering score
		return
	}
	b.tiles[t.x][t.y] = t.id
	if t.id == Ball {
		b.ball = t
	}
	if t.id == Paddle {
		b.paddle = t
	}
}
 
func (b *board) render() {
	o := b.renderer

	for y := 0; y < b.extent.y; y++ {
		buf := strings.Builder{}
		buf.Grow(b.extent.x)
		for x := 0; x < b.extent.x; x++ {
			buf.WriteByte(glyphs[b.tiles[x][y]])
		}
		fmt.Fprintln(o, buf.String())
	}
}
