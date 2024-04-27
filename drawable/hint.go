package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hint struct {
	w, h      int
	hit, blow int
	txtSize   int
	bgColor   color.Color
	txtColor  color.Color
}

func NewHint(w, h, hi, b, ts int, bgc, tc color.Color) *Hint {
	return &Hint{w, h, hi, b, ts, bgc, tc}
}

func (h *Hint) Image() *ebiten.Image {
	img := ebiten.NewImage(h.w, h.h)
	img.Fill(h.bgColor)
	return img
}

func (h *Hint) Draw(screen *ebiten.Image, x, y int) {
	font := mplusNormalFont(h.txtSize)
	hitImg := h.Image()
	NewText(fmt.Sprint(h.hit), font, h.txtColor).Draw(hitImg, 0, 0)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(hitImg, op)

	blowImg := h.Image()
	NewText(fmt.Sprint(h.blow), font, h.txtColor).Draw(blowImg, 0, 0)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x+h.w), float64(y))
	screen.DrawImage(blowImg, op)
}
