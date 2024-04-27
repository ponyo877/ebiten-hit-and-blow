package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Tenkey struct {
	buttons []*Button
	margin  int
}

func NewTenkey(bs []*Button, m int) *Tenkey {
	return &Tenkey{bs, m}
}

func (t *Tenkey) Image() *ebiten.Image {
	return nil
}

func (t *Tenkey) Draw(screen *ebiten.Image) {
	for i, b := range t.buttons {
		b.Draw(screen, i*(b.w+t.margin), 0)
	}
}
