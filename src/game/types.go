package game

import (
	"image"
	"time"
)

type Cell struct {
	X, Y int
}

type Game struct {
	cells          map[Cell]int
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
		cells:        make(map[Cell]int),
		currentState: Conductor,
		speed:        1.0,
		scale:        16.0,
		offsetX:      0,
		offsetY:      0,
		startButton:  image.Rect(10, 300, 190, 340),
		saveButton:   image.Rect(10, 400, 190, 440),
		loadButton:   image.Rect(10, 450, 190, 490),
		stateButtons: []image.Rectangle{
			image.Rect(10, 10, 190, 50),
			image.Rect(10, 60, 190, 100),
			image.Rect(10, 110, 190, 150),
			image.Rect(10, 160, 190, 200),
		},
		slider:     image.Rect(10, 250, 190, 270),
		sliderPos:  0.0,
		lastUpdate: time.Now(),
		isDrawing:  false,
		lastCellX:  0,
		lastCellY:  0,
	}
}
