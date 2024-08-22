package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	margin = 5
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

func (hb *HistoryBoard) MyHistory() *History {
	return hb.myHistory
}

func (hb *HistoryBoard) EmHistory() *History {
	return hb.emHistory
}

func (hb *HistoryBoard) Draw(screen *ebiten.Image, x, y int) {
	hb.base.Fill()
	hb.myHistory.Draw(hb.base.Image(), margin, margin)
	hb.emHistory.Draw(hb.base.Image(), hb.w/2+margin, margin)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(hb.base.Image(), op)
}
