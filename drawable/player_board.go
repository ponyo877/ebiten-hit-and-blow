package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerBoard struct {
	myPlayer *Player
	emPlayer *Player
	myHand   *Cards
	emHand   *Cards
	w, h     int
	bgColor  color.Color
	base     *Rect
}

func NewPlayerBoard(mp, ep *Player, mh, eh *Cards, w, h int, bgc color.Color) *PlayerBoard {
	base := NewRect(w, h, bgc)
	return &PlayerBoard{mp, ep, mh, eh, w, h, bgc, base}
}

func (pb *PlayerBoard) Draw(screen *ebiten.Image, x, y int) {
	pb.base.Fill()
	pb.myPlayer.Draw(pb.base.Image(), 0, 0)
	pb.myHand.Draw(pb.base.Image(), 0, 0+pb.myPlayer.h)

	pb.emPlayer.Draw(pb.base.Image(), pb.w/2, 0)
	pb.emHand.Draw(pb.base.Image(), pb.w/2, 0+pb.emPlayer.h)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(pb.base.Image(), op)
}
