package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerBoard struct {
	myPlayer *Player
	emPlayer *Player
	myHand   *Hand
	emHand   *Hand
	w, h     int
	myBase   *Rect
	emBase   *Rect
}

func NewPlayerBoard(mp, ep *Player, mh, eh *Hand, w, h int, mybgc, embgc color.Color) *PlayerBoard {
	myBase := NewRect(w/2, h, mybgc)
	emBase := NewRect(w/2, h, embgc)
	return &PlayerBoard{mp, ep, mh, eh, w, h, myBase, emBase}
}

func (pb *PlayerBoard) MyHand() *Hand {
	return pb.myHand
}

func (pb *PlayerBoard) Draw(screen *ebiten.Image, x, y int) {
	pb.myBase.Fill()
	pb.myPlayer.Draw(pb.myBase.Image(), 0, 0)
	w, _ := pb.myHand.Bounds()
	pb.myHand.Draw(pb.myBase.Image(), pb.w/4-w/2, pb.myPlayer.h+pb.h*20/133)

	pb.emBase.Fill()
	pb.emPlayer.Draw(pb.emBase.Image(), 0, 0)
	w, _ = pb.emHand.Bounds()
	pb.emHand.Draw(pb.emBase.Image(), pb.w/4-w/2, pb.emPlayer.h+pb.h*20/133)

	myOp := &ebiten.DrawImageOptions{}
	myOp.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(pb.myBase.Image(), myOp)
	emOp := &ebiten.DrawImageOptions{}
	emOp.GeoM.Translate(float64(x+pb.w/2), float64(y))
	screen.DrawImage(pb.emBase.Image(), emOp)
}
