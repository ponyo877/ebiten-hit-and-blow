package drawable

import (
	"image/color"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type EffectButton struct {
	card       *Card
	x, y       int
	inputField *Input
}

func NewEffectButton(n string, w, h, ts int, bgc, txtc color.Color, x, y int, inputField *Input) *EffectButton {
	return &EffectButton{NewCard(n, w, h, ts, bgc, txtc), x, y, inputField}
}

func (b *EffectButton) Bounds() (int, int) {
	return b.card.Bounds()
}

func (b *EffectButton) Push() {
	b.inputField.Add(b.card.Text())
}

func (b *EffectButton) Send(do func([]int)) {
	log.Print("send1")
	do(b.inputField.Numbers())
	log.Print("send2")
	b.inputField.Clear()
}

func (b *EffectButton) Clear() {
	b.inputField.Clear()
}

func (b *EffectButton) In(x, y int) bool {
	w, h := b.Bounds()
	return x >= b.x && x <= b.x+w && y >= b.y && y <= b.y+h
}

func (b *EffectButton) Draw(screen *ebiten.Image) {
	b.card.Draw(screen, b.x, b.y)
}
