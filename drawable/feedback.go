package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Feedback struct {
	w, h     int
	estimate *Estimate
	hint     *Hint
}

func NewFeedback(w, h int, e *Estimate, hi *Hint) *Feedback {
	return &Feedback{w, h, e, hi}
}

func (f *Feedback) Draw(screen *ebiten.Image, x, y int) {
	f.estimate.Draw(screen, x, y)
	f.hint.Draw(screen, x+f.estimate.w, y)
}
