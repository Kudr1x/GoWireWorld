package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	PanelWidth   = 200

	UIButtonX      = 10
	UIButtonWidth  = 180
	UIButtonHeight = 40
	UIStateStartY  = 10
	UIStateGap     = 50
	UISliderY      = 250
	UISliderHeight = 20
	UIStartY       = 300
	UISaveY        = 400
	UILoadY        = 450
)

var (
	PanelBackground    = color.RGBA{R: 50, G: 50, B: 50, A: 255}
	ButtonBorderActive = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	SliderBackground   = color.RGBA{R: 100, G: 100, B: 100, A: 255}
	SliderHandle       = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	StartButtonBlue    = color.RGBA{G: 128, B: 255, A: 255}
	StartButtonRed     = color.RGBA{R: 255, A: 255}
	ActionButtonGreen  = color.RGBA{G: 255, A: 255}
	GridLines          = color.RGBA{R: 50, G: 50, B: 50, A: 50}
)

var (
	ColorEmpty        = color.RGBA{A: 255}
	ColorConductor    = color.RGBA{R: 255, G: 255, A: 255}
	ColorElectronHead = color.RGBA{B: 255, A: 255}
	ColorElectronTail = color.RGBA{R: 255, A: 255}
)

var StateButtonColors = []color.RGBA{
	ColorEmpty,
	ColorConductor,
	ColorElectronHead,
	ColorElectronTail,
}

var cellTile *ebiten.Image

func GetCellTile() *ebiten.Image {
	if cellTile == nil {
		cellTile = ebiten.NewImage(1, 1)
		cellTile.Fill(color.White)
	}
	return cellTile
}
