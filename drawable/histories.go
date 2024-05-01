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
	base := NewRect(w, h, bgc)
	text := NewText(hdTxt, h, txc) // .Draw(img, 0, 0)
	fw, fh := fs[0].Bounds()
	emptyFB := NewRect(fw, fh, bgc)
	return &History{fs, w, h, base, text, emptyFB}
}

func (h *History) Draw(screen *ebiten.Image, x, y int) {
	var i int
	h.base.Fill()
	h.text.Draw(h.base.Image(), 0, 0)
	h.base.Draw(screen, x, y)
	for i, f := range h.feedbacks {
		f.Draw(screen, x, y+(i+1)*h.h)
	}
	for i = len(h.feedbacks); i < maxCount; i++ {
		h.emptyFB.Draw(screen, x, y+i*h.h)
	}
}
