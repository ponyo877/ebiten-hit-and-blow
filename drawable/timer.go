package drawable

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Timer struct {
	r       int
	bgColor color.Color
	base    *Rect
	text    *Text
}

func NewTimer(t, r, ts int, bgc, tc color.Color) *Timer {
	base := NewRect(2*r, 2*r, bgc)
	text := NewText(fmt.Sprint(t), ts, tc)
	return &Timer{r, bgc, base, text}
}

func (t *Timer) Set(second int) {
	t.text.SetText(strconv.Itoa(second))
}

func (t *Timer) Draw(screen *ebiten.Image, x, y int) {
	vector.DrawFilledCircle(t.base.Image(), float32(t.r), float32(t.r), float32(t.r), t.bgColor, true)
	t.text.Draw(t.base.Image(), 0, 0)
	t.base.Draw(screen, x, y)
}
