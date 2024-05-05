package game

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ponyo877/ebiten-hit-and-blow/drawable"
	"github.com/ponyo877/ebiten-hit-and-blow/static"
)

type Game struct {
	playerBoard  *drawable.PlayerBoard
	historyBoard *drawable.HistoryBoard
	inputBoard   *drawable.InputBoard
	timer        *drawable.Timer
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
	w, h := g.Layout(375, 667)
	name := "NoName"
	rate := 1500
	reader := bytes.NewReader(static.Profile)

	img, _, _ := image.Decode(reader)
	icon := drawable.NewIcon(w*90/750, w*90/750, img)
	myPlayer := drawable.NewPlayer(w/2, h*90/1334, h*30/1334, icon, name, rate, drawable.TransColor, color.White)
	emPlayer := drawable.NewPlayer(w/2, h*90/1334, h*30/1334, icon, name, rate, drawable.TransColor, color.White)
	myHand := drawable.NewNumberHand([]int{0, 1, 2}, 45, h*120/1334, h*100/1334, 15, color.White, drawable.HistoryFrameColor)
	emHand := drawable.NewNumberHand([]int{9, 8, 7}, 45, h*120/1334, h*100/1334, 15, color.White, drawable.HistoryFrameColor)
	g.playerBoard = drawable.NewPlayerBoard(myPlayer, emPlayer, myHand, emHand, w, h*2/10, drawable.MyPlayerColor, drawable.EnemyPlayerColor)
	es := []*drawable.Cards{
		drawable.NewNumberCards([]int{1, 2, 3}, w*30/750, h*40/1334, h*40/1334, 5, drawable.TransColor, drawable.HistoryFrameColor),
		drawable.NewNumberCards([]int{4, 5, 6}, w*30/750, h*40/1334, h*40/1334, 5, drawable.TransColor, drawable.HistoryFrameColor),
		drawable.NewNumberCards([]int{7, 8, 9}, w*30/750, h*40/1334, h*40/1334, 5, drawable.TransColor, drawable.HistoryFrameColor),
	}
	hs := []*drawable.Cards{
		drawable.NewNumberCards([]int{0, 0}, w*30/750, h*40/1334, h*40/1334, 5, drawable.TransColor, drawable.HistoryFrameColor),
		drawable.NewNumberCards([]int{1, 0}, w*30/750, h*40/1334, h*40/1334, 5, drawable.TransColor, drawable.HistoryFrameColor),
		drawable.NewNumberCards([]int{2, 0}, w*30/750, h*40/1334, h*40/1334, 5, drawable.TransColor, drawable.HistoryFrameColor),
	}
	feedback := []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[0], w*350/750, h*55/1334),
		drawable.NewFeedback(es[1], hs[1], w*350/750, h*55/1334),
		drawable.NewFeedback(es[2], hs[2], w*350/750, h*55/1334),
	}
	myHistory := drawable.NewHistory(feedback, w*350/750, h*55/1334, "あなたの推理", drawable.HistoryFrameColor, color.White)
	feedback = []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[1], w*350/750, h*55/1334),
		drawable.NewFeedback(es[2], hs[0], w*350/750, h*55/1334),
		drawable.NewFeedback(es[1], hs[2], w*350/750, h*55/1334),
	}
	emHistory := drawable.NewHistory(feedback, w*350/750, h*55/1334, "相手の推理", drawable.HistoryFrameColor, color.White)
	g.historyBoard = drawable.NewHistoryBoard(myHistory, emHistory, w, h*5/10, drawable.HistoryBackgroundColor)
	g.timer = drawable.NewTimer(w*60/750, w*60/750, w*60/750, color.White, drawable.HistoryFrameColor)

	inputField := drawable.NewNumberInput([]int{0, 1, 2}, w*75/750, h*100/1334, h*100/1334, 30, drawable.HistoryFrameColor, drawable.MessageColor)
	buttons := []*drawable.Button{
		drawable.NewNumberButton(0, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(1, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(2, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(3, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(4, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(5, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(6, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(7, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(8, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
		drawable.NewNumberButton(9, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White),
	}
	tenkey := drawable.NewTenkey(buttons, w*12/750, h*16/1334)
	g.inputBoard = drawable.NewInputBoard(w, h*3/10, "相手は考えています...", inputField, tenkey, drawable.HistoryFrameColor, drawable.MessageColor)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x80, 0x80})
	if g.inputBoard == nil {
		g.Update()
	}
	_, h := g.Layout(375, 667)
	g.playerBoard.Draw(screen, 0, 0)
	g.historyBoard.Draw(screen, 0, h*2/10)
	g.inputBoard.Draw(screen, 0, h*7/10)
	g.timer.Draw(screen, 0, h*7/10)
}
