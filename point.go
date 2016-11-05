package main

type point struct {
	X, Y int
}

func (p point) Translate(dx, dy int) point {
	return point{p.X + dx, p.Y + dy}
}
