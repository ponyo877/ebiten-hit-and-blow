package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Card struct {
	w, h    int
	txtSize int
	base    *Rect
	text    *Text
}

func NewCard(n string, w, h, ts int, bgc, txc color.Color) *Card {
	base := NewRect(w, h, bgc)
	text := NewText(n, ts, txc)
	return &Card{w, h, ts, base, text}
}

func NewNumberCard(n, w, h, t int, bgc, txtc color.Color) *Card {
	return NewCard(fmt.Sprint(n), w, h, t, bgc, txtc)
}

func (c *Card) Bounds() (int, int) {
	return c.w, c.h
}

func (c *Card) Draw(screen *ebiten.Image, x, y int) {
	c.base.Fill()
	c.text.Draw(c.base.Image(), 0, 0)
	c.base.Draw(screen, x, y)
}
