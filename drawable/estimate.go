package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Estimate struct {
	w, h     int
	numbers  []int
	txtSize  int
	bgColor  color.Color
	txtColor color.Color
}

func NewEstimate(w, h int, ns []int, ts int, bgc, txc color.Color) *Estimate {
	return &Estimate{w, h, ns, ts, bgc, txc}
}

func (e *Estimate) Image() *ebiten.Image {
	img := ebiten.NewImage(e.w, e.h)
	img.Fill(e.bgColor)
	return img
}

func (e *Estimate) Draw(screen *ebiten.Image, x, y int) {
	for i, n := range e.numbers {
		img := e.Image()
		txt := NewText(fmt.Sprint(n), mplusNormalFont(e.txtSize), e.txtColor)
		txt.Draw(img, 0, 0)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x+i*e.w), float64(y))
		screen.DrawImage(img, op)
	}
}
