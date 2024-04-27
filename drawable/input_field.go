package drawable

import "github.com/hajimehoshi/ebiten/v2"

type InputField struct{}

func NewInputField() *InputField {
	return &InputField{}
}

func (i *InputField) Image() *ebiten.Image {
	return nil
}

func (i *InputField) Draw(screen *ebiten.Image) {
	inputField := i.Image()
	screen.DrawImage(inputField, nil)
}
