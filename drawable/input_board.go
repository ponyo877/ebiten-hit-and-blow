package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputBoard struct {
	w, h       int
	message    string
	timer      *Timer
	inputField *Cards
	tenkey     *Tenkey
	color      color.Color
}

func NewInputBoard(w, h int, m string, ti *Timer, i *Cards, te *Tenkey, c color.Color) *InputBoard {
	return &InputBoard{w, h, m, ti, i, te, c}
}

func (ib *InputBoard) Image() *ebiten.Image {
	img := ebiten.NewImage(ib.w, ib.h)
	img.Fill(ib.color)
	return img
}

func (ib *InputBoard) SetMessage(m string) {
	ib.message = m
}

func (ib *InputBoard) Draw(screen *ebiten.Image, x, y int) {
	img := ib.Image()
	img.Fill(ib.color)
	NewText(ib.message, 10, color.Black).Draw(img, 0, 0)
	ib.timer.Draw(img, 0, 10)
	iw, _ := ib.inputField.Bounds()
	ib.inputField.Draw(img, img.Bounds().Dx()/2-iw/2, 40)
	tw, _ := ib.tenkey.Bounds()
	ib.tenkey.Draw(img, img.Bounds().Dx()/2-tw/2, 100)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
}
