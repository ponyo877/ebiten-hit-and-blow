package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type HistoryBoard struct {
	myHistory *History
	emHistory *History
	w, h      int
	bgColor   color.Color
	base      *Rect
}

func NewHistoryBoard(mh, eh *History, w, h int, bgc color.Color) *HistoryBoard {
	rect := NewRect(w, h, bgc)
	return &HistoryBoard{mh, eh, w, h, bgc, rect}
}

func (hb *HistoryBoard) Draw(screen *ebiten.Image, x, y int) {
	hb.base.Fill()
	hb.myHistory.Draw(hb.base.Image(), 0+5, 5)
	hb.emHistory.Draw(hb.base.Image(), hb.w/2+5, 5)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(hb.base.Image(), op)
}
