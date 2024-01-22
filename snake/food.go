package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
)

type Food struct {
	position Vec2
}

func NewFood() *Food {
	var initialX = rand.Intn(ScreenWidth/20-20) * 20
	var initialY = rand.Intn(ScreenHeight/20-20) * 20

	return &Food{
		position: Vec2{x: initialX, y: initialY},
	}
}

func (f *Food) Update() {
	var newX = rand.Intn(ScreenWidth/20-20) * 20
	var newY = rand.Intn(ScreenHeight/20-20) * 20
	f.position = Vec2{
		x: newX,
		y: newY,
	}
}

func (f *Food) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(f.position.x),
		float32(f.position.y),
		20, 20, color.RGBA{R: 0xFF, A: 0xFF}, true)
}
