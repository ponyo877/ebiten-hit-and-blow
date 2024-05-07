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
	bgColor color.Color
	base    *Rect
	text    *Text
}

func NewTimer(t, r, ts int, bgc, tc color.Color) *Timer {
	base := NewRect(2*r, 2*r, bgc)
	text := NewText(fmt.Sprint(t), ts, tc)
	return &Timer{t, r, bgc, base, text}
}

func (t *Timer) Draw(screen *ebiten.Image, x, y int) {
	vector.DrawFilledCircle(screen, float32(x+t.r), float32(y+t.r), float32(t.r), t.bgColor, true)
	t.text.Draw(t.base.Image(), 0, 0)
	t.base.Draw(screen, x, y)
}
