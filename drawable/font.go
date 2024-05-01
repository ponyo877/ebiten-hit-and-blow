package drawable

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type FontOption struct {
	src       []byte
	size, dpi int
	hinting   font.Hinting
}

func NewFontOption(src []byte, si, d int, h font.Hinting) *FontOption {
	return &FontOption{src, si, d, h}
}

func (f *FontOption) Font() font.Face {
	tt, _ := opentype.Parse(f.src)
	font, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(f.size),
		DPI:     float64(f.dpi),
		Hinting: font.HintingFull,
	})
	return font
}
