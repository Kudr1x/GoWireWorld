package game

import (
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) handlePanning() {
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
}

func (g *Game) handleZooming() {
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
}

func (g *Game) handleMouseClicks() {
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
				g.running = false
				go g.SaveWithDialog()
			}
			if image.Pt(x, y).In(g.loadButton) {
				g.running = false
				go g.LoadWithDialog()
			}
			if image.Pt(x, y).In(g.slider) {
				g.draggingSlider = true
			}
		} else {
			g.isDrawing = true
			cellX, cellY := g.screenToCell(x, y)
			g.lastCellX, g.lastCellY = cellX, cellY

			g.mu.Lock()
			if g.currentState == Empty {
				delete(g.cells, Cell{cellX, cellY})
			} else {
				g.cells[Cell{cellX, cellY}] = g.currentState
			}
			g.mu.Unlock()
		}
	}
}

func (g *Game) handleDrawing() {
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
}

func (g *Game) handleSlider() {
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
}

func (g *Game) handleSimulation() {
	if g.running {
		now := time.Now()
		elapsed := now.Sub(g.lastUpdate).Seconds()
		if elapsed >= 1.0/g.speed {
			g.UpdateCells()
			g.lastUpdate = now
		}
	}
}

func (g *Game) Update() error {
	g.handlePanning()
	g.handleZooming()
	g.handleMouseClicks()
	g.handleDrawing()
	g.handleSlider()
	g.handleSimulation()
	return nil
}

func (g *Game) UpdateCells() {
	g.mu.RLock()
	newCells := CalculateNextState(g.cells)
	g.mu.RUnlock()

	g.mu.Lock()
	g.cells = newCells
	g.mu.Unlock()
}
