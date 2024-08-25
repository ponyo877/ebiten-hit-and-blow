package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type NumberButton struct {
	card             *Card
	ebgc, dbgc, txtc color.Color
	inputField       *Input
	enable           bool
}

func NewNumberButton(n, w, h, ts int, ebgc, dbgc, txtc color.Color, inputField *Input) *NumberButton {
	return &NumberButton{NewNumberCard(n, w, h, ts, ebgc, txtc), ebgc, dbgc, txtc, inputField, true}
}

func (b *NumberButton) Bounds() (int, int) {
	return b.card.Bounds()
}

func (b *NumberButton) Push() {
	if !b.enable {
		return
	}
	b.inputField.Add(b.card.Text())
}

func (b *NumberButton) Enable() {
	b.enable = true
	b.card.SetColor(b.ebgc)
}

func (b *NumberButton) Disable() {
	b.enable = false
	b.card.SetColor(b.dbgc)
}

func (b *NumberButton) Draw(screen *ebiten.Image, x, y int) {
	b.card.Draw(screen, x, y)
}
