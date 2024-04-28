package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Feedback struct {
	estimate *Estimate
	hint     *Hint
}

func NewFeedback(e *Estimate, hi *Hint) *Feedback {
	return &Feedback{e, hi}
}

func (f *Feedback) Bounds() (int, int) {
	ew, eh := f.estimate.Bounds()
	hw, _ := f.hint.Bounds()
	return ew + hw, eh
}

func (f *Feedback) Draw(screen *ebiten.Image, x, y int) {
	f.estimate.Draw(screen, x, y)
	ew, _ := f.estimate.Bounds()
	f.hint.Draw(screen, x+ew, y)
}
