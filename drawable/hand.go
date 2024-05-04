package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hand struct {
	w, h            int
	txtSize, margin int
	Hand            []*Card
	bg              []*Card
}

func NewNumberHand(ns []int, w, h, ts, m int, bgc, txc color.Color) *Hand {
	cs := make([]*Card, len(ns))
	bg := make([]*Card, len(ns))
	for i, n := range ns {
		cs[i] = NewNumberCard(n, w, h, ts, bgc, txc)
		bg[i] = NewCard("", w+6, h+6, ts, HistoryBackgroundColor, txc)
	}
	return &Hand{w, h, ts, m, cs, bg}
}

func (cs *Hand) Bounds() (int, int) {
	cnt := len(cs.Hand)
	return (cs.w)*cnt + cs.margin*(cnt-1), cs.h
}

func (cs *Hand) Draw(screen *ebiten.Image, x, y int) {
	for i, c := range cs.Hand {
		cs.bg[i].Draw(screen, x+i*(cs.w+cs.margin)-3, y-3)
		c.Draw(screen, x+i*(cs.w+cs.margin), y)
	}
}
