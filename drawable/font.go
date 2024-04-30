package drawable

import (
	"os"

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

func mplusNormalFont(size int) font.Face {
	var fontsRoGSanSrfStd, _ = os.ReadFile("./drawable/RoGSanSrfStd-Bd.otf")
	// return NewFontOption(fonts.MPlus1pRegular_ttf, size, 72, font.HintingVertical).Font()
	return NewFontOption(fontsRoGSanSrfStd, size, 72, font.HintingVertical).Font()
}

// var fontsRoGSanSrfStd, _ = os.ReadFile("RoGSanSrfStd-Bd.otf")
