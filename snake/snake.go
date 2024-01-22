package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Direction Vec2

var UP = Direction{0, -1}
var DOWN = Direction{0, 1}
var LEFT = Direction{-1, 0}
var RIGHT = Direction{1, 0}
var STOPPED = Direction{0, 0}

type Snake struct {
	segments  []Vec2
	direction Direction
}

func NewSnake() *Snake {
	var initialSegments []Vec2
	for i := 0; i < 3; i++ {
		initialSegments = append(initialSegments, Vec2{
			400, 300 + (i * 20),
		})
	}

	snake := &Snake{
		segments:  initialSegments,
		direction: STOPPED,
	}

	return snake
}

func (s *Snake) Head() Vec2 {
	return s.segments[0]
}

func (s *Snake) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && s.direction != RIGHT {
		s.direction = LEFT
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && s.direction != LEFT {
		s.direction = RIGHT
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && s.direction != DOWN {
		s.direction = UP
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && s.direction != UP {
		s.direction = DOWN
	}

	if s.direction != STOPPED {
		s.segments = s.segments[:len(s.segments)-1]

		snakeHead := s.Head()
		snakeHead.x += s.direction.x * 20
		snakeHead.y += s.direction.y * 20

		s.segments = append([]Vec2{snakeHead}, s.segments...)
	}
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, segment := range s.segments {
		vector.DrawFilledRect(screen,
			float32(segment.x),
			float32(segment.y),
			20, 20, color.RGBA{G: 0xFF, A: 0xFF}, true)
	}
}
