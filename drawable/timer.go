package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	time     int
	r        int
	txtSize  int
	bgColor  color.Color
	txtColor color.Color
}

func NewTimer(t, r, ts int, bgc, tc color.Color) *Timer {
	return &Timer{t, r, ts, bgc, tc}
}

func (t *Timer) Draw(screen *ebiten.Image, x, y int) {
	rect := NewRect(x, y, t.r, t.r, t.bgColor)
	NewText(fmt.Sprint(t.time), mplusNormalFont(t.txtSize), t.txtColor).Draw(rect.Image(), 0, 0)
	rect.Draw(screen)
}
