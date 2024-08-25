package drawable

import (
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
	cards           []*Card
	preCards        map[string]*Card
	numbers         []int
}

func NewInput(texts []string, w, h, ts, m int, bgc, txc color.Color) *Input {
	cs := make([]*Card, len(texts))
	for i, n := range texts {
		cs[i] = NewCard(n, w, h, ts, bgc, txc)
	}
	preCards := make(map[string]*Card)
	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i)
		preCards[num] = NewCard(num, w, h, ts, bgc, txc)
	}
	return &Input{w, h, ts, m, bgc, txc, cs, preCards, []int{}}
}

func NewEmptyInput(w, h, ts, m int, bgc, txc color.Color) *Input {
	return NewInput([]string{}, w, h, ts, m, bgc, txc)
}

func (cs *Input) Bounds() (int, int) {
	cnt := maxLength
	return cs.w*cnt + cs.margin*(cnt-1), cs.h
}

func (cs *Input) Addble() bool {
	return len(cs.cards) < maxLength
}

func (cs *Input) Add(t string) {
	if !cs.Addble() {
		return
	}
	num, _ := strconv.Atoi(t)
	cs.numbers = append(cs.numbers, num)
	cs.cards = append(cs.cards, cs.preCards[t])
}

func (cs *Input) Clear() {
	cs.numbers = []int{}
	cs.cards = []*Card{}
}

func (cs *Input) EndNumber() int {
	if len(cs.numbers) == 0 {
		return -1
	}
	return cs.numbers[len(cs.numbers)-1]
}

func (cs *Input) Numbers() []int {
	return cs.numbers
}

func (cs *Input) Draw(screen *ebiten.Image, x, y int) {
	w, h := cs.Bounds()
	harfScreenWidth := screen.Bounds().Dx() / 2
	fsize := h * 3 / 50
	NewRounded(w+4*fsize, h+4*fsize, color.White).Draw(screen, x-2*fsize-w/2+harfScreenWidth, y-2*fsize)
	NewRounded(w+2*fsize, h+2*fsize, HistoryFrameColor).Draw(screen, x-fsize-w/2+harfScreenWidth, y-fsize)
	for i := 0; i < maxLength; i++ {
		if i < len(cs.cards) {
			cs.cards[i].Draw(screen, x+i*(cs.w+cs.margin)-w/2+harfScreenWidth, y)
			continue
		}
		NewCard("-", cs.w, cs.h, cs.txtSize, cs.bgc, cs.txc).Draw(screen, x+i*(cs.w+cs.margin)-w/2+harfScreenWidth, y)
	}
}
