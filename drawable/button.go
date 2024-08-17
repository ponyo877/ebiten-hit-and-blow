package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	card       *Card
	inputField *Input
}

func NewNumberButton(n, w, h, ts int, bgc, txtc color.Color, inputField *Input) *Button {
	return &Button{NewNumberCard(n, w, h, ts, bgc, txtc), inputField}
}

func NewButton(n string, w, h, ts int, bgc, txtc color.Color, inputField *Input) *Button {
	return &Button{NewCard(n, w, h, ts, bgc, txtc), inputField}
}

func (b *Button) Bounds() (int, int) {
	return b.card.Bounds()
}

func (b *Button) Push() {
	b.inputField.Add(b.card.Text())
}

func (b *Button) Send(do func([]int)) {
	do(b.inputField.Numbers())
	b.inputField.Clear()
}

func (b *Button) Clear() {
	b.inputField.Clear()
}

func (b *Button) In(x, y int) bool {
	w, h := b.Bounds()
	return x >= 0 && x <= w && y >= 0 && y <= h
}

func (b *Button) Draw(screen *ebiten.Image, x, y int) {
	b.card.Draw(screen, x, y)
}
