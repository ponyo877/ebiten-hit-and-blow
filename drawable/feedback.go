package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Feedback struct {
	hand *Cards
	hint *Cards
}

func NewFeedback(ha *Cards, hi *Cards) *Feedback {
	return &Feedback{ha, hi}
}

func NewEmptyFeedback(w, h, ts, m int, bgc, txc color.Color) *Feedback {
	return NewFeedback(NewEmptyCards(w, h, ts, m, bgc, txc), NewEmptyCards(w, h, ts, m, bgc, txc))
}

func (f *Feedback) Bounds() (int, int) {
	ew, eh := f.hand.Bounds()
	hw, _ := f.hint.Bounds()
	return ew + hw, eh
}

func (f *Feedback) Draw(screen *ebiten.Image, x, y int) {
	f.hand.Draw(screen, x, y)
	ew, _ := f.hand.Bounds()
	f.hint.Draw(screen, x+ew, y)
}
