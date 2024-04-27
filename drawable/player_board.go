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
	bgColor  color.Color
}

func NewPlayerBoard(mp, ep *Player, mh, eh *Hand, w, h int, bgc color.Color) *PlayerBoard {
	return &PlayerBoard{mp, ep, mh, eh, w, h, bgc}
}

func (pb *PlayerBoard) Image() *ebiten.Image {
	img := ebiten.NewImage(pb.w, pb.h)
	img.Fill(pb.bgColor)
	return img

}

func (pb *PlayerBoard) Draw(screen *ebiten.Image, x, y int) {
	img := pb.Image()
	pb.myPlayer.Draw(img, 0, 0)
	pb.myHand.Draw(img, 0, 0+pb.myPlayer.h)

	cx := img.Bounds().Dx() / 2
	pb.emPlayer.Draw(img, cx, 0)
	pb.emHand.Draw(img, cx, 0+pb.emPlayer.h)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
}
