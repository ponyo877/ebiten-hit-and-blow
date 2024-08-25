package static

import (
	_ "embed"
)

var (
	//go:embed me.png
	Me []byte
	//go:embed enemy.png
	Enemy []byte
	//go:embed number_font.otf
	NumberFont []byte
)
