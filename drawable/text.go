package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

type Text struct {
	text  string
	font  font.Face
	color color.Color
}

func NewText(t string, f font.Face, c color.Color) *Text {
	return &Text{t, f, c}
}

func (t *Text) Bounds() (int, int) {
	w, h := text.Measure(t.text, text.NewGoXFace(t.font), 0)
	return int(w), int(h)
}

func (t *Text) Draw(screen *ebiten.Image, x, y int) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(float64(x)+float64(screen.Bounds().Dx())/2, float64(y)+float64(screen.Bounds().Dy())/2)
	textOp.ColorScale.ScaleWithColor(t.color)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	text.Draw(screen, t.text, text.NewGoXFace(t.font), textOp)
}
