package util

type Point struct {
	Y int
	X int
}

func (p *Point) Add(o Point) Point {
	return Point{Y: p.Y + o.Y, X: p.X + o.X}
}

func (p *Point) Subtract(o Point) Point {
	return Point{Y: p.Y - o.Y, X: p.X - o.X}
}

func (p *Point) IsInBounds(height, width int) bool {
	if p.Y < 0 || p.Y >= height {
		return false
	}
	if p.X < 0 || p.X >= width {
		return false
	}
	return true
}

func ByteAt(lines [][]byte, p Point) (byte, bool) {
	if p.Y < 0 || p.X < 0 {
		return 0, false
	}
	if p.Y >= len(lines) {
		return 0, false
	}
	if p.X >= len(lines[p.Y]) {
		return 0, false
	}
	return lines[p.Y][p.X], true
}
