//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/ponyo877/ebiten-hit-and-blow/game"
)

func main() {
	g := game.NewGame()
	if err := g.Start(); err != nil {
		panic(err)
	}
}
