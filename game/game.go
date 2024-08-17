package game

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ponyo877/ebiten-hit-and-blow/drawable"
	"github.com/ponyo877/ebiten-hit-and-blow/static"
)

type Game struct {
	playerBoard   *drawable.PlayerBoard
	historyBoard  *drawable.HistoryBoard
	inputBoard    *drawable.InputBoard
	timer         *drawable.Timer
	inputField    *drawable.Input
	tenkey        *drawable.Tenkey
	enterKey      *drawable.EffectButton
	deleteKey     *drawable.EffectButton
	numberButtons []*drawable.Button
	tmp           *ebiten.Image
	changeInput   bool
	changeTimer   bool
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() error {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g.changeInput = true
	return ebiten.RunGame(g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	if g.tmp == nil {
		w, h := screenWidth, screenHeight
		name := "NoName"
		rate := 1500
		reader := bytes.NewReader(static.Profile)

		img, _, _ := image.Decode(reader)
		icon := drawable.NewIcon(w*90/750, w*90/750, img)
		myPlayer := drawable.NewPlayer(w/2, h*90/1334, h*30/1334, icon, name, rate, drawable.TransColor, color.White)
		emPlayer := drawable.NewPlayer(w/2, h*90/1334, h*30/1334, icon, name, rate, drawable.TransColor, color.White)
		myHand := drawable.NewNumberHand([]int{0, 1, 2}, w*45/375, h*120/1334, h*100/1334, w*15/375, color.White, drawable.HistoryFrameColor, drawable.HistoryBackgroundColor)
		emHand := drawable.NewNumberHand([]int{9, 8, 7}, w*45/375, h*120/1334, h*100/1334, w*15/375, color.White, drawable.HistoryFrameColor, drawable.HistoryBackgroundColor)
		g.playerBoard = drawable.NewPlayerBoard(myPlayer, emPlayer, myHand, emHand, w, h*2/10, drawable.MyPlayerColor, drawable.EnemyPlayerColor)
		es := []*drawable.Cards{
			drawable.NewNumberCards([]int{1, 2, 3}, w*30/750, h*40/1334, h*40/1334, w*5/375, drawable.TransColor, drawable.HistoryFrameColor),
			drawable.NewNumberCards([]int{4, 5, 6}, w*30/750, h*40/1334, h*40/1334, w*5/375, drawable.TransColor, drawable.HistoryFrameColor),
			drawable.NewNumberCards([]int{7, 8, 9}, w*30/750, h*40/1334, h*40/1334, w*5/375, drawable.TransColor, drawable.HistoryFrameColor),
		}
		hs := []*drawable.Cards{
			drawable.NewNumberCards([]int{0, 0}, w*30/750, h*40/1334, h*40/1334, w*5/375, drawable.TransColor, drawable.HistoryFrameColor),
			drawable.NewNumberCards([]int{1, 0}, w*30/750, h*40/1334, h*40/1334, w*5/375, drawable.TransColor, drawable.HistoryFrameColor),
			drawable.NewNumberCards([]int{2, 0}, w*30/750, h*40/1334, h*40/1334, w*5/375, drawable.TransColor, drawable.HistoryFrameColor),
		}
		feedback := []*drawable.Feedback{
			drawable.NewFeedback(es[0], hs[0]),
			drawable.NewFeedback(es[1], hs[1]),
			drawable.NewFeedback(es[2], hs[2]),
		}
		myHistory := drawable.NewHistory(feedback, w*350/750, h*55/1334, "あなたの推理", drawable.HistoryFrameColor, color.White)
		feedback = []*drawable.Feedback{
			drawable.NewFeedback(es[0], hs[1]),
			drawable.NewFeedback(es[2], hs[0]),
			drawable.NewFeedback(es[1], hs[2]),
		}
		emHistory := drawable.NewHistory(feedback, w*350/750, h*55/1334, "相手の推理", drawable.HistoryFrameColor, color.White)
		g.historyBoard = drawable.NewHistoryBoard(myHistory, emHistory, w, h*5/10, drawable.HistoryBackgroundColor)
		g.timer = drawable.NewTimer(w*60/750, w*60/750, w*60/750, color.White, drawable.HistoryFrameColor)

		g.inputField = drawable.NewInput([]string{}, w*75/750, h*100/1334, h*100/1334, w*30/375, drawable.HistoryFrameColor, drawable.MessageColor)
		g.numberButtons = []*drawable.Button{
			drawable.NewNumberButton(0, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(1, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(2, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(3, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(4, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(5, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(6, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(7, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(8, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(9, w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, g.inputField),
		}
		wmargin := w * 12 / 750
		hmargin := h * 16 / 1334
		g.tenkey = drawable.NewTenkey(g.numberButtons, w*12/750, h*16/1334, w, h*25/40+(h*3/10)*19/40)
		g.enterKey = drawable.NewEffectButton("決定", w*330/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, w*110/750+wmargin, h*25/40+(h*3/10)*19/40+h*180/1334+2*hmargin, g.inputField)
		g.deleteKey = drawable.NewEffectButton("←", w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, w*440/750+3*wmargin, h*25/40+(h*3/10)*19/40+h*180/1334+2*hmargin, g.inputField)
		g.inputBoard = drawable.NewInputBoard(w, h*4/10, h*30/1334, "相手は考えています...", g.inputField, drawable.HistoryFrameColor, drawable.MessageColor)
	}

	// ボタンのクリック判定
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		pushedButton := g.tenkey.WhichButtonByPosition(ebiten.CursorPosition())
		if pushedButton != nil {
			pushedButton.Push()
			g.changeInput = true
		}
		if g.enterKey.In(ebiten.CursorPosition()) {
			g.enterKey.Send(func(ns []int) {
				// ここで入力された数字を送信する処理
			})
			g.changeInput = true
		}
		if g.deleteKey.In(ebiten.CursorPosition()) {
			g.deleteKey.Clear()
			g.changeInput = true
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.tmp == nil {
		g.tmp = ebiten.NewImage(screenWidth, screenHeight)
		g.tmp.Fill(drawable.TransColor)
		g.playerBoard.Draw(g.tmp, 0, 0)
		g.historyBoard.Draw(g.tmp, 0, screenHeight*2/10)
		g.inputBoard.Draw(g.tmp, 0, screenHeight*25/40)
		g.timer.Draw(g.tmp, 0, screenHeight*25/40)
	}
	if g.changeInput {
		g.inputBoard.Draw(g.tmp, 0, screenHeight*25/40)
		g.changeInput = false
	}
	// delete not
	if !g.changeTimer {
		g.timer.Draw(g.tmp, 0, screenHeight*25/40)
	}
	ocOp := &ebiten.DrawImageOptions{}
	ocOp.GeoM.Translate(0, 0)
	screen.DrawImage(g.tmp, ocOp)

	// インタラクションな描画
	g.tenkey.Draw(screen)
	g.enterKey.Draw(screen)
	g.deleteKey.Draw(screen)
}
