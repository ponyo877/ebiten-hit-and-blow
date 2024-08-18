package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type NumberButton struct {
	card       *Card
	inputField *Input
}

func NewNumberButton(n, w, h, ts int, bgc, txtc color.Color, inputField *Input) *NumberButton {
	return &NumberButton{NewNumberCard(n, w, h, ts, bgc, txtc), inputField}
}

func (b *NumberButton) Bounds() (int, int) {
	return b.card.Bounds()
}

func (b *NumberButton) Push() {
	b.inputField.Add(b.card.Text())
}

func (b *NumberButton) Draw(screen *ebiten.Image, x, y int) {
	b.card.Draw(screen, x, y)
}
