package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Timer struct {
	time    int
	r       int
	txtSize int
	bgColor color.Color
	base    *Rect
	text    *Text
}

func NewTimer(t, r, ts int, bgc, tc color.Color) *Timer {
	base := NewRect(2*r, 2*r, color.NRGBA{0, 0, 0, 0})
	text := NewText(fmt.Sprint(t), ts, tc)
	return &Timer{t, r, ts, bgc, base, text}
}

func (t *Timer) Draw(screen *ebiten.Image, x, y int) {
	vector.DrawFilledCircle(t.base.Image(), float32(x+t.r), float32(y), float32(t.r), t.bgColor, false)
	t.text.Draw(t.base.Image(), 0, 0)
	t.base.Draw(screen, x, y)
}
