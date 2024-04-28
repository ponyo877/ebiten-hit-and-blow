package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	time  int
	r     int
	color color.Color
}

func NewTimer(t, r int, c color.Color) *Timer {
	return &Timer{t, r, c}
}

func (t *Timer) Image() *ebiten.Image {
	img := ebiten.NewImage(t.r, t.r)
	img.Fill(t.color)
	return nil
}

func (t *Timer) Draw(screen *ebiten.Image, x, y int) {
	img := t.Image()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, nil)
}
