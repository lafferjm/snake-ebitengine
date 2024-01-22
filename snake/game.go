package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	x float32
	y float32
}

func NewGame() (*Game, error) {
	g := &Game{
		x: ScreenWidth / 2,
		y: ScreenHeight / 2,
	}

	return g, nil
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
	return ScreenWidth, ScreenHeight
}
