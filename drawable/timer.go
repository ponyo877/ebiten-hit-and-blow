package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	rect := NewRect(2*t.r, 2*t.r, color.NRGBA{0, 0, 0, 0})
	vector.DrawFilledCircle(rect.Image(), float32(x+t.r), float32(y), float32(t.r), t.bgColor, false)
	NewText(fmt.Sprint(t.time), t.txtSize, t.txtColor).Draw(rect.Image(), 0, 0)
	rect.Draw(screen, x, y)
}
