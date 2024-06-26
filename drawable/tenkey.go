package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Tenkey struct {
	buttons []*Button
	wmargin int
	hmargin int
}

func NewTenkey(bs []*Button, wm, hm int) *Tenkey {
	return &Tenkey{bs, wm, hm}
}

func (t *Tenkey) Bounds() (int, int) {
	w, h := t.buttons[0].Bounds()
	cnt := len(t.buttons) / 2
	return w*cnt + t.wmargin*(cnt-1), 2*h + t.hmargin
}

func (t *Tenkey) Draw(screen *ebiten.Image, x, y int) {
	turn := len(t.buttons) / 2
	for i, b := range t.buttons {
		w, h := b.Bounds()
		if i < turn {
			b.Draw(screen, x+i*(w+t.wmargin), y)
			continue
		}
		b.Draw(screen, x+(i-turn)*(w+t.wmargin), y+t.hmargin+h)
	}
}
