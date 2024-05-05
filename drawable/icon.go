package drawable

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Icon struct {
	w, h  int
	image image.Image
}

func NewIcon(w, h int, i image.Image) *Icon {
	return &Icon{w, h, i}
}

func (i *Icon) Bounds() (int, int) {
	return i.w, i.h
}

func (i *Icon) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(i.w)/float64(i.image.Bounds().Dx()), float64(i.h)/float64(i.image.Bounds().Dy()))
	op.GeoM.Translate(float64(x), float64(y))
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(ebiten.NewImageFromImage(i.image), op)
}
