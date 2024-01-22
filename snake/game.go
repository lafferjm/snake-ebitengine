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
	food  *Food
}

func NewGame() (*Game, error) {
	g := &Game{
		snake: NewSnake(),
		food:  NewFood(),
	}

	return g, nil
}

func (g *Game) Update() error {
	g.snake.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.food.Draw(screen)
	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
