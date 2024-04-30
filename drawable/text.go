package drawable

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	text string
	size int
	// font  font.Face
	color color.Color
}

func NewNumber(n, s int, c color.Color) *Text {
	return NewText(fmt.Sprint(n), s, c)
}

func NewText(t string, s int, c color.Color) *Text {
	return &Text{t, s, c}
}

func (t *Text) Bounds() (int, int) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	font := &text.GoTextFace{
		Source: s,
		Size:   float64(t.size),
	}
	w, h := text.Measure(t.text, font, 0)
	return int(w), int(h)
}

func (t *Text) Draw(screen *ebiten.Image, x, y int) {
	textOp := &text.DrawOptions{}
	// textOp.GeoM.Translate(float64(x)+float64(screen.Bounds().Dx())/2, float64(y)+float64(screen.Bounds().Dy())/2)
	textOp.ColorScale.ScaleWithColor(t.color)
	textOp.LineSpacing = float64(t.size)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	textOp.Filter = ebiten.FilterLinear
	// fontsRoGSanSrfStd, _ := os.ReadFile("./drawable/RoGSanSrfStd-Bd.otf")
	// s, err := text.NewGoTextFaceSource(bytes.NewReader(fontsRoGSanSrfStd))
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	font := &text.GoTextFace{
		Source: s,
		Size:   float64(t.size),
	}
	textOp.GeoM.Translate(float64(x+screen.Bounds().Dx()/2), float64(y+screen.Bounds().Dy()/2))
	// text.Draw(screen, t.text, text.NewGoXFace(t.font), textOp)
	text.Draw(screen, t.text, font, textOp)
}
