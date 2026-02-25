package game

import (
	"goWireWorld/src/core"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

var basicFace = text.NewGoXFace(basicfont.Face7x13)

func drawText(screen *ebiten.Image, str string, x, y float64, clr color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(clr)
	text.Draw(screen, str, basicFace, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, 0, 0, PanelWidth, ScreenHeight, PanelBackground, false)
	g.DrawPanel(screen)
	g.DrawGrid(screen)
	g.DrawCells(screen)
}

func (g *Game) DrawPanel(screen *ebiten.Image) {
	for i, btn := range g.stateButtons {
		col := StateButtonColors[i]
		if i == g.currentState {
			vector.FillRect(screen, float32(btn.Min.X), float32(btn.Min.Y), float32(btn.Dx()), float32(btn.Dy()), ButtonBorderActive, false)
			vector.FillRect(screen, float32(btn.Min.X)+2, float32(btn.Min.Y)+2, float32(btn.Dx())-4, float32(btn.Dy())-4, col, false)
		} else {
			vector.FillRect(screen, float32(btn.Min.X), float32(btn.Min.Y), float32(btn.Dx()), float32(btn.Dy()), col, false)
		}
	}

	btnColor := StartButtonBlue
	if g.running {
		btnColor = StartButtonRed
	}
	vector.FillRect(screen, float32(g.startButton.Min.X), float32(g.startButton.Min.Y), float32(g.startButton.Dx()), float32(g.startButton.Dy()), btnColor, false)
	drawText(screen, "Start/Stop", float64(g.startButton.Min.X+20), float64(g.startButton.Min.Y+20), color.White)

	vector.FillRect(screen, float32(g.saveButton.Min.X), float32(g.saveButton.Min.Y), float32(g.saveButton.Dx()), float32(g.saveButton.Dy()), ActionButtonGreen, false)
	drawText(screen, "Save", float64(g.saveButton.Min.X+60), float64(g.saveButton.Min.Y+20), color.White)

	vector.FillRect(screen, float32(g.loadButton.Min.X), float32(g.loadButton.Min.Y), float32(g.loadButton.Dx()), float32(g.loadButton.Dy()), ActionButtonGreen, false)
	drawText(screen, "Load", float64(g.loadButton.Min.X+60), float64(g.loadButton.Min.Y+20), color.White)

	vector.FillRect(screen, float32(g.slider.Min.X), float32(g.slider.Min.Y), float32(g.slider.Dx()), float32(g.slider.Dy()), SliderBackground, false)
	sliderHandleX := float32(g.slider.Min.X) + float32(g.sliderPos)
	vector.FillRect(screen, sliderHandleX, float32(g.slider.Min.Y), 4, float32(g.slider.Dy()), SliderHandle, false)
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
			vector.StrokeLine(screen, float32(sx), 0, float32(sx), float32(ScreenHeight), 1, GridLines, false)
		}
	}

	for y := startY; y <= endY; y++ {
		sy := g.offsetY + float64(y)*g.scale
		vector.StrokeLine(screen, float32(PanelWidth), float32(sy), float32(ScreenWidth), float32(sy), 1, GridLines, false)
	}
}

func (g *Game) DrawCells(screen *ebiten.Image) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	colors := map[int]color.RGBA{
		core.Empty:        ColorEmpty,
		core.Conductor:    ColorConductor,
		core.ElectronHead: ColorElectronHead,
		core.ElectronTail: ColorElectronTail,
	}

	minX := (0 - g.offsetX) / g.scale
	maxX := (float64(ScreenWidth-PanelWidth) - g.offsetX) / g.scale
	minY := (0 - g.offsetY) / g.scale
	maxY := (float64(ScreenHeight) - g.offsetY) / g.scale

	startX := int(math.Floor(minX))
	endX := int(math.Ceil(maxX))
	startY := int(math.Floor(minY))
	endY := int(math.Ceil(maxY))

	op := &ebiten.DrawImageOptions{}

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			cell := core.Cell{X: x, Y: y}
			if state, ok := g.cells[cell]; ok {
				sx := PanelWidth + g.offsetX + float64(x)*g.scale
				sy := g.offsetY + float64(y)*g.scale
				if sx >= PanelWidth {
					op.GeoM.Reset()
					op.GeoM.Scale(g.scale, g.scale)
					op.GeoM.Translate(sx, sy)
					op.ColorScale.Reset()
					op.ColorScale.ScaleWithColor(colors[state])
					screen.DrawImage(GetCellTile(), op)
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
