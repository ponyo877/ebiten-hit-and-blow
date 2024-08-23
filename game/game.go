package game

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"

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
	isMyturn           bool
	tmp                *ebiten.Image
	changeHistoryBoard bool
	changePlayerBoard  bool
	changeInput        bool
	changeTimer        bool
	changeMode         bool
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
	return ebiten.RunGame(g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	w, h := screenWidth, screenHeight
	name := "NoName"
	rate := 1500
	reader := bytes.NewReader(static.Profile)
	if g.tmp == nil {
		img, _, _ := image.Decode(reader)
		icon := drawable.NewIcon(w*90/750, w*90/750, img)
		myPlayer := drawable.NewPlayer(w/2, h*90/1334, h*30/1334, icon, name, rate, drawable.TransColor, color.White)
		emPlayer := drawable.NewPlayer(w/2, h*90/1334, h*30/1334, icon, name, rate, drawable.TransColor, color.White)
		myHand := drawable.NewHand([]string{"?", "?", "?"}, w*45/375, h*120/1334, h*100/1334, w*15/375, color.White, drawable.HistoryFrameColor, drawable.HistoryBackgroundColor)
		emHand := drawable.NewHand([]string{"?", "?", "?"}, w*45/375, h*120/1334, h*100/1334, w*15/375, color.White, drawable.HistoryFrameColor, drawable.HistoryBackgroundColor)
		g.playerBoard = drawable.NewPlayerBoard(myPlayer, emPlayer, myHand, emHand, w, h*2/10, drawable.MyPlayerColor, drawable.EnemyPlayerColor)
		myHistory := drawable.NewHistory([]*drawable.Feedback{}, w*350/750, h*55/1334, "あなたの推理", drawable.HistoryFrameColor, color.White)
		emHistory := drawable.NewHistory([]*drawable.Feedback{}, w*350/750, h*55/1334, "相手の推理", drawable.HistoryFrameColor, color.White)
		g.historyBoard = drawable.NewHistoryBoard(myHistory, emHistory, w, h*5/10, drawable.HistoryBackgroundColor)
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
		g.enterKey = drawable.NewEffectButton("決定", w*330/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, w*110/750+wmargin, h*25/40+(h*3/10)*19/40+h*180/1334+2*hmargin, g.inputField)
		g.deleteKey = drawable.NewEffectButton("←", w*110/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, w*440/750+3*wmargin, h*25/40+(h*3/10)*19/40+h*180/1334+2*hmargin, g.inputField)
		g.inputBoard = drawable.NewInputBoard(w, h*4/10, h*30/1334, "相手は考えています...", g.inputField, drawable.HistoryFrameColor, drawable.MessageColor)
		go func(ch chan *entity.Hand) {
			h := <-ch
			g.playerBoard.MyHand().SetHand(h)
			g.changePlayerBoard = true
		}(g.hch)
	}

	switch g.mode {
	case ModeInit:
		g.searchButton = drawable.NewEffectButton("対戦相手を探す", w*600/750, h*90/1334, h*80/1334, drawable.GrayColor, color.White, w*110/750, h*25/40+(h*3/10)*19/40+h*180/1334, g.inputField)
		// ボタンのクリック判定
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if g.searchButton.In(ebiten.CursorPosition()) {
				go conn.Matching(g.cch, g.hch, g.gch, g.qch, g.tch)
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
						if g.isMyturn {
							g.historyBoard.EmHistory().AddFeedback(drawable.NewFeedback(gs, hs))
						} else {
							g.historyBoard.MyHistory().AddFeedback(drawable.NewFeedback(gs, hs))
						}
						g.changeHistoryBoard = true
					}
				}(g.qch)
				go func(ch chan bool) {
					for {
						g.isMyturn = <-ch
						log.Println("g.isMyturn", g.isMyturn)
						// g.changeTimer = true
						// g.timer.Start()
					}
				}(g.tch)
			}
		}

	case ModeWaiting:

	case ModePlaying:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			pushedButton := g.tenkey.WhichButtonByPosition(ebiten.CursorPosition())
			log.Println("pushedButton!!")
			if pushedButton != nil {
				pushedButton.Push()
				pushedButton.Disable()
				g.changeInput = true
			}
			if g.enterKey.In(ebiten.CursorPosition()) {
				g.enterKey.Send(func(ns []int) {
					conn.Send(g.gch, ns)
				})
				g.changeInput = true
			}
			if g.deleteKey.In(ebiten.CursorPosition()) {
				g.deleteKey.Clear()
				for _, nb := range g.numberButtons {
					nb.Enable()
				}
				g.changeInput = true
			}
		}
	case ModeFinished:

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.tmp == nil || g.changeMode {
		g.tmp = ebiten.NewImage(screenWidth, screenHeight)
		g.tmp.Fill(drawable.TransColor)
		g.playerBoard.Draw(g.tmp, 0, 0)
		g.historyBoard.Draw(g.tmp, 0, screenHeight*2/10)
		// g.inputBoard.Draw(g.tmp, 0, screenHeight*25/40)
		// g.changeMode = false
	}
	if g.changePlayerBoard {
		g.playerBoard.Draw(g.tmp, 0, 0)
		g.changePlayerBoard = false
	}
	// インタラクションな描画
	switch g.mode {
	case ModeInit:
		g.searchButton.Draw(g.tmp)
	default:
		if g.changeMode || g.changeInput {
			g.inputBoard.Draw(g.tmp, 0, screenHeight*25/40)
			g.timer.Draw(g.tmp, 0, screenHeight*25/40)
			g.tenkey.Draw(g.tmp)
			g.enterKey.Draw(g.tmp)
			g.deleteKey.Draw(g.tmp)
			g.changeMode = false
			g.changeInput = false
		}
		if g.changeHistoryBoard {
			g.historyBoard.Draw(g.tmp, 0, screenHeight*2/10)
			g.changeHistoryBoard = false
		}
		// if g.changeInput {
		// 	g.inputBoard.Draw(g.tmp, 0, screenHeight*25/40)
		// 	g.changeInput = false
		// }
		if !g.changeTimer {
			g.timer.Draw(g.tmp, 0, screenHeight*25/40)
		}
	}
	ocOp := &ebiten.DrawImageOptions{}
	ocOp.GeoM.Translate(0, 0)
	screen.DrawImage(g.tmp, ocOp)
}
