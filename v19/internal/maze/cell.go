package maze

type Cell struct {
	id          string
	maze        *Maze
	loc         Coord
	Traversable bool
}

func (c *Cell) Id() string {
	return c.id
}

func (c *Cell) Loc() Coord {
	return c.loc
}

func (c *Cell) Neighbors() []*Cell {
	return []*Cell{
		c.maze.At(c.loc.X,   c.loc.Y-1),
		c.maze.At(c.loc.X+1, c.loc.Y),
		c.maze.At(c.loc.X,   c.loc.Y+1),
		c.maze.At(c.loc.X-1, c.loc.Y),
	}
}