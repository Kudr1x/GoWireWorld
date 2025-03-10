package game

import "math"

func (g *Game) drawLine(x0, y0, x1, y1 int) {
	dx := abs(x1 - x0)
	dy := -abs(y1 - y0)
	sx := 1
	if x0 >= x1 {
		sx = -1
	}
	sy := 1
	if y0 >= y1 {
		sy = -1
	}
	err := dx + dy

	for {
		if g.currentState == Empty {
			delete(g.cells, Cell{x0, y0})
		} else {
			g.cells[Cell{x0, y0}] = g.currentState
		}
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (g *Game) screenToCell(x, y int) (int, int) {
	if x < PanelWidth {
		return 0, 0
	}
	gameX := x - PanelWidth
	fx := (float64(gameX) - g.offsetX) / g.scale
	fy := (float64(y) - g.offsetY) / g.scale
	return int(math.Floor(fx)), int(math.Floor(fy))
}
