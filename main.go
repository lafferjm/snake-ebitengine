package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	x float32
	y float32
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.x -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.x += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.y += 1
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, g.x, g.y, 20, 20, color.RGBA{G: 0xFF, A: 0xFF}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{x: 20, y: 20}); err != nil {
		log.Fatal(err)
	}
}
