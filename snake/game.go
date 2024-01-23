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

func (g *Game) CheckCollision() bool {
	snakeHead := g.snake.Head()
	foodPosition := g.food.position

	return snakeHead.x == foodPosition.x && snakeHead.y == foodPosition.y
}

func (g *Game) CheckGameOver() bool {
	return g.snake.IsOffScreen() || g.snake.DidIntersectSelf()
}

func (g *Game) Update() error {
	g.snake.Update()

	if g.CheckCollision() {
		g.food.Update()
		g.snake.Grow()
	}

	if g.CheckGameOver() {
		g.snake.Stop()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.food.Draw(screen)
	g.snake.Draw(screen)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
