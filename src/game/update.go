package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"math"
	"time"
)

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		x, y := ebiten.CursorPosition()
		if g.prevMouseX != 0 || g.prevMouseY != 0 {
			dx := float64(x - g.prevMouseX)
			dy := float64(y - g.prevMouseY)
			g.offsetX += dx
			g.offsetY += dy
		}
		g.prevMouseX, g.prevMouseY = x, y
	} else {
		g.prevMouseX, g.prevMouseY = 0, 0
	}

	_, wy := ebiten.Wheel()
	if wy != 0 {
		x, y := ebiten.CursorPosition()
		oldScale := g.scale
		g.scale *= math.Pow(1.1, wy)
		if g.scale < 4 {
			g.scale = 4
		} else if g.scale > 64 {
			g.scale = 64
		}
		if x >= PanelWidth {
			gameX := x - PanelWidth
			fx := float64(gameX)
			fy := float64(y)
			g.offsetX = fx - (fx-g.offsetX)*(g.scale/oldScale)
			g.offsetY = fy - (fy-g.offsetY)*(g.scale/oldScale)
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x < PanelWidth {
			for i, btn := range g.stateButtons {
				if image.Pt(x, y).In(btn) {
					g.currentState = i
				}
			}
			if image.Pt(x, y).In(g.startButton) {
				g.running = !g.running
			}
			if image.Pt(x, y).In(g.saveButton) {
				go g.SaveWithDialog()
			}
			if image.Pt(x, y).In(g.loadButton) {
				go g.LoadWithDialog()
			}
			if image.Pt(x, y).In(g.slider) {
				g.draggingSlider = true
			}
		} else {
			g.isDrawing = true
			cellX, cellY := g.screenToCell(x, y)
			g.lastCellX, g.lastCellY = cellX, cellY
			if g.currentState == Empty {
				delete(g.cells, Cell{cellX, cellY})
			} else {
				g.cells[Cell{cellX, cellY}] = g.currentState
			}
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.isDrawing {
		x, y := ebiten.CursorPosition()
		if x >= PanelWidth {
			cellX, cellY := g.screenToCell(x, y)
			if cellX != g.lastCellX || cellY != g.lastCellY {
				g.drawLine(g.lastCellX, g.lastCellY, cellX, cellY)
				g.lastCellX, g.lastCellY = cellX, cellY
			}
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.isDrawing = false
	}

	if g.draggingSlider {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, _ := ebiten.CursorPosition()
			if x < 10 {
				x = 10
			} else if x > 190 {
				x = 190
			}
			g.sliderPos = float64(x - 10)
			g.speed = (g.sliderPos/180.0)*9.0 + 1.0
		} else {
			g.draggingSlider = false
		}
	}

	if g.running {
		now := time.Now()
		elapsed := now.Sub(g.lastUpdate).Seconds()
		if elapsed >= 1.0/g.speed {
			g.UpdateCells()
			g.lastUpdate = now
		}
	}

	return nil
}

func (g *Game) UpdateCells() {
	newCells := make(map[Cell]int)
	temp := make(map[Cell]int)

	for cell := range g.cells {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				temp[Cell{cell.X + dx, cell.Y + dy}] = 0
			}
		}
	}

	for cell := range temp {
		currentState := g.cells[cell]
		switch currentState {
		case Empty:
		case ElectronHead:
			newCells[cell] = ElectronTail
		case ElectronTail:
			newCells[cell] = Conductor
		case Conductor:
			count := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					neighbor := Cell{cell.X + dx, cell.Y + dy}
					if g.cells[neighbor] == ElectronHead {
						count++
					}
				}
			}
			if count == 1 || count == 2 {
				newCells[cell] = ElectronHead
			} else {
				newCells[cell] = Conductor
			}
		}
	}

	g.cells = newCells
}
