package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Vec2 struct {
	x int
	y int
}

type Game struct {
	snake *Snake
}

func NewGame() (*Game, error) {
	g := &Game{
		snake: NewSnake(),
	}

	return g, nil
}

func (g *Game) Update() error {
	g.snake.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
