package maze

type Cell struct {
	id          string
	celler      Celler
	loc         Coord
	Traversable bool
	Explored    bool
}

func NewCell(id string, celler Celler, loc Coord, traversable bool) *Cell {
	return &Cell{id, celler, loc, traversable, !traversable}
}

func (c *Cell) String() string {
	if c == nil {
		return " "
	}
	return c.id
}

func (c *Cell) Id() string {
	return c.id
}

func (c *Cell) Loc() Coord {
	return c.loc
}

func (c *Cell) Neighbors() []*Cell {
	return []*Cell{
		c.celler.At(c.loc.X,   c.loc.Y-1),
		c.celler.At(c.loc.X+1, c.loc.Y),
		c.celler.At(c.loc.X,   c.loc.Y+1),
		c.celler.At(c.loc.X-1, c.loc.Y),
	}
}