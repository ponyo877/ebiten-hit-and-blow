package drawable

import (
	"bytes"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	text  string
	size  int
	color color.Color
}

func NewNumber(n, s int, c color.Color) *Text {
	return &Text{fmt.Sprint(n), s, c}
}

func NewText(t string, s int, c color.Color) *Text {
	return &Text{t, s, c}
}

func (t *Text) Bounds() (int, int) {
	w, h := text.Measure(t.text, t.font(), 0)
	return int(w), int(h)
}

func (t *Text) Text() string {
	return t.text
}

func (t *Text) SetText(text string) {
	t.text = text
}

func (t *Text) SetColor(c color.Color) {
	t.color = c
}

func (t *Text) font() *text.GoTextFace {
	var source *text.GoTextFaceSource
	source, _ = text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	// if _, err := strconv.Atoi(t.text); err == nil {
	// 	source, _ = text.NewGoTextFaceSource(bytes.NewReader(static.NumberFont))
	// }
	return &text.GoTextFace{
		Source: source,
		Size:   float64(t.size),
	}
}

func (t *Text) Draw(screen *ebiten.Image, x, y int) {
	textOp := &text.DrawOptions{}
	textOp.ColorScale.ScaleWithColor(t.color)
	textOp.LineSpacing = float64(t.size)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	textOp.Filter = ebiten.FilterLinear
	textOp.GeoM.Translate(float64(x+screen.Bounds().Dx()/2), float64(y+screen.Bounds().Dy()/2))
	text.Draw(screen, t.text, t.font(), textOp)
}
