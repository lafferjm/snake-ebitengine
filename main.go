package main

import (
	"log"
	"snake-ebitengine/snake"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := snake.NewGame()
	if err != nil {
		log.Fatal("Failed to initialize game")
	}

	ebiten.SetWindowSize(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowTitle("Snake")
	ebiten.SetTPS(30)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
