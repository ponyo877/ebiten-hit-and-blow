package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputBoard struct {
	w, h       int
	inputField *Cards
	tenkey     *Tenkey
	txtColor   color.Color
	base       *Rect
	text       *Text
}

func NewInputBoard(w, h int, m string, i *Cards, te *Tenkey, bgc, txc color.Color) *InputBoard {
	base := NewRect(w, h, bgc)
	text := NewText(m, 10, txc)
	return &InputBoard{w, h, i, te, txc, base, text}
}

func (ib *InputBoard) SetMessage(m string) {
	ib.text = NewText(m, 10, ib.txtColor)
}

func (ib *InputBoard) Draw(screen *ebiten.Image, x, y int) {
	ib.base.Fill()
	ib.text.Draw(ib.base.Image(), 0, -ib.h/2+10)
	iw, _ := ib.inputField.Bounds()
	ib.inputField.Draw(ib.base.Image(), ib.w/2-iw/2, 40)
	tw, _ := ib.tenkey.Bounds()
	ib.tenkey.Draw(ib.base.Image(), ib.w/2-tw/2, 100)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(ib.base.Image(), op)
}
