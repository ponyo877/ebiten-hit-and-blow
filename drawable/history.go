package drawable

import "github.com/hajimehoshi/ebiten/v2"

type History struct {
	w, h      int
	feedbacks []*Feedback
}

func NewHistory(w, h int, fs []*Feedback) *History {
	return &History{w, h, fs}
}

func (h *History) Draw(screen *ebiten.Image, x, y int) {
	for i, f := range h.feedbacks {
		f.Draw(screen, x, y+i*f.h)
	}
}
