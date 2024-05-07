package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hand struct {
	w, h            int
	txtSize, margin int
	Hand            []*Card
	frameColor      color.Color
}

func NewNumberHand(ns []int, w, h, ts, m int, bgc, txc, frc color.Color) *Hand {
	cs := make([]*Card, len(ns))
	for i, n := range ns {
		cs[i] = NewNumberCard(n, w, h, ts, bgc, txc)
	}
	return &Hand{w, h, ts, m, cs, frc}
}

func (cs *Hand) Bounds() (int, int) {
	cnt := len(cs.Hand)
	return (cs.w)*cnt + cs.margin*(cnt-1), cs.h
}

func (cs *Hand) Draw(screen *ebiten.Image, x, y int) {
	fsize := cs.h * 3 / 50
	for i, c := range cs.Hand {
		NewRounded(cs.w+2*fsize, cs.h+2*fsize, cs.frameColor).Draw(screen, x+i*(cs.w+cs.margin)-fsize, y-fsize)
		c.Draw(screen, x+i*(cs.w+cs.margin), y)
	}
}
