package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	w, h    int
	txtSize int
	icon    *Icon
	name    *Text
	rate    *Text
	base    *Rect
}

func NewPlayer(w, h, t int, i *Icon, n string, r string, bgc, txc color.Color) *Player {
	base := NewRect(w, h, bgc)
	name := NewText(n, t, txc)
	rate := NewText(r, t, txc)
	return &Player{w, h, t, i, name, rate, base}
}

func (p *Player) SetName(n string) {
	p.name.SetText(n)
}

func (p *Player) SetRate(r int) {
	p.rate.SetText(fmt.Sprint(r))
}

func (p *Player) Draw(screen *ebiten.Image, x, y int) {
	p.icon.Draw(p.base.Image(), 0, 0)
	iw, _ := p.icon.Bounds()
	margin := p.h / 9

	p.name.Draw(p.base.Image(), iw, -p.txtSize+margin)
	p.rate.Draw(p.base.Image(), iw, 2*margin)
	p.base.Draw(screen, x, y)
}
