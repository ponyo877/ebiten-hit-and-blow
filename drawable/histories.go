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
	emptyFB   *Rect
}

func NewHistory(fs []*Feedback, w, h int, hdTxt string, bgc, txc color.Color) *History {
	base := NewRect(w, h*(maxCount+1), bgc)
	text := NewText(hdTxt, h, txc) // .Draw(img, 0, 0)
	fw, fh := fs[0].Bounds()
	emptyFB := NewRect(fw, fh, bgc)
	return &History{fs, w, h, base, text, emptyFB}
}

func (h *History) Draw(screen *ebiten.Image, x, y int) {
	var i int
	h.base.Fill()
	// h.text.Draw(h.base.Image(), 0, 0)
	for i, f := range h.feedbacks {
		image := ebiten.NewImage(h.w-2, h.h-2)
		image.Fill(color.White)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64((i+1)*h.h))
		h.base.Image().DrawImage(image, op)
		f.Draw(h.base.Image(), 0, (i+1)*h.h)
	}
	for i = len(h.feedbacks); i < maxCount; i++ {
		image := ebiten.NewImage(h.w-2, h.h-2)
		image.Fill(color.White)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64((i+1)*h.h))
		h.base.Image().DrawImage(image, op)
		h.emptyFB.Draw(h.base.Image(), 0, i*h.h)
	}
	h.base.Draw(screen, x, y)

}
