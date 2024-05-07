package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputBoard struct {
	w, h       int
	inputField *Input
	tenkey     *Tenkey
	txtColor   color.Color
	base       *Rect
	text       *Text
}

func NewInputBoard(w, h, ts int, m string, i *Input, te *Tenkey, bgc, txc color.Color) *InputBoard {
	base := NewRect(w, h, bgc)
	text := NewText(m, ts, txc)
	return &InputBoard{w, h, i, te, txc, base, text}
}

func (ib *InputBoard) SetMessage(m string) {
	ib.text.SetText(m)
}

func (ib *InputBoard) Draw(screen *ebiten.Image, x, y int) {
	ib.base.Fill()
	ib.text.Draw(ib.base.Image(), 0, -ib.h/2+ib.h/20)
	iw, _ := ib.inputField.Bounds()
	ib.inputField.Draw(ib.base.Image(), ib.w/2-iw/2, ib.h*3/20)
	tw, _ := ib.tenkey.Bounds()
	ib.tenkey.Draw(ib.base.Image(), ib.w/2-tw/2, ib.h*19/40)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(ib.base.Image(), op)
}
