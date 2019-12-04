package event

type Point struct {
	X, Y int
}

func (p Point) Magnitude() int {
	return abs(p.X) + abs(p.Y)
}

func (p Point) Distance(o Point) int {
	return abs(p.X - o.X) + abs(p.Y - o.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p Point) Offset(x, y int) Point {
	return Point{p.X + x, p.Y + y}
}

func EqualPoints(lhs, rhs PointSlice) bool {
	if len(lhs) != len(rhs) { return false}
	for i, p := range lhs {
		if rhs[i] != p {
			return false
		}
	}
	return true
}

type PointSlice []Point

func (p PointSlice) Len() int {
	return len(p)
}

func (p PointSlice) Less(i, j int) bool {
	iDist := p[i].Magnitude()
	jDist := p[j].Magnitude()
	return iDist < jDist
}

func (p PointSlice) Swap(i, j int) {
	t := p[i]
	p[i] = p[j]
	p[j] = t
}

type Segment struct {
	Head, Tail Point
}

func (s Segment) Length() int {
	return s.Head.Distance(s.Tail)
}

func Intersections(a, b Segment) []Point {
	if a.H() && b.V() {
		if intersects, p := orthoIntersect(a.norm(), b.norm()); intersects {
			return []Point{p}
		}
		return nil
	}
	if a.V() && b.H() {
		return Intersections(b, a)
	}
	if a.H() && b.H() {
		return hIntersects(a.norm(), b.norm())
	}
	if a.V() && b.V() {
		return vIntersects(a.norm(), b.norm())
	}
	return nil
}

func (s Segment) H() bool {
	return s.Head.Y == s.Tail.Y
}
func (s Segment) V() bool {
	return s.Head.X == s.Tail.X
}

func hIntersects(a, b Segment) []Point {
	if a.Head.X > b.Head.X { return hIntersects(b, a) }
	if a.Head.Y != b.Head.Y { return nil }
	points := make([]Point, 0)
	for i := b.Head.X; i <= a.Tail.X; i++ {
		points = append(points, Point{i, a.Head.Y})
	}
	return points
}

func vIntersects(a, b Segment) []Point {
	if a.Head.Y > b.Head.Y { return vIntersects(b, a) }
	if a.Head.X != b.Head.X { return nil }
	points := make([]Point, 0)
	for i := b.Head.Y; i <= a.Tail.Y; i++ {
		points = append(points, Point{a.Head.X, i})
	}
	return points
}

func orthoIntersect(h, v Segment) (bool, Point) {
	if (h.Head.X <= v.Head.X && v.Head.X <= h.Tail.X) &&
		(v.Head.Y <= h.Head.Y && h.Head.Y <= v.Tail.Y) {
		return true, Point{v.Head.X, h.Head.Y}
	}
	return false, Point{}
}

func (s Segment) norm() Segment {
	if s.H() {
		if s.Head.X > s.Tail.X {
			return s.swap()
		}
	}
	if s.Head.Y > s.Tail.Y {
		return s.swap()
	}
	return s
}

func (s Segment) swap() Segment {
	return Segment{s.Tail, s.Head}
}