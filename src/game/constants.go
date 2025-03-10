package game

import "image/color"

const (
	Empty = iota
	Conductor
	ElectronHead
	ElectronTail
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	PanelWidth   = 200
)

var (
	PanelBackground    = color.RGBA{50, 50, 50, 255}
	ButtonBorderActive = color.RGBA{255, 255, 255, 255}
	SliderBackground   = color.RGBA{100, 100, 100, 255}
	SliderHandle       = color.RGBA{255, 255, 255, 255}
	StartButtonBlue    = color.RGBA{0, 128, 255, 255}
	StartButtonRed     = color.RGBA{255, 0, 0, 255}
	ActionButtonGreen  = color.RGBA{0, 255, 0, 255}
	GridLines          = color.RGBA{50, 50, 50, 50}
)

var (
	ColorEmpty        = color.RGBA{0, 0, 0, 255}
	ColorConductor    = color.RGBA{255, 255, 0, 255}
	ColorElectronHead = color.RGBA{0, 0, 255, 255}
	ColorElectronTail = color.RGBA{255, 0, 0, 255}
)

var StateButtonColors = []color.RGBA{
	ColorEmpty,
	ColorConductor,
	ColorElectronHead,
	ColorElectronTail,
}
