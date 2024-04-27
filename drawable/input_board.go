package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputBoard struct {
	message    string
	inputField *InputField
	tenkey     *Tenkey
	color      color.Color
}

func NewInputBoard(m string, i *InputField, t *Tenkey, c color.Color) *InputBoard {
	return &InputBoard{m, i, t, c}
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
	i.inputField.Draw(inputBoard)
	i.tenkey.Draw(inputBoard)
	screen.DrawImage(inputBoard, nil)
}
