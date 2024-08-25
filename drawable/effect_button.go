package drawable

import (
	"image/color"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type EffectButton struct {
	card       *Card
	x, y       int
	ebgc, dbgc color.Color
	enable     bool
	inputField *Input
}

func NewEffectButton(n string, w, h, ts int, ebgc, dbgc, txtc color.Color, x, y int, inputField *Input) *EffectButton {
	return &EffectButton{NewCard(n, w, h, ts, ebgc, txtc), x, y, ebgc, dbgc, true, inputField}
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

func (b *EffectButton) Disable() {
	b.enable = false
	b.card.SetColor(b.dbgc)
}

func (b *EffectButton) Enable() {
	b.enable = true
	b.card.SetColor(b.ebgc)
}

func (b *EffectButton) In(x, y int) bool {
	w, h := b.Bounds()
	return x >= b.x && x <= b.x+w && y >= b.y && y <= b.y+h
}

func (b *EffectButton) Draw(screen *ebiten.Image) {
	b.card.Draw(screen, b.x, b.y)
}

func (b *EffectButton) DrawCenter(screen *ebiten.Image) {
	w, _ := b.Bounds()
	b.card.Draw(screen, screen.Bounds().Dx()/2-w/2, b.y)
}
