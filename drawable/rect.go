package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Rect struct {
	w, h  int
	color color.Color
	image *ebiten.Image
}

func NewRect(w, h int, c color.Color) *Rect {
	rect := ebiten.NewImage(w, h)
	rect.Fill(c)
	return &Rect{w, h, c, rect}
}

func (r *Rect) Image() *ebiten.Image {
	return r.image
}

func (r *Rect) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(r.image, op)
}
