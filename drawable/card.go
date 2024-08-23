package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Card struct {
	w, h    int
	base    *Rect
	rounded *Rounded
	text    *Text
}

func NewCard(n string, w, h, ts int, bgc, txc color.Color) *Card {
	base := NewRect(w, h, bgc)
	rounded := NewRounded(w, h, bgc)
	text := NewText(n, ts, txc)
	return &Card{w, h, base, rounded, text}
}

func NewNumberCard(n, w, h, t int, bgc, txtc color.Color) *Card {
	return NewCard(fmt.Sprint(n), w, h, t, bgc, txtc)
}

func (c *Card) Bounds() (int, int) {
	return c.w, c.h
}

func (c *Card) Text() string {
	return c.text.Text()
}

func (c *Card) SetColor(bgc color.Color) {
	c.rounded = NewRounded(c.w, c.h, bgc)
}

func (c *Card) Draw(screen *ebiten.Image, x, y int) {
	c.rounded.Draw(c.base.Image(), 0, 0)
	c.text.Draw(c.base.Image(), 0, 0)
	c.base.Draw(screen, x, y)
}
