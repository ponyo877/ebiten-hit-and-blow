//go:build js && wasm
// +build js,wasm

package game

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ponyo877/ebiten-hit-and-blow/conn"
	"github.com/ponyo877/ebiten-hit-and-blow/drawable"
	"github.com/ponyo877/ebiten-hit-and-blow/entity"
	"github.com/ponyo877/ebiten-hit-and-blow/static"
)

type Mode int

const (
	ModeInit Mode = iota
	ModeWaiting
	ModePlaying
	ModeFinished
)

type Game struct {
	mode               Mode
	cch                chan struct{}
	hch                chan *entity.Hand
	gch                chan *entity.Guess
	qch                chan *entity.QA
	tch                chan bool
	tich               chan int
	jch                chan entity.JudgeStatus
	rch                chan *entity.Rating
	playerBoard        *drawable.PlayerBoard
	historyBoard       *drawable.HistoryBoard
	inputBoard         *drawable.InputBoard
	timer              *drawable.Timer
	inputField         *drawable.Input
	tenkey             *drawable.Tenkey
	enterKey           *drawable.EffectButton
	deleteKey          *drawable.EffectButton
	numberButtons      []*drawable.NumberButton
	searchButton       *drawable.EffectButton
	message            *drawable.Message
	resultText         *drawable.Text
	isMyTurn           bool
	isMyTurnInit       *bool
	bliCounter         int
	aniCounter         int
	tmp                *ebiten.Image
	changeHistoryBoard bool
	changePlayerBoard  bool
	isJustPushed       bool
	isJustCleared      bool
	changeTimer        bool
	changeMode         bool
	changeTurn         bool
	touchIDs           []ebiten.TouchID
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() error {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g.mode = ModeInit
	g.cch = make(chan struct{})
	g.hch = make(chan *entity.Hand)
	g.gch = make(chan *entity.Guess)
	g.qch = make(chan *entity.QA)
	g.tch = make(chan bool)
	g.tich = make(chan int)
	g.jch = make(chan entity.JudgeStatus)
	g.rch = make(chan *entity.Rating)
	return ebiten.RunGame(g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	w, h := screenWidth, screenHeight
	if g.tmp == nil {
		name := ""
		rate := ""
		reader := bytes.NewReader(static.Me)
		img, _, _ := image.Decode(reader)
		myIcon := drawable.NewIcon(w*90/750, w*90/750, img)
		reader = bytes.NewReader(static.Enemy)
		img, _, _ = image.Decode(reader)
		emIcon := drawable.NewIcon(w*90/750, w*90/750, img)
		myPlayer := drawable.NewPlayer(w/2, h*90/1334, h*30/1334, myIcon, name, rate, drawable.TransColor, color.White)
		emPlayer := drawable.NewPlayer(w/2, h*90/1334, h*30/1334, emIcon, name, rate, drawable.TransColor, color.White)
		myHand := drawable.NewHand([]string{"?", "?", "?"}, w*45/375, h*120/1334, h*100/1334, w*15/375, color.White, drawable.HistoryFrameColor, drawable.HistoryBackgroundColor)
		emHand := drawable.NewHand([]string{"?", "?", "?"}, w*45/375, h*120/1334, h*100/1334, w*15/375, color.White, drawable.HistoryFrameColor, drawable.HistoryBackgroundColor)
		myInitT := drawable.NewCard("??", w*45/375, h*20/1334, h*20/1334, drawable.HistoryFrameColor, color.White)
		emInitT := drawable.NewCard("??", w*45/375, h*20/1334, h*20/1334, drawable.HistoryFrameColor, color.White)
		g.playerBoard = drawable.NewPlayerBoard(myPlayer, emPlayer, myHand, emHand, myInitT, emInitT, w, h*2/10, drawable.MyPlayerColor, drawable.EnemyPlayerColor)
		myHistory := drawable.NewHistory([]*drawable.Feedback{}, w*350/750, h*55/1334, "あなたの推理", drawable.HistoryFrameColor, color.White)
		emHistory := drawable.NewHistory([]*drawable.Feedback{}, w*350/750, h*55/1334, "相手の推理", drawable.HistoryFrameColor, color.White)
		g.historyBoard = drawable.NewHistoryBoard(myHistory, emHistory, w, h*17/40, drawable.HistoryBackgroundColor)
		g.timer = drawable.NewTimer(w*60/750, w*60/750, w*60/750, color.White, drawable.HistoryFrameColor)
		g.inputField = drawable.NewInput([]string{}, w*75/750, h*100/1334, h*100/1334, w*30/375, drawable.HistoryFrameColor, drawable.MessageColor)
		g.numberButtons = []*drawable.NumberButton{
			drawable.NewNumberButton(0, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(1, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(2, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(3, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(4, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(5, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(6, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(7, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(8, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
			drawable.NewNumberButton(9, w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, g.inputField),
		}
		wmargin := w * 12 / 750
		hmargin := h * 16 / 1334
		g.tenkey = drawable.NewTenkey(g.numberButtons, w*12/750, h*16/1334, w, h*25/40+(h*3/10)*19/40)
		g.enterKey = drawable.NewEffectButton("決定", w*330/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, w*110/750+wmargin, h*25/40+(h*3/10)*19/40+h*180/1334+2*hmargin, g.inputField)
		g.deleteKey = drawable.NewEffectButton("←", w*110/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, w*440/750+3*wmargin, h*25/40+(h*3/10)*19/40+h*180/1334+2*hmargin, g.inputField)
		g.enterKey.Disable()
		g.inputBoard = drawable.NewInputBoard(w, h*15/40, drawable.HistoryFrameColor)
		g.message = drawable.NewMessage("対戦相手を探してください", w*30/750, screenWidth/2, w*30/750, drawable.MessageColor, drawable.HistoryFrameColor) // y :=  screenHeight*25/40+(h*4/10)/20
		g.resultText = drawable.NewText("", w*100/750, color.White)
		g.searchButton = drawable.NewEffectButton("対戦相手を探す", w*600/750, h*90/1334, h*80/1334, drawable.HistoryBackgroundColor, drawable.GrayColor, color.White, w/2, h*25/40+(h*3/10)*19/40, g.inputField)
		go func(ch chan *entity.Hand) {
			h := <-ch
			g.playerBoard.MyHand().SetHand(h)
			g.changePlayerBoard = true
		}(g.hch)
	}

	switch g.mode {
	case ModeInit:
		var x, y int
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y = ebiten.CursorPosition()
		}
		g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
		for _, touchID := range g.touchIDs {
			x, y = ebiten.TouchPosition(touchID)
		}
		if x != 0 && y != 0 {
			// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if g.searchButton.InWithCenter(x, y) {
				g.changeMode = true
				g.mode = ModeWaiting
				g.message.SetMessage("対戦相手を探しています...")
				go conn.Matching(g.cch, g.hch, g.gch, g.qch, g.tch, g.tich, g.jch, g.rch)
				go func(ch chan struct{}) {
					<-ch
					g.changeMode = true
					g.mode = ModePlaying
				}(g.cch)
				go func(ch chan *entity.QA) {
					for {
						qa := <-ch
						gs := drawable.NewNumberCards(qa.GuessView(), w*30/750, h*40/1334, h*40/1334, w*5/375, drawable.TransColor, drawable.HistoryFrameColor)
						hs := drawable.NewNumberCards(qa.HistoryView(), w*30/750, h*40/1334, h*40/1334, w*5/375, drawable.TransColor, drawable.HistoryFrameColor)
						if g.isMyTurn {
							g.historyBoard.EmHistory().Add(drawable.NewFeedback(gs, hs))
						} else {
							g.historyBoard.MyHistory().Add(drawable.NewFeedback(gs, hs))
						}
						g.changeHistoryBoard = true
					}
				}(g.qch)
				go func(ch chan bool, rch chan *entity.Rating) {
					for {
						g.isMyTurn = <-ch
						g.timer.Set(60)
						if g.isMyTurnInit == nil {
							g.isMyTurnInit = &g.isMyTurn
							myTurnText, emTurnText := "後攻", "先攻"
							if *g.isMyTurnInit {
								myTurnText, emTurnText = "先攻", "後攻"
							}
							g.playerBoard.MyInitTurn().SetText(myTurnText)
							g.playerBoard.EmInitTurn().SetText(emTurnText)
							myr := <-rch
							emr := <-rch
							g.playerBoard.MyPlayer().SetName(myr.ID())
							g.playerBoard.MyPlayer().SetRate(myr.Rating())
							g.playerBoard.EmPlayer().SetName(emr.ID())
							g.playerBoard.EmPlayer().SetRate(emr.Rating())
							g.changePlayerBoard = true
						}
						g.message.SetMessage("相手は考えています...")
						if g.isMyTurn {
							g.message.SetMessage("推理した3桁の数字を入れてください")
						}
						g.changeTurn = true
						g.changeTimer = true
						log.Println("g.isMyTurn", g.isMyTurn)
					}
				}(g.tch, g.rch)
				go func(ch chan int) {
					for {
						g.timer.Set(<-ch)
						g.changeTimer = true
					}
				}(g.tich)
				go func(ch chan entity.JudgeStatus) {
					for {
						switch <-ch {
						case entity.Win:
							g.message.SetMessage("あなたの勝ちです")
							g.resultText.SetText("YOU WIN!")
						case entity.Lose:
							g.message.SetMessage("あなたの負けです")
							g.resultText.SetText("YOU LOSE!")
						case entity.Draw:
							g.message.SetMessage("引き分けです")
							g.resultText.SetText("DRAW!")
						}
						g.changeMode = true
						g.mode = ModeFinished
					}
				}(g.jch)
			}
		}
	case ModePlaying:
		var x, y int
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y = ebiten.CursorPosition()
		}
		g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
		for _, touchID := range g.touchIDs {
			x, y = ebiten.TouchPosition(touchID)
		}
		if x != 0 && y != 0 {
			pushedButton := g.tenkey.WhichButtonByPosition(x, y)
			if !g.isMyTurn {
				return nil
			}
			if pushedButton != nil && g.inputField.Addble() {
				pushedButton.Push()
				pushedButton.Disable()
				if !g.inputField.Addble() {
					g.enterKey.Enable()
				}
				g.isJustPushed = true
				return nil
			}
			if g.enterKey.In(x, y) && !g.inputField.Addble() {
				if g.inputField.Addble() {
					return nil
				}
				g.enterKey.Send(func(ns []int) {
					conn.Send(g.gch, ns)
				})
				for _, nb := range g.numberButtons {
					nb.Enable()
				}
				g.enterKey.Disable()
				g.isJustCleared = true
				return nil
			}
			if g.deleteKey.In(x, y) {
				g.deleteKey.Clear()
				for _, nb := range g.numberButtons {
					nb.Enable()
				}
				g.enterKey.Disable()
				g.isJustCleared = true
				return nil
			}
		}
	case ModeFinished:
		go func(ch chan *entity.Hand) {
			h := <-ch
			g.playerBoard.EmHand().SetHand(h)
			g.changePlayerBoard = true
		}(g.hch)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.tmp == nil || g.changeMode {
		g.tmp = ebiten.NewImage(screenWidth, screenHeight)
		if !g.changePlayerBoard {
			g.playerBoard.Draw(g.tmp, 0, 0)
		}
		if !g.changeHistoryBoard {
			g.historyBoard.Draw(g.tmp, 0, screenHeight*2/10)
		}
		g.inputBoard.Draw(g.tmp, 0, screenHeight*25/40)
	}
	if g.changePlayerBoard {
		g.playerBoard.Draw(g.tmp, 0, 0)
		g.changePlayerBoard = false
	}
	if g.changeHistoryBoard {
		g.historyBoard.Draw(g.tmp, 0, screenHeight*2/10)
		g.changeHistoryBoard = false
	}
	// インタラクションな描画
	switch g.mode {
	case ModeInit:
		g.searchButton.DrawCenter(g.tmp)
	case ModeWaiting:
		g.changeMode = false
	case ModePlaying:
		if g.isJustPushed {
			g.inputField.Draw(g.tmp, 0, screenHeight*27/40)
			g.tenkey.DrawPart(g.tmp, g.inputField.EndNumber())
			g.enterKey.Draw(g.tmp)
			g.isJustPushed = false
		}
		if g.isJustCleared {
			g.inputField.Draw(g.tmp, 0, screenHeight*27/40)
			g.tenkey.Draw(g.tmp)
			g.enterKey.Draw(g.tmp)
			g.isJustCleared = false
		}
		if g.changeMode {
			g.inputBoard.Draw(g.tmp, 0, screenHeight*25/40)
			g.inputField.Draw(g.tmp, 0, screenHeight*27/40)
			g.timer.Draw(g.tmp, 0, screenHeight*25/40)
			g.tenkey.Draw(g.tmp)
			g.enterKey.Draw(g.tmp)
			g.deleteKey.Draw(g.tmp)
			g.changeMode = false
		}
		if g.changeTimer {
			g.timer.Draw(g.tmp, 0, screenHeight*25/40)
			g.changeTimer = false
		}
	case ModeFinished:
		g.resultText.Draw(g.tmp, 0, screenHeight/4)
	}
	screen.DrawImage(g.tmp, &ebiten.DrawImageOptions{})
	g.counterHelper(func() {
		ocOp2 := &ebiten.DrawImageOptions{}
		ocOp2.GeoM.Translate(0, float64(screenHeight*25/40))
		tmp2 := ebiten.NewImage(screenWidth, screenWidth*50/750)
		g.message.Draw(tmp2)
		screen.DrawImage(tmp2, ocOp2)
		if g.mode == ModeInit {
			return
		}
		ocOp3 := &ebiten.DrawImageOptions{}
		tmp3 := &ebiten.Image{}
		if g.isMyTurn {
			w, h, x, y := g.historyBoard.MyHistory().WaitingRect()
			ocOp3.GeoM.Translate(x+5, y+float64(screenHeight*2/10+5))
			tmp3 = ebiten.NewImage(w, h)
			g.historyBoard.MyHistory().DrawWaiting(tmp3)
		} else {
			w, h, x, y := g.historyBoard.EmHistory().WaitingRect()
			ocOp3.GeoM.Translate(x+5+float64(screenWidth/2), y+float64(screenHeight*2/10+5))
			tmp3 = ebiten.NewImage(w, h)
			g.historyBoard.EmHistory().DrawWaiting(tmp3)
		}
		screen.DrawImage(tmp3, ocOp3)
	},
		func() {
			g.aniCounter++
			text := drawable.NewText("相手のターンです", 30, drawable.HistoryFrameColor)
			if g.isMyTurn {
				text = drawable.NewText("あなたのターンです", 30, drawable.HistoryFrameColor)
			}
			w, _ := text.Bounds()
			tmp := ebiten.NewImage(w, 30)
			tmp.Fill(color.White)
			aniOp := &ebiten.DrawImageOptions{}
			aniOp.GeoM.Translate(25*math.Cos(float64(g.aniCounter))+float64(screenWidth/2-w/2), float64(screenHeight/2))

			text.Draw(tmp, 0, 0)
			screen.DrawImage(tmp, aniOp)
		},
	)
}

func (g *Game) counterHelper(blinkingFunc func(), animationFunc func()) {
	g.bliCounter++
	bliMaxCount := 20
	aniMaxCount := 20
	lightRatio := 0.7
	changeCount := int(float64(bliMaxCount) * lightRatio)
	if !g.changeTurn && g.bliCounter <= changeCount {
		blinkingFunc()
	}
	if g.changeTurn && g.mode == ModePlaying {
		animationFunc()
	}

	if g.bliCounter > bliMaxCount {
		g.bliCounter = 0
	}
	if g.aniCounter > aniMaxCount {
		g.aniCounter = 0
		g.changeTurn = false
	}
}
