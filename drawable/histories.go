package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var maxCount = 9

type History struct {
	feedbacks    []*Feedback
	w, h         int
	base         *Rect
	headerCard   *Card
	hbHeaderCard *Card
	emptyRect    *Rect
	hbEmptyRect  *Rect
}

func NewHistory(fs []*Feedback, w, h int, hdTxt string, bgc, txc color.Color) *History {
	base := NewRect(w, h*(maxCount+1)+1, bgc)
	headerCard := NewCard(hdTxt, w*2/3-5, h-2, h*2/3, bgc, txc)
	hbHeaderCard := NewCard("H B", w/3-2, h-2, h*2/3, bgc, txc)
	emptyRect := NewRect(w*2/3-5, h-2, color.White)
	hbEmptyRect := NewRect(w/3-2, h-2, color.White)
	return &History{fs, w, h, base, headerCard, hbHeaderCard, emptyRect, hbEmptyRect}
}

func (h *History) AddFeedback(fb *Feedback) {
	h.feedbacks = append(h.feedbacks, fb)
}

func (h *History) Draw(screen *ebiten.Image, x, y int) {
	h.base.Fill()
	h.headerCard.Draw(h.base.Image(), 3, 0)
	h.hbHeaderCard.Draw(h.base.Image(), h.w*2/3, 0)
	for i := 0; i < maxCount; i++ {
		h.emptyRect.Fill()
		h.hbEmptyRect.Fill()
		if i < len(h.feedbacks) {
			haw, _ := h.feedbacks[i].Hand().Bounds()
			hiw, _ := h.feedbacks[i].Hint().Bounds()
			h.feedbacks[i].Hand().Draw(h.emptyRect.Image(), -haw/2+(h.w*2/3-5)/2, 2)
			h.feedbacks[i].Hint().Draw(h.hbEmptyRect.Image(), -hiw/2+(h.w/3-2)/2, 2)
		}
		h.emptyRect.Draw(h.base.Image(), 3, (i+1)*h.h)
		h.hbEmptyRect.Draw(h.base.Image(), h.w*2/3, (i+1)*h.h)
	}
	h.base.Draw(screen, x, y)
}
