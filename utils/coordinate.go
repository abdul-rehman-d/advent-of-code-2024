package utils

type Coord struct {
	Row int
	Col int
}

func (c Coord) Plus(d Coord) Coord {
	return Coord{c.Row + d.Row, c.Col + d.Col}
}

func GetDirections() []Coord {
	return []Coord{
		{-1, 0}, // up
		{1, 0},  // down
		{0, 1},  // right
		{0, -1}, // left
	}
}
