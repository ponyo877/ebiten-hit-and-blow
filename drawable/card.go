package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Card struct {
	text     string
	w, h     int
	txtSize  int
	bgColor  color.Color
	txtColor color.Color
	r        *Rect
	t        *Text
}

func NewCard(n string, w, h, ts int, bgc, txc color.Color) *Card {
	r := NewRect(w, h, bgc)
	t := NewText(n, ts, txc)
	return &Card{n, w, h, ts, bgc, txc, r, t}
}

func NewNumberCard(n, w, h, t int, bgc, txtc color.Color) *Card {
	return NewCard(fmt.Sprint(n), w, h, t, bgc, txtc)
}

func (c *Card) Bounds() (int, int) {
	return c.w, c.h
}

func (c *Card) Draw(screen *ebiten.Image, x, y int) {
	c.t.Draw(c.r.Image(), 0, 0)
	c.r.Draw(screen, x, y)
}
