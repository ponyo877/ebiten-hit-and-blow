package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputBoard struct {
	w, h       int
	timer      *Timer
	inputField *Cards
	tenkey     *Tenkey
	base       *Rect
	text       *Text
}

func NewInputBoard(w, h int, m string, ti *Timer, i *Cards, te *Tenkey, c color.Color) *InputBoard {
	base := NewRect(w, h, c)
	text := NewText(m, 10, color.Black)
	return &InputBoard{w, h, ti, i, te, base, text}
}

func (ib *InputBoard) SetMessage(m string) {
	ib.text = NewText(m, 10, color.Black)
}

func (ib *InputBoard) Draw(screen *ebiten.Image, x, y int) {
	ib.base.Fill()
	ib.text.Draw(ib.base.Image(), 0, 0)
	ib.timer.Draw(ib.base.Image(), 0, 10)
	iw, _ := ib.inputField.Bounds()
	ib.inputField.Draw(ib.base.Image(), ib.w/2-iw/2, 40)
	tw, _ := ib.tenkey.Bounds()
	ib.tenkey.Draw(ib.base.Image(), ib.w/2-tw/2, 100)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(ib.base.Image(), op)
}
