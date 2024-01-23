package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Vec2 struct {
	x int
	y int
}

type GameText struct {
	fontFace  font.Face
	textValue string
	position  Vec2
}

type Game struct {
	snake     *Snake
	food      *Food
	hud       *HUD
	gameState *GameState
	score     int
}

func NewGame() (*Game, error) {
	g := &Game{
		snake:     NewSnake(),
		food:      NewFood(),
		hud:       NewHud(),
		gameState: NewGameState(),
		score:     0,
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
	if g.gameState.PollState(MainMenu) {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.gameState.SetState(Playing)
		}
	}

	if g.gameState.PollState(Playing) {
		g.snake.Update()

		if g.CheckCollision() {
			g.food.Update()
			g.snake.Grow()
			g.score += 1
			g.hud.UpdateScoreText(g.score)
		}

		if g.CheckGameOver() {
			g.snake.Stop()
			g.gameState.SetState(GameOver)
		}
	}

	if g.gameState.PollState(GameOver) {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.snake = NewSnake()
			g.food = NewFood()
			g.hud = NewHud()
			g.score = 0

			g.gameState.SetState(Playing)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameState.PollState(MainMenu) {
		mainGameScreen.Draw(screen)
	}

	if g.gameState.PollState(Playing) {
		g.hud.Draw(screen)
		g.food.Draw(screen)
		g.snake.Draw(screen)
	}

	if g.gameState.PollState(GameOver) {
		gameOverScreen.Draw(screen)
	}

}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
