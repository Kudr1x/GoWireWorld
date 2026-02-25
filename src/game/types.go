package game

import (
	"goWireWorld/src/core"
	"image"
	"sync"
	"time"
)

type Game struct {
	mu sync.RWMutex

	cells          map[core.Cell]int
	currentState   int
	running        bool
	speed          float64
	scale          float64
	offsetX        float64
	offsetY        float64
	prevMouseX     int
	prevMouseY     int
	startButton    image.Rectangle
	saveButton     image.Rectangle
	loadButton     image.Rectangle
	stateButtons   []image.Rectangle
	slider         image.Rectangle
	sliderPos      float64
	draggingSlider bool
	lastUpdate     time.Time
	isDrawing      bool
	lastCellX      int
	lastCellY      int
}

func NewGame() *Game {
	return &Game{
		cells:        make(map[core.Cell]int),
		currentState: core.Conductor,
		speed:        1.0,
		scale:        16.0,
		offsetX:      0,
		offsetY:      0,
		startButton:  image.Rect(UIButtonX, UIStartY, UIButtonX+UIButtonWidth, UIStartY+UIButtonHeight),
		saveButton:   image.Rect(UIButtonX, UISaveY, UIButtonX+UIButtonWidth, UISaveY+UIButtonHeight),
		loadButton:   image.Rect(UIButtonX, UILoadY, UIButtonX+UIButtonWidth, UILoadY+UIButtonHeight),
		stateButtons: []image.Rectangle{
			image.Rect(UIButtonX, UIStateStartY, UIButtonX+UIButtonWidth, UIStateStartY+UIButtonHeight),
			image.Rect(UIButtonX, UIStateStartY+UIStateGap, UIButtonX+UIButtonWidth, UIStateStartY+UIStateGap+UIButtonHeight),
			image.Rect(UIButtonX, UIStateStartY+UIStateGap*2, UIButtonX+UIButtonWidth, UIStateStartY+UIStateGap*2+UIButtonHeight),
			image.Rect(UIButtonX, UIStateStartY+UIStateGap*3, UIButtonX+UIButtonWidth, UIStateStartY+UIStateGap*3+UIButtonHeight),
		},
		slider:     image.Rect(UIButtonX, UISliderY, UIButtonX+UIButtonWidth, UISliderY+UISliderHeight),
		sliderPos:  0.0,
		lastUpdate: time.Now(),
		isDrawing:  false,
	}
}
