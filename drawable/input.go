package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	w, h            int
	txtSize, margin int
	Input           []*Card
}

func NewNumberInput(ns []int, w, h, ts, m int, bgc, txc color.Color) *Input {
	cs := make([]*Card, len(ns))
	for i, n := range ns {
		cs[i] = NewNumberCard(n, w, h, ts, bgc, txc)
	}
	return &Input{w, h, ts, m, cs}
}

func NewInput(texts []string, w, h, ts, m int, bgc, txc color.Color) *Input {
	cs := make([]*Card, len(texts))
	for i, n := range texts {
		cs[i] = NewCard(n, w, h, ts, bgc, txc)
	}
	return &Input{w, h, ts, m, cs}
}

func NewEmptyInput(w, h, ts, m int, bgc, txc color.Color) *Input {
	return NewInput([]string{"", "", ""}, w, h, ts, m, bgc, txc)
}

func (cs *Input) Bounds() (int, int) {
	cnt := len(cs.Input)
	return cs.w*cnt + cs.margin*(cnt-1), cs.h
}

func (cs *Input) Draw(screen *ebiten.Image, x, y int) {
	w, h := cs.Bounds()
	NewCard("", w+12, h+12, cs.txtSize, color.White, color.White).Draw(screen, x-6, y-6)
	NewCard("", w+6, h+6, cs.txtSize, HistoryFrameColor, color.White).Draw(screen, x-3, y-3)
	for i, c := range cs.Input {
		c.Draw(screen, x+i*(cs.w+cs.margin), y)
	}
}
