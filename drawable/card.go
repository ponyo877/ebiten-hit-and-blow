package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Card struct {
	w, h    int
	txtSize int
	// base    *Rect
	bgColor color.Color
	base    *Rect
	baseV   *ebiten.Image
	text    *Text
}

func NewCard(n string, w, h, ts int, bgc, txc color.Color) *Card {
	base := NewRect(w, h, bgc)
	baseV := ebiten.NewImage(w, h)
	// whiteSubImage := base.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	text := NewText(n, ts, txc)
	return &Card{w, h, ts, bgc, base, baseV, text}
}

func NewNumberCard(n, w, h, t int, bgc, txtc color.Color) *Card {
	return NewCard(fmt.Sprint(n), w, h, t, bgc, txtc)
}

func (c *Card) Bounds() (int, int) {
	return c.w, c.h
}

func (c *Card) Draw(screen *ebiten.Image, x, y int) {
	// c.base.Fill()
	// c.text.Draw(c.base.Image(), 0, 0)
	// c.base.Draw(screen, x, y)

	wf, hf, xf, yf, rf := float32(c.w), float32(c.h), float32(0), float32(0), float32(10)
	var path vector.Path
	path.MoveTo(xf, yf)
	path.ArcTo(xf+wf, yf, xf+wf, yf+hf/2, rf)
	path.ArcTo(xf+wf, yf+hf, xf+wf/2, yf+hf, rf)
	path.ArcTo(xf, yf+hf, xf, yf+hf/2, rf)
	path.ArcTo(xf, yf, xf+wf/2, yf, rf)
	path.Close()

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	op := &ebiten.DrawTrianglesOptions{}
	op.AntiAlias = true
	op.FillRule = ebiten.NonZero
	c.baseV.Fill(c.bgColor)
	c.base.Image().DrawTriangles(vs, is, c.baseV, op)
	c.text.Draw(c.base.Image(), 0, 0)
	c.base.Draw(screen, x, y)
}
