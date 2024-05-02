package game

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ponyo877/ebiten-hit-and-blow/drawable"
)

type Game struct {
	playerBoard  *drawable.PlayerBoard
	historyBoard *drawable.HistoryBoard
	inputBoard   *drawable.InputBoard
}

func (g *Game) init() {
	w, h := g.Layout(375, 667)
	name := "NoName"
	rate := 1500
	img := image.NewRGBA(image.Rect(0, 0, w*90/750, w*90/750))
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{255, 255, 0, 0})
		}
	}
	icon := drawable.NewIcon(w*90/750, w*90/750, img)
	myPlayer := drawable.NewPlayer(w/2, h*90/1334, h*20/1334, icon, name, rate, color.White, color.Black)
	emPlayer := drawable.NewPlayer(w/2, h*90/1334, h*20/1334, icon, name, rate, color.White, color.Black)
	myHand := drawable.NewNumberCards([]int{0, 1, 2}, 45, h*120/1334, h*120/1334, 5, color.White, color.Black)
	emHand := drawable.NewNumberCards([]int{9, 8, 7}, 45, h*120/1334, h*120/1334, 5, color.White, color.Black)
	g.playerBoard = drawable.NewPlayerBoard(myPlayer, emPlayer, myHand, emHand, w, h*2/10, color.RGBA{0, 0, 255, 255})
	es := []*drawable.Cards{
		drawable.NewNumberCards([]int{1, 2, 3}, w*30/750, h*40/1334, h*40/1334, 0, color.White, color.Black),
		drawable.NewNumberCards([]int{4, 5, 6}, w*30/750, h*40/1334, h*40/1334, 0, color.White, color.Black),
		drawable.NewNumberCards([]int{7, 8, 9}, w*30/750, h*40/1334, h*40/1334, 0, color.White, color.Black),
	}
	hs := []*drawable.Cards{
		drawable.NewNumberCards([]int{0, 0}, w*30/750, h*40/1334, h*40/1334, 0, color.White, color.Black),
		drawable.NewNumberCards([]int{1, 0}, w*30/750, h*40/1334, h*40/1334, 0, color.White, color.Black),
		drawable.NewNumberCards([]int{2, 0}, w*30/750, h*40/1334, h*40/1334, 0, color.White, color.Black),
	}
	feedback := []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[0]),
		drawable.NewFeedback(es[1], hs[1]),
		drawable.NewFeedback(es[2], hs[2]),
	}
	myHistory := drawable.NewHistory(feedback, w*350/750, h*45/1334, "あなたの推理", color.White, color.Black)
	feedback = []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[1]),
		drawable.NewFeedback(es[2], hs[0]),
		drawable.NewFeedback(es[1], hs[2]),
	}
	emHistory := drawable.NewHistory(feedback, w*350/750, h*45/1334, "相手の推理", color.White, color.Black)
	g.historyBoard = drawable.NewHistoryBoard(myHistory, emHistory, w, h*5/10, color.RGBA{0, 0, 255, 255})
	timer := drawable.NewTimer(50, w*120/750, w*120/750, color.White, color.Black)
	inputField := drawable.NewNumberCards([]int{0, 1, 2}, w*75/750, h*95/1334, h*95/1334, 0, color.White, color.Black)
	buttons := []*drawable.Button{
		drawable.NewNumberButton(0, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(1, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(2, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(3, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(4, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(5, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(6, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(7, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(8, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
		drawable.NewNumberButton(9, w*110/750, h*90/1334, h*90/1334, color.White, color.Black),
	}
	tenkey := drawable.NewTenkey(buttons, w*12/750, h*16/1334)
	g.inputBoard = drawable.NewInputBoard(w, h*3/10, "相手は考えています...", timer, inputField, tenkey, color.RGBA{0, 0, 255, 255})
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() error {
	w, h := g.Layout(375, 667)
	ebiten.SetWindowSize(w, h)
	return ebiten.RunGame(g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x80, 0x80})
	if g.inputBoard == nil {
		g.init()
	}
	_, h := g.Layout(375, 667)
	g.playerBoard.Draw(screen, 0, 0)
	g.historyBoard.Draw(screen, 0, h*2/10)
	g.inputBoard.Draw(screen, 0, h*7/10)
}
