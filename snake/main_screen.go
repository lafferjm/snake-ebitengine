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

type MainGameScreen struct {
	titleText        GameText
	instructionsText GameText
}

func NewMainGameScreen() *MainGameScreen {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	titleFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
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

	const titleText = "Snake"
	const instructionsText = "Press <space> to play!"

	titleMeasure := font.MeasureString(titleFont, titleText) / 148 // dpi / 2
	instructionsMeasure := font.MeasureString(instructionsFont, instructionsText) / 148

	titleX := int(ScreenWidth/2 - titleMeasure)
	instructionsX := int(ScreenWidth/2 - instructionsMeasure)

	return &MainGameScreen{
		titleText: GameText{
			textValue: titleText,
			fontFace:  titleFont,
			position: Vec2{
				x: titleX,
				y: 250,
			},
		},
		instructionsText: GameText{
			textValue: instructionsText,
			fontFace:  instructionsFont,
			position: Vec2{
				x: instructionsX,
				y: 300,
			},
		},
	}
}

func (ms *MainGameScreen) Draw(screen *ebiten.Image) {
	text.Draw(screen, ms.titleText.textValue, ms.titleText.fontFace, ms.titleText.position.x, ms.titleText.position.y, color.White)
	text.Draw(screen, ms.instructionsText.textValue, ms.instructionsText.fontFace, ms.instructionsText.position.x, ms.instructionsText.position.y, color.White)
}

var mainGameScreen = NewMainGameScreen()
