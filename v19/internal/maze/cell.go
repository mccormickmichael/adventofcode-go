package maze

type Cell struct {
	Id          string
	celler      Locator
	loc         Coord
	Traversable bool
	Explored    bool
}

func NewCell(id string, celler Locator, loc Coord, traversable bool) *Cell {
	return &Cell{id, celler, loc, traversable, !traversable}
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
		c.celler.At(c.loc.Up()),
		c.celler.At(c.loc.Right()),
		c.celler.At(c.loc.Down()),
		c.celler.At(c.loc.Left()),
	}
}