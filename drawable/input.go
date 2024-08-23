package drawable

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	maxLength = 3
)

type Input struct {
	w, h            int
	txtSize, margin int
	bgc, txc        color.Color
	Input           []*Card
}

func NewInput(texts []string, w, h, ts, m int, bgc, txc color.Color) *Input {
	cs := make([]*Card, len(texts))
	for i, n := range texts {
		cs[i] = NewCard(n, w, h, ts, bgc, txc)
	}
	return &Input{w, h, ts, m, bgc, txc, cs}
}

func NewNumberInput(ns []int, w, h, ts, m int, bgc, txc color.Color) *Input {
	cs := make([]string, len(ns))
	for i, n := range ns {
		cs[i] = fmt.Sprint(n)
	}
	return NewInput(cs, w, h, ts, m, bgc, txc)
}

func NewEmptyInput(w, h, ts, m int, bgc, txc color.Color) *Input {
	return NewInput([]string{}, w, h, ts, m, bgc, txc)
}

func (cs *Input) Bounds() (int, int) {
	cnt := maxLength
	return cs.w*cnt + cs.margin*(cnt-1), cs.h
}

func (cs *Input) Addble() bool {
	return len(cs.Input) < maxLength
}

func (cs *Input) Add(t string) {
	if !cs.Addble() {
		return
	}
	cs.Input = append(cs.Input, NewCard(t, cs.w, cs.h, cs.txtSize, HistoryFrameColor, color.Black))
}

func (cs *Input) Clear() {
	cs.Input = []*Card{}
}

func (cs *Input) Numbers() []int {
	ns := make([]int, len(cs.Input))
	for i, c := range cs.Input {
		num, _ := strconv.Atoi(c.Text())
		ns[i] = num
	}
	return ns
}

func (cs *Input) Draw(screen *ebiten.Image, x, y int) {
	w, h := cs.Bounds()
	fsize := h * 3 / 50
	NewRounded(w+4*fsize, h+4*fsize, color.White).Draw(screen, x-2*fsize, y-2*fsize)
	NewRounded(w+2*fsize, h+2*fsize, HistoryFrameColor).Draw(screen, x-fsize, y-fsize)
	// for i, c := range cs.Input {
	// 	c.Draw(screen, x+i*(cs.w+cs.margin), y)
	// }
	for i := 0; i < maxLength; i++ {
		if i < len(cs.Input) {
			cs.Input[i].Draw(screen, x+i*(cs.w+cs.margin), y)
			continue
		}
		NewCard("-", cs.w, cs.h, cs.txtSize, cs.bgc, cs.txc).Draw(screen, x+i*(cs.w+cs.margin), y)
	}
}
