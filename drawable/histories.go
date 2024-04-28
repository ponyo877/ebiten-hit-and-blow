package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type History struct {
	w, h        int
	header      string
	headerColor color.Color
	feedbacks   []*Feedback
}

func NewHistory(w, h int, hd string, hdColor color.Color, fs []*Feedback) *History {
	return &History{w, h, hd, hdColor, fs}
}

func (h *History) Draw(screen *ebiten.Image, x, y int) {
	img := ebiten.NewImage(h.w, h.h)
	img.Fill(h.headerColor)
	NewText(h.header, mplusNormalFont(h.h), color.Black).Draw(img, 0, 0)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
	for i, f := range h.feedbacks {
		f.Draw(screen, x, y+(i+1)*h.h)
	}
}
