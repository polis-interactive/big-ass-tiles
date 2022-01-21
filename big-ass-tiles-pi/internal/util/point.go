package util

import "math/rand"

type Point struct {
	X int
	Y int
}

func (p *Point) PointIsEqual(p1 *Point) bool {
	return p.X == p1.X && p.Y == p1.Y
}

func GetRandomPoint(g *GridDefinition) *Point {
	return &Point{
		X: rand.Intn(g.Columns),
		Y: rand.Intn(g.Rows+2*g.RowExtension) - 2*g.RowExtension,
	}
}

func (p *Point) GetNextRandomPoint(g *GridDefinition) *Point {
	dir := rand.Intn(4)
	xp := p.X
	yp := p.Y
	if dir == 0 {
		xp -= 1
		if xp < 0 {
			xp = g.Columns - 1
		}
	} else if dir == 1 {
		xp += 1
		if xp >= g.Columns {
			xp = xp - g.Columns
		}
	} else if dir == 2 {
		yp -= 1
		if yp <= -g.RowExtension {
			yp = g.Rows + g.RowExtension - 1
		}
	} else {
		yp += 1
		if yp >= g.Rows+g.RowExtension {
			yp = yp - g.Rows - g.RowExtension
		}
	}
	return &Point{
		X: xp,
		Y: yp,
	}
}
