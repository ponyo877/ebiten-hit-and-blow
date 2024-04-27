package drawable

import "github.com/hajimehoshi/ebiten/v2"

type Hand struct {
	cards  []*Card
	margin int
}

func NewHand(c []*Card, m int) *Hand {
	return &Hand{c, m}
}

func (h *Hand) Draw(screen *ebiten.Image, x, y int) {
	for i, c := range h.cards {
		c.Draw(screen, x+i*(c.w+h.margin), y)
	}
}
