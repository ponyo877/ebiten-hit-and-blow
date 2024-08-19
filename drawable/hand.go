package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ponyo877/ebiten-hit-and-blow/entity"
)

type Hand struct {
	w, h            int
	txtSize, margin int
	Hand            []*Card
	bgc, txc        color.Color
	frameColor      color.Color
}

func NewHand(ns []string, w, h, ts, m int, bgc, txc, frc color.Color) *Hand {
	cs := make([]*Card, len(ns))
	for i, n := range ns {
		cs[i] = NewCard(n, w, h, ts, bgc, txc)
	}
	return &Hand{w, h, ts, m, cs, bgc, txc, frc}
}

func NewNumberHand(ns []int, w, h, ts, m int, bgc, txc, frc color.Color) *Hand {
	cs := make([]*Card, len(ns))
	for i, n := range ns {
		cs[i] = NewNumberCard(n, w, h, ts, bgc, txc)
	}
	return &Hand{w, h, ts, m, cs, bgc, txc, frc}
}

func (cs *Hand) Bounds() (int, int) {
	cnt := len(cs.Hand)
	return (cs.w)*cnt + cs.margin*(cnt-1), cs.h
}

func (cs *Hand) SetHand(hand *entity.Hand) {
	handInt := []int(*hand)
	cs.Hand = make([]*Card, len(handInt))
	for i, c := range handInt {
		cs.Hand[i] = NewNumberCard(c, cs.w, cs.h, cs.txtSize, cs.bgc, cs.txc)
	}
}

func (cs *Hand) Draw(screen *ebiten.Image, x, y int) {
	fsize := cs.h * 3 / 50
	for i, c := range cs.Hand {
		NewRounded(cs.w+2*fsize, cs.h+2*fsize, cs.frameColor).Draw(screen, x+i*(cs.w+cs.margin)-fsize, y-fsize)
		c.Draw(screen, x+i*(cs.w+cs.margin), y)
	}
}
