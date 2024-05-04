package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cards struct {
	w, h            int
	txtSize, margin int
	cards           []*Card
}

func NewNumberCards(ns []int, w, h, ts, m int, bgc, txc color.Color) *Cards {
	cs := make([]*Card, len(ns))
	for i, n := range ns {
		cs[i] = NewNumberCard(n, w, h, ts, bgc, txc)
	}
	return &Cards{w, h, ts, m, cs}
}

func NewCards(texts []string, w, h, ts, m int, bgc, txc color.Color) *Cards {
	cs := make([]*Card, len(texts))
	for i, n := range texts {
		cs[i] = NewCard(n, w, h, ts, bgc, txc)
	}
	return &Cards{w, h, ts, m, cs}
}

func NewEmptyCards(w, h, ts, m int, bgc, txc color.Color) *Cards {
	return NewCards([]string{"", "", ""}, w, h, ts, m, bgc, txc)
}

func (cs *Cards) Bounds() (int, int) {
	cnt := len(cs.cards)
	return cs.w*cnt + cs.margin*(cnt-1), cs.h
}

func (cs *Cards) Draw(screen *ebiten.Image, x, y int) {
	for i, c := range cs.cards {
		c.Draw(screen, x+i*(cs.w+cs.margin), y)
	}
}
