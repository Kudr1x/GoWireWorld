package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goWireWorld/src/game"
	"log"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("goWireWorld")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
