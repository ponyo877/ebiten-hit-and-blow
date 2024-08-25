package drawable

import (
	"bytes"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/ponyo877/ebiten-hit-and-blow/static"
)

type Message struct {
	text       string
	numberFlag bool
	size       int
	x, y       int
	light      color.Color
	dark       color.Color
	counter    int
}

func NewMessage(t string, s, x, y int, light, dark color.Color) *Message {
	return &Message{t, false, s, x, y, light, dark, 0}
}

func (t *Message) Bounds() (int, int) {
	w, h := text.Measure(t.text, t.font(), 0)
	return int(w), int(h)
}

func (t *Message) Message() string {
	return t.text
}

func (t *Message) SetMessage(text string) {
	t.text = text
}

func (t *Message) font() *text.GoTextFace {
	var source *text.GoTextFaceSource
	source, _ = text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if _, err := strconv.Atoi(t.text); err == nil {
		source, _ = text.NewGoTextFaceSource(bytes.NewReader(static.NumberFont))
	}
	return &text.GoTextFace{
		Source: source,
		Size:   float64(t.size),
	}
}

func (t *Message) getFrameOp() color.Color {
	t.counter++
	maxCount := 20
	lightRatio := 0.7
	changeCount := int(float64(maxCount) * lightRatio)
	switch {
	case t.counter <= changeCount:
		return t.light
	case changeCount < t.counter && t.counter <= maxCount:
		return t.dark
	case maxCount < t.counter:
		t.counter = 0
		return t.light
	default:
		return t.dark
	}
}

func (t *Message) Draw(screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.ColorScale.ScaleWithColor(t.light)
	textOp.LineSpacing = float64(t.size)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	textOp.Filter = ebiten.FilterLinear
	textOp.GeoM.Translate(float64(t.x), float64(t.y))
	text.Draw(screen, t.text, t.font(), textOp)
}

func colorScale(clr color.Color) (rf, gf, bf, af float32) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float32(r) / float32(a)
	gf = float32(g) / float32(a)
	bf = float32(b) / float32(a)
	af = float32(a) / 0xffff
	return
}
