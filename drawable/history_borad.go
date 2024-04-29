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
}

func NewHistoryBoard(mh, eh *History, w, h int, bgc color.Color) *HistoryBoard {
	return &HistoryBoard{mh, eh, w, h, bgc}
}

func (hb *HistoryBoard) Image() *ebiten.Image {
	img := ebiten.NewImage(hb.w, hb.h)
	img.Fill(hb.bgColor)
	return img
}

func (hb *HistoryBoard) Draw(screen *ebiten.Image, x, y int) {
	img := hb.Image()
	hb.myHistory.Draw(img, 0, 0)

	cx := img.Bounds().Dx() / 2
	hb.emHistory.Draw(img, cx, 0)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
}
