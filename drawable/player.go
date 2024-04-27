package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	margin = 5
)

type Player struct {
	w, h    int
	txtSize int
	icon    *Icon
	name    string
	rate    int
	bgColor color.Color
}

func NewPlayer(w, h, t int, i *Icon, n string, r int, bgc color.Color) *Player {
	return &Player{w, h, t, i, n, r, bgc}
}

func (p *Player) Image() *ebiten.Image {
	img := ebiten.NewImage(p.w, p.h)
	img.Fill(p.bgColor)
	return img
}

func (p *Player) Draw(screen *ebiten.Image, x, y int) {
	img := p.Image()
	p.icon.Draw(img, 0, 0)
	iconW := p.icon.w
	txtFont := mplusNormalFont(p.txtSize)
	NewText(p.name, txtFont, color.Black).Draw(img, iconW, 0)
	NewText(fmt.Sprint(p.rate), txtFont, color.Black).Draw(img, iconW, p.txtSize+margin)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
}
