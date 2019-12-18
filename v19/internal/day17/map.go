package day17

type dir int

const (
	up    dir = 0
	right dir = 1
	down  dir = 2
	left  dir = 3
)

type coordinate struct {
	x, y int
}

type robot struct {
	x,y  int
	dir  dir
	dead bool
}

type scaffold struct {
	cells [][]byte  // row, col
	robot *robot
}

func (s *scaffold) extent() (cols, rows int) {
	rows = len(s.cells)
	cols = 0
	for _, row := range s.cells {
		if len(row) > cols {
			cols = len(row)
		}
	}
	return
}

func newScaffold(r *robot) *scaffold {
	cells := make([][]byte, 1)
	cells[0] = make([]byte, 0)
	return &scaffold{cells:cells, robot:r}
}

func (s *scaffold) square() {
	cols, _ := s.extent()
	for r, row := range s.cells {
		if len(row) < cols {
			for i := len(row); i < cols; i++ {
				s.cells[r] = append(s.cells[r], '.')
			}
		}
	}
}

func (s *scaffold) intersections() []coordinate {
	intersections := make([]coordinate, 0)
	
	cols, rows := s.extent()
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if s.isIntersection(r, c) {
				intersections = append(intersections, coordinate{c,r})
			}
		}
	}
	return intersections
}

func (s *scaffold) isIntersection(row, col int) bool {
	if s.cells[row][col] != '#' {
		return false
	}
	if s.cells[row-1][col] != '#' {
		return false
	}
	if s.cells[row+1][col] != '#' {
		return false
	}
	if s.cells[row][col-1] != '#' {
		return false
	}
	if s.cells[row][col+1] != '#' {
		return false
	}
	return true
}


type builder struct {
	row, col int
	scaffold *scaffold
}

func (b *builder) Output(o int) {
	b.set(byte(o))
}

func (b *builder) Close() {
}

func (b *builder) set(c byte) {
	switch c {
	case '.', '#':
		b.col++
		b.scaffold.cells[b.row] = append(b.scaffold.cells[b.row], c)
	case '\n':
		b.row++
		b.scaffold.cells = append(b.scaffold.cells, make([]byte, 0))
	case '^':
		b.col++
		b.scaffold.cells[b.row] = append(b.scaffold.cells[b.row], '^')
		b.scaffold.robot = &robot{x:b.col, y:b.row, dir:up}
	case '>':
		b.col++
		b.scaffold.cells[b.row] = append(b.scaffold.cells[b.row], '>')
		b.scaffold.robot = &robot{x:b.col, y:b.row, dir:right}
	case 'v':
		b.col++
		b.scaffold.cells[b.row] = append(b.scaffold.cells[b.row], 'v')
		b.scaffold.robot = &robot{x:b.col, y:b.row, dir:down}
	case '<':
		b.col++
		b.scaffold.cells[b.row] = append(b.scaffold.cells[b.row], '<')
		b.scaffold.robot = &robot{x:b.col, y:b.row, dir:left}
	case 'X':
		b.col++
		b.scaffold.cells[b.row] = append(b.scaffold.cells[b.row], 'X')
		b.scaffold.robot = &robot{x:b.col, y:b.row, dir:up, dead:true}
	}
}
			
