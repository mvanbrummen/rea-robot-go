package main

type table struct {
	min, max point
}

func (t table) withinBounds(p point) bool {
	if p.X < t.min.X || p.X > t.max.X {
		return false
	}
	if p.Y < t.min.Y || p.Y > t.max.Y {
		return false
	}
	return true
}
