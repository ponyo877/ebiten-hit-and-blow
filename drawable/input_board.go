package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputBoard struct {
	message    string
	timer      *Timer
	inputField *InputField
	tenkey     *Tenkey
	color      color.Color
}

func NewInputBoard(m string, ti *Timer, i *InputField, te *Tenkey, c color.Color) *InputBoard {
	return &InputBoard{m, ti, i, te, c}
}

func (i *InputBoard) Image() *ebiten.Image {
	return nil
}

func (i *InputBoard) SetMessage(m string) {
	i.message = m
}

func (i *InputBoard) Draw(screen *ebiten.Image) {
	inputBoard := i.Image()
	inputBoard.Fill(i.color)
	NewText(i.message, mplusNormalFont(10), color.Black).Draw(inputBoard, 0, 0)
	i.timer.Draw(inputBoard, 0, 10)
	i.inputField.Draw(inputBoard)
	i.tenkey.Draw(inputBoard)
	screen.DrawImage(inputBoard, nil)
}
