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
}

func NewNumberCard(n, w, h, t int, bgc, txtc color.Color) *Card {
	return NewCard(fmt.Sprint(n), w, h, t, bgc, txtc)
}

func NewCard(n string, w, h, t int, bgc, txtc color.Color) *Card {
	return &Card{n, w, h, t, bgc, txtc}
}

func (c *Card) Draw(screen *ebiten.Image, x, y int) {
	rect := NewRect(x, y, c.w, c.h, c.bgColor)
	NewText(c.text, mplusNormalFont(c.txtSize), c.txtColor).Draw(rect.Image(), 0, 0)
	rect.Draw(screen)
}
