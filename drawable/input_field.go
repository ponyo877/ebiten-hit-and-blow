package drawable

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type InputField struct {
	w, h     int
	txtSize  int
	bgColor  color.Color
	txtColor color.Color
	numbers  []int
}

func NewInputField(w, h, txtSize int, bgColor, txtColor color.Color, ns []int) *InputField {
	return &InputField{w, h, txtSize, bgColor, txtColor, ns}
}

func (i *InputField) Bounds() (int, int) {
	return i.w * len(i.numbers), i.h
}

func (i *InputField) Draw(screen *ebiten.Image, x, y int) {
	for j, n := range i.numbers {
		rect := NewRect(x+j*i.w, y, i.w, i.h, i.bgColor)
		NewText(fmt.Sprint(n), mplusNormalFont(i.txtSize), i.txtColor).Draw(rect.Image(), 0, 0)
		rect.Draw(screen)
	}
}
