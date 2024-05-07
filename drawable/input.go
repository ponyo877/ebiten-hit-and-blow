package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	w, h            int
	txtSize, margin int
	Input           []*Card
}

func NewInput(texts []string, w, h, ts, m int, bgc, txc color.Color) *Input {
	cs := make([]*Card, len(texts))
	for i, n := range texts {
		cs[i] = NewCard(n, w, h, ts, bgc, txc)
	}
	return &Input{w, h, ts, m, cs}
}

func NewNumberInput(ns []int, w, h, ts, m int, bgc, txc color.Color) *Input {
	cs := make([]string, len(ns))
	for i, n := range ns {
		cs[i] = fmt.Sprint(n)
	}
	return NewInput(cs, w, h, ts, m, bgc, txc)
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
	fsize := h * 3 / 50
	NewRounded(w+4*fsize, h+4*fsize, color.White).Draw(screen, x-2*fsize, y-2*fsize)
	NewRounded(w+2*fsize, h+2*fsize, HistoryFrameColor).Draw(screen, x-fsize, y-fsize)
	for i, c := range cs.Input {
		c.Draw(screen, x+i*(cs.w+cs.margin), y)
	}
}
