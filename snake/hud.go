package snake

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

type HUD struct {
	scoreFont font.Face
	scoreText string
}

func NewHud() *HUD {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	scoreFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingVertical,
	})

	if err != nil {
		log.Fatal(err)
	}

	return &HUD{
		scoreFont: scoreFont,
		scoreText: "Score: 0",
	}
}

func (h *HUD) Draw(screen *ebiten.Image) {
	text.Draw(screen, h.scoreText, h.scoreFont, 5, 25, color.White)
}

func (h *HUD) UpdateScoreText(score int) {
	h.scoreText = fmt.Sprintf("Score: %d", score)
}
