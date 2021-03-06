package main

import (
	"fmt"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/day15"
	"github.com/mccormickmichael/adventofcode-go/v19/internal/maze"
	"os"
)

func main() {
	testMaze := maze.NewMaze(0, 0, 4, 4)
	testMaze.Render(os.Stdout)

	tiny := day15.Tiny(testMaze)
	mapper := day15.NewMapper(testMaze, tiny)

	mapper.Start(tiny.Start)
	if err := mapper.Map(); err != nil {
		_, _ = fmt.Fprint(os.Stdout, err)
	}

	testMaze.Render(os.Stdout)
}