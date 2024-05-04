package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var maxCount = 10

type History struct {
	feedbacks []*Feedback
	w, h      int
	base      *Rect
	text      *Text
	text2     *Text
	emptyFB   *Rect
}

func NewHistory(fs []*Feedback, w, h int, hdTxt string, bgc, txc color.Color) *History {
	base := NewRect(w, h*(maxCount+1)+2, bgc)
	text := NewText(hdTxt, h*2/3, txc)
	text2 := NewText("H B", h*2/3, txc)
	fw, fh := fs[0].Bounds()
	emptyFB := NewRect(fw, fh, bgc)
	return &History{fs, w, h, base, text, text2, emptyFB}
}

func (h *History) Draw(screen *ebiten.Image, x, y int) {
	var i int
	h.base.Fill()
	// h.text.Draw(h.base.Image(), 0, 0)
	image := ebiten.NewImage(h.w*2/3-2, h.h-2)
	image.Fill(HistoryFrameColor)
	h.text.Draw(image, 0, 0)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	h.base.Image().DrawImage(image, op)

	image2 := ebiten.NewImage(h.w/3-2, h.h-2)
	image2.Fill(HistoryFrameColor)
	h.text2.Draw(image2, 0, 0)
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(h.w*2/3), 0)
	h.base.Image().DrawImage(image2, op2)
	for i, f := range h.feedbacks {
		f.Draw(h.base.Image(), 0, (i+1)*h.h)
	}
	for i = len(h.feedbacks); i < maxCount; i++ {
		image := ebiten.NewImage(h.w*2/3-5, h.h-2)
		image.Fill(color.White)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(3, float64((i+1)*h.h))
		h.base.Image().DrawImage(image, op)

		image2 := ebiten.NewImage(h.w/3-2, h.h-2)
		image2.Fill(color.White)
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(h.w*2/3), float64((i+1)*h.h))
		h.base.Image().DrawImage(image2, op2)

		h.emptyFB.Draw(h.base.Image(), 0, i*h.h)
	}
	h.base.Draw(screen, x, y)

}
