package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputBoard struct {
	w, h int
	base *Rect
}

func NewInputBoard(w, h int, bgc color.Color) *InputBoard {
	base := NewRect(w, h, bgc)
	base.Fill()
	return &InputBoard{w, h, base}
}

func (ib *InputBoard) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(ib.base.Image(), op)
}
