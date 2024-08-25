package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Rounded struct {
	w, h    int
	bgColor color.Color
}

var (
	r          float32 = 10
	whiteImage         = ebiten.NewImage(3, 3)
	// whiteSubImage         = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func NewRounded(w, h int, bgc color.Color) *Rounded {
	return &Rounded{w, h, bgc}
}

func (c *Rounded) Bounds() (int, int) {
	return c.w, c.h
}

func (c *Rounded) SetColor(bgc color.Color) {
	c.bgColor = bgc
}

func (c *Rounded) Draw(screen *ebiten.Image, x, y int) {
	wf, hf, xf, yf := float32(c.w), float32(c.h), float32(x), float32(y)
	var path vector.Path
	path.MoveTo(xf, yf)
	path.ArcTo(xf+wf, yf, xf+wf, yf+hf/2, r)
	path.ArcTo(xf+wf, yf+hf, xf+wf/2, yf+hf, r)
	path.ArcTo(xf, yf+hf, xf, yf+hf/2, r)
	path.ArcTo(xf, yf, xf+wf/2, yf, r)
	path.Close()

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	whiteImage.Fill(c.bgColor)
	op := &ebiten.DrawTrianglesOptions{}
	op.FillRule = ebiten.NonZero
	op.AntiAlias = true
	screen.DrawTriangles(vs, is, whiteImage, op)
}
