package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Tenkey struct {
	buttons    []*Button
	wmargin    int
	hmargin    int
	wholeWidth int
	y          int
}

func NewTenkey(bs []*Button, wm, hm int, w, y int) *Tenkey {
	return &Tenkey{bs, wm, hm, w, y}
}

func (t *Tenkey) Bounds() (int, int) {
	w, h := t.buttons[0].Bounds()
	cnt := len(t.buttons) / 2
	return w*cnt + t.wmargin*(cnt-1), 2*h + t.hmargin
}

func (t *Tenkey) WhichButtonByPosition(x, y int) *Button {
	turn := len(t.buttons) / 2
	for i, b := range t.buttons {
		w, h := b.Bounds()
		if i < turn {
			if x >= t.x()+i*(w+t.wmargin) && x <= t.x()+i*(w+t.wmargin)+w && y >= t.y && y <= t.y+h {
				return b
			}
			continue
		}
		if x >= t.x()+(i-turn)*(w+t.wmargin) && x <= t.x()+(i-turn)*(w+t.wmargin)+w && y >= t.y+t.hmargin+h && y <= t.y+t.hmargin+2*h {
			return b
		}
	}
	return nil
}

func (t *Tenkey) x() int {
	w, _ := t.Bounds()
	return t.wholeWidth/2 - w/2
}

func (t *Tenkey) Draw(screen *ebiten.Image) {
	turn := len(t.buttons) / 2
	for i, b := range t.buttons {
		w, h := b.Bounds()
		if i < turn {
			b.Draw(screen, t.x()+i*(w+t.wmargin), t.y)
			continue
		}
		b.Draw(screen, t.x()+(i-turn)*(w+t.wmargin), t.y+t.hmargin+h)
	}
}
