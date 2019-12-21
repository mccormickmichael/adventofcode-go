package maze

type Cell struct {
	Id          string
	locator     Locator
	loc         Coord
	Traversable bool
	Explored    bool
}

func NewCell(id string, locator Locator, loc Coord, traversable bool) *Cell {
	return &Cell{id, locator, loc, traversable, !traversable}
}

func (c *Cell) String() string {
	if c == nil {
		return " "
	}
	return c.Id
}


func (c *Cell) Loc() Coord {
	return c.loc
}

func (c *Cell) Neighbors() []*Cell {
	return []*Cell{
		c.locator.At(c.loc.Up()),
		c.locator.At(c.loc.Right()),
		c.locator.At(c.loc.Down()),
		c.locator.At(c.loc.Left()),
	}
}

func (c * Cell) TraversableNeighbors() []*Cell {
	traversable := make([]*Cell, 0, 4)
	for _, cell := range c.Neighbors() {
		if cell.Traversable {
			traversable = append(traversable, cell)
		}
	}
	return traversable
}