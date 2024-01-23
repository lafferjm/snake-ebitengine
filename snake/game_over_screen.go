package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

type GameOverScreen struct {
	gameOverText     GameText
	instructionsText GameText
}

func NewGameOverScreen() *GameOverScreen {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	gameOverFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    64,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	instructionsFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	const gameOverText = "Game Over!"
	const instructionsText = "Press <space> to restart!"

	gameOverMeasure := font.MeasureString(gameOverFont, gameOverText) / 148
	instructionsMeasure := font.MeasureString(instructionsFont, instructionsText) / 148

	gameOverX := int(ScreenWidth/2 - gameOverMeasure)
	instructionsX := int(ScreenWidth/2 - instructionsMeasure)

	return &GameOverScreen{
		gameOverText: GameText{
			fontFace:  gameOverFont,
			textValue: gameOverText,
			position: Vec2{
				x: gameOverX,
				y: 250,
			},
		},
		instructionsText: GameText{
			fontFace:  instructionsFont,
			textValue: instructionsText,
			position: Vec2{
				x: instructionsX,
				y: 300,
			},
		},
	}
}

func (ms *GameOverScreen) Draw(screen *ebiten.Image) {
	text.Draw(screen, ms.gameOverText.textValue, ms.gameOverText.fontFace, ms.gameOverText.position.x, ms.gameOverText.position.y, color.White)
	text.Draw(screen, ms.instructionsText.textValue, ms.instructionsText.fontFace, ms.instructionsText.position.x, ms.instructionsText.position.y, color.White)
}

var gameOverScreen = NewGameOverScreen()
