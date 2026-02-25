package main

import (
	"goWireWorld/src/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("goWireWorld")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
