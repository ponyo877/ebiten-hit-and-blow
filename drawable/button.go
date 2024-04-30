package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	card *Card
}

func NewNumberButton(n, w, h, ts int, bgc, txtc color.Color) *Button {
	return &Button{NewNumberCard(n, w, h, ts, bgc, txtc)}
}

func NewButton(t string, w, h, ts int, bgc, txtc color.Color) *Button {
	return &Button{NewCard(t, w, h, ts, bgc, txtc)}
}

func (b *Button) Bounds() (int, int) {
	return b.card.Bounds()
}

func (b *Button) Draw(screen *ebiten.Image, x, y int) {
	b.card.Draw(screen, x, y)
}
