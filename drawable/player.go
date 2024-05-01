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
	name    *Text
	rate    *Text
	base    *Rect
}

func NewPlayer(w, h, t int, i *Icon, n string, r int, bgc, txc color.Color) *Player {
	base := NewRect(w, h, bgc)
	name := NewText(n, t, txc)
	rate := NewText(fmt.Sprint(r), t, txc)
	return &Player{w, h, t, i, name, rate, base}
}

func (p *Player) Draw(screen *ebiten.Image, x, y int) {
	p.icon.Draw(p.base.Image(), 0, 0)
	iw, _ := p.icon.Bounds()
	p.name.Draw(p.base.Image(), iw, 0)
	p.rate.Draw(p.base.Image(), iw, p.txtSize+margin)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(p.base.Image(), op)
}
