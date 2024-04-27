package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Rect struct {
	x, y  int
	w, h  int
	color color.Color
	img   *ebiten.Image
}

func NewRect(x, y, w, h int, c color.Color) *Rect {
	rect := ebiten.NewImage(w, h)
	rect.Fill(c)
	return &Rect{x, y, w, h, c, rect}
}

func (r *Rect) Image() *ebiten.Image {
	return r.img
}

func (r *Rect) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.x), float64(r.y))
	screen.DrawImage(r.img, op)
}
