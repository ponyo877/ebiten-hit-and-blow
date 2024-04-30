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
	w, h := t.buttons[0].Bounds()
	cnt := len(t.buttons)
	return w*cnt + t.margin*(cnt-1), h
}

func (t *Tenkey) Draw(screen *ebiten.Image, x, y int) {
	for i, b := range t.buttons {
		w, _ := b.Bounds()
		b.Draw(screen, x+i*(w+t.margin), y)
	}
}
