package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	text     string
	w, h     int
	txtSize  int
	bgColor  color.Color
	txtColor color.Color
}

func NewNumberButton(n, w, h, ts int, bgc, txtc color.Color) *Button {
	return NewButton(fmt.Sprint(n), w, h, ts, bgc, txtc)
}

func NewButton(t string, w, h, ts int, bgc, txtc color.Color) *Button {
	return &Button{t, w, h, ts, bgc, txtc}
}

func (b *Button) Draw(screen *ebiten.Image, x, y int) {
	rect := NewRect(x, y, b.w, b.h, b.bgColor)
	NewText(b.text, mplusNormalFont(b.txtSize), b.txtColor).Draw(rect.Image(), 0, 0)
	rect.Draw(screen)
}
