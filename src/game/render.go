package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"math"
)

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, PanelWidth, ScreenHeight, PanelBackground)
	g.DrawPanel(screen)
	g.DrawGrid(screen)
	g.DrawCells(screen)
}

func (g *Game) DrawPanel(screen *ebiten.Image) {
	for i, btn := range g.stateButtons {
		col := StateButtonColors[i]
		if i == g.currentState {
			ebitenutil.DrawRect(screen, float64(btn.Min.X), float64(btn.Min.Y), float64(btn.Dx()), float64(btn.Dy()), ButtonBorderActive)
			ebitenutil.DrawRect(screen, float64(btn.Min.X)+2, float64(btn.Min.Y)+2, float64(btn.Dx())-4, float64(btn.Dy())-4, col)
		} else {
			ebitenutil.DrawRect(screen, float64(btn.Min.X), float64(btn.Min.Y), float64(btn.Dx()), float64(btn.Dy()), col)
		}
	}

	// Start/Stop button
	btnColor := StartButtonBlue
	if g.running {
		btnColor = StartButtonRed
	}
	ebitenutil.DrawRect(screen, float64(g.startButton.Min.X), float64(g.startButton.Min.Y), float64(g.startButton.Dx()), float64(g.startButton.Dy()), btnColor)
	text.Draw(screen, "Start/Stop", basicfont.Face7x13, g.startButton.Min.X+20, g.startButton.Min.Y+20, color.White)

	// Save button
	ebitenutil.DrawRect(screen, float64(g.saveButton.Min.X), float64(g.saveButton.Min.Y), float64(g.saveButton.Dx()), float64(g.saveButton.Dy()), ActionButtonGreen)
	text.Draw(screen, "Save", basicfont.Face7x13, g.saveButton.Min.X+60, g.saveButton.Min.Y+20, color.White)

	// Load button
	ebitenutil.DrawRect(screen, float64(g.loadButton.Min.X), float64(g.loadButton.Min.Y), float64(g.loadButton.Dx()), float64(g.loadButton.Dy()), ActionButtonGreen)
	text.Draw(screen, "Load", basicfont.Face7x13, g.loadButton.Min.X+60, g.loadButton.Min.Y+20, color.White)

	// Slider
	ebitenutil.DrawRect(screen, float64(g.slider.Min.X), float64(g.slider.Min.Y), float64(g.slider.Dx()), float64(g.slider.Dy()), SliderBackground)
	sliderHandleX := float64(g.slider.Min.X) + g.sliderPos
	ebitenutil.DrawRect(screen, sliderHandleX, float64(g.slider.Min.Y), 4, float64(g.slider.Dy()), SliderHandle)
}

func (g *Game) DrawGrid(screen *ebiten.Image) {
	minX := (0 - g.offsetX) / g.scale
	maxX := (float64(ScreenWidth-PanelWidth) - g.offsetX) / g.scale
	minY := (0 - g.offsetY) / g.scale
	maxY := (float64(ScreenHeight) - g.offsetY) / g.scale

	startX := int(math.Floor(minX))
	endX := int(math.Ceil(maxX))
	startY := int(math.Floor(minY))
	endY := int(math.Ceil(maxY))

	for x := startX; x <= endX; x++ {
		sx := PanelWidth + g.offsetX + float64(x)*g.scale
		if sx >= PanelWidth {
			ebitenutil.DrawLine(screen, sx, 0, sx, float64(ScreenHeight), GridLines)
		}
	}

	for y := startY; y <= endY; y++ {
		sy := g.offsetY + float64(y)*g.scale
		ebitenutil.DrawLine(screen, PanelWidth, sy, float64(ScreenWidth), sy, GridLines)
	}
}

func (g *Game) DrawCells(screen *ebiten.Image) {
	colors := map[int]color.RGBA{
		Empty:        {0, 0, 0, 255},
		Conductor:    {255, 255, 0, 255},
		ElectronHead: {0, 0, 255, 255},
		ElectronTail: {255, 0, 0, 255},
	}

	minX := (0 - g.offsetX) / g.scale
	maxX := (float64(ScreenWidth-PanelWidth) - g.offsetX) / g.scale
	minY := (0 - g.offsetY) / g.scale
	maxY := (float64(ScreenHeight) - g.offsetY) / g.scale

	startX := int(math.Floor(minX))
	endX := int(math.Ceil(maxX))
	startY := int(math.Floor(minY))
	endY := int(math.Ceil(maxY))

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			cell := Cell{x, y}
			if state, ok := g.cells[cell]; ok {
				sx := PanelWidth + g.offsetX + float64(x)*g.scale
				sy := g.offsetY + float64(y)*g.scale
				if sx >= PanelWidth {
					ebitenutil.DrawRect(screen, sx, sy, g.scale, g.scale, colors[state])
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
