package static

import (
	_ "embed"
)

var (
	//go:embed profile.png
	Profile []byte
	//go:embed number_font.otf
	NumberFont []byte
)
