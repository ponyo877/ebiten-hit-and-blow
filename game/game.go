package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() error {
	return ebiten.RunGame(g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(Screen *ebiten.Image) {}
