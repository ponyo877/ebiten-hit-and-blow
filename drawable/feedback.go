package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Feedback struct {
	hand *Cards
	hint *Cards
	w, h int
	base *Rect
}

func NewFeedback(ha *Cards, hi *Cards, w, h int) *Feedback {
	base := NewRect(w, h, color.White)
	return &Feedback{ha, hi, w, h, base}
}

func NewEmptyFeedback(w, h, ts, m int, bgc, txc color.Color) *Feedback {
	return NewFeedback(NewEmptyCards(w, h, ts, m, bgc, txc), NewEmptyCards(w, h, ts, m, bgc, txc), w, h)
}

func (f *Feedback) Bounds() (int, int) {
	ew, eh := f.hand.Bounds()
	hw, _ := f.hint.Bounds()
	return ew + hw, eh
}

func (f *Feedback) Draw(screen *ebiten.Image, x, y int) {
	image := ebiten.NewImage(f.w*2/3-5, f.h-2)
	image.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x+3), float64(y))
	haw, _ := f.hand.Bounds()
	f.hand.Draw(image, -haw/2+image.Bounds().Dx()/2, 2)
	// image.DrawImage(f.base.Image(), op)
	screen.DrawImage(image, op)

	image2 := ebiten.NewImage(f.w/3-2, f.h-2)
	image2.Fill(color.White)
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(x+f.w*2/3), float64(y))
	hiw, _ := f.hint.Bounds()
	f.hint.Draw(image2, -hiw/2+image2.Bounds().Dx()/2, 2)
	// image2.DrawImage(f.base.Image(), op)
	screen.DrawImage(image2, op2)

	// haw, _ := f.hand.Bounds()
	// f.hand.Draw(screen, x+haw/2, y)
	// hiw, _ := f.hint.Bounds()
	// f.hint.Draw(screen, x+haw+hiw/2, y)
}
