package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func New(x int, y int, a uint) Square {
	return Square{Point{x, y}, a}
}

func (s Square) End() Point {
	return Point{s.start.x + int(s.a), s.start.y - int(s.a)}
}

func (s Square) Perimeter() uint {
	return s.a * 4
}

func (s Square) Area() uint {
	return s.a * s.a
}
