package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Tenkey struct {
	buttons []*Button
	margin  int
}

func NewTenkey(bs []*Button, m int) *Tenkey {
	return &Tenkey{bs, m}
}

func (t *Tenkey) Bounds() (int, int) {
	return t.buttons[0].w*len(t.buttons) + t.margin*(len(t.buttons)-1), t.buttons[0].h
}

func (t *Tenkey) Draw(screen *ebiten.Image, x, y int) {
	for i, b := range t.buttons {
		b.Draw(screen, x+i*(b.w+t.margin), y)
	}
}
