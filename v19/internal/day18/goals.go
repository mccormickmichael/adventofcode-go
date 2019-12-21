package day18

import mz "github.com/mccormickmichael/adventofcode-go/v19/internal/maze"

type goal struct {
	key  Key
	dist int
}

type goals struct {
	maze    *mz.Maze
	pending map[mz.Coord]int
	marked  map[mz.Coord]int
	found   []goal
	open    map[mz.Coord]Key
}

func newGoals(maze *mz.Maze, scenario *scenario) *goals {
	openKeys := map[mz.Coord]Key{}
	for _, k := range scenario.outstandingKeys {
		openKeys[k.loc] = k
	}
	return &goals{
		maze:maze,
		pending: map[mz.Coord]int{scenario.loc:0},
		marked:  make(map[mz.Coord]int),
		found:   make([]goal, 0),
		open:    openKeys,
	}
}

func (g *goals) find() {

	for len(g.pending) > 0 && len(g.open) > 0 {
		newp := make(map[mz.Coord]int, 0)

		for loc, dist := range g.pending {
			for _, cell := range g.maze.At(loc).TraversableNeighbors() {
				cloc := cell.Loc()
				if md, ok := g.marked[cloc]; ok {
					g.shortcut(cloc, md, dist + 1)
				} else {
					if !g.markGoal(cloc, dist + 1) {
						newp[cloc] = dist + 1
					}
				}
			}
			g.marked[loc] = dist
		}
		g.pending = newp
	}
}

func (g *goals) shortcut(loc mz.Coord, dist, newDist int) {
	if newDist < dist {
		g.marked[loc] = newDist
	}
}

func (g *goals) markGoal(loc mz.Coord, dist int) bool {
	if maybeGoal, ok := g.open[loc]; ok {
		g.found = append(g.found, goal{maybeGoal, dist})
		delete(g.open, loc)
		return true
	}
	return false
}