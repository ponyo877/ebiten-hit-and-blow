package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Tenkey struct {
	buttons    []*NumberButton
	wmargin    int
	hmargin    int
	wholeWidth int
	y          int
}

func NewTenkey(bs []*NumberButton, wm, hm int, w, y int) *Tenkey {
	return &Tenkey{bs, wm, hm, w, y}
}

func (t *Tenkey) Bounds() (int, int) {
	w, h := t.buttons[0].Bounds()
	cnt := len(t.buttons) / 2
	return w*cnt + t.wmargin*(cnt-1), 2*h + t.hmargin
}

func (t *Tenkey) WhichButtonByPosition(x, y int) *NumberButton {
	turn := len(t.buttons) / 2
	w, h := t.buttons[0].Bounds()
	tx := t.x()
	for i, b := range t.buttons {
		if i < turn {
			if x >= tx+i*(w+t.wmargin) && x <= tx+i*(w+t.wmargin)+w && y >= t.y && y <= t.y+h {
				return b
			}
			continue
		}
		if x >= tx+(i-turn)*(w+t.wmargin) && x <= tx+(i-turn)*(w+t.wmargin)+w && y >= t.y+t.hmargin+h && y <= t.y+t.hmargin+2*h {
			return b
		}
	}
	return nil
}

func (t *Tenkey) x() int {
	w, _ := t.Bounds()
	return t.wholeWidth/2 - w/2
}

func (t *Tenkey) DrawPart(screen *ebiten.Image, num int) {
	turn := len(t.buttons) / 2
	if num < 0 {
		return
	}
	w, h := t.buttons[num].Bounds()
	if num < turn {
		t.buttons[num].Draw(screen, t.x()+num*(w+t.wmargin), t.y)
		return
	}
	t.buttons[num].Draw(screen, t.x()+(num-turn)*(w+t.wmargin), t.y+t.hmargin+h)
}

func (t *Tenkey) Draw(screen *ebiten.Image) {
	turn := len(t.buttons) / 2
	w, h := t.buttons[0].Bounds()
	tx := t.x()
	for i, b := range t.buttons {
		if i < turn {
			b.Draw(screen, tx+i*(w+t.wmargin), t.y)
			continue
		}
		b.Draw(screen, tx+(i-turn)*(w+t.wmargin), t.y+t.hmargin+h)
	}
}
