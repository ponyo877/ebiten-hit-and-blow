package game

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ponyo877/ebiten-hit-and-blow/drawable"
)

type Game struct {
	name         string
	rate         int
	icon         *drawable.Icon
	myPlayer     *drawable.Player
	emPlayer     *drawable.Player
	myNumbers    []int
	emNumbers    []int
	myHand       *drawable.Hand
	emHand       *drawable.Hand
	playerBoard  *drawable.PlayerBoard
	myHistory    *drawable.History
	emHistory    *drawable.History
	historyBoard *drawable.HistoryBoard
}

func (g *Game) init() {
	g.name = "NoName"
	g.rate = 1500
	file, _ := os.Open("sample.png")
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	g.icon = drawable.NewIcon(50, 50, image)
	g.myPlayer = drawable.NewPlayer(150, 50, 10, g.icon, g.name, g.rate, color.White)
	g.emPlayer = drawable.NewPlayer(150, 50, 10, g.icon, g.name, g.rate, color.White)
	cards := []*drawable.Card{}
	g.myNumbers = []int{0, 1, 2}
	for _, n := range g.myNumbers {
		c := drawable.NewNumberCard(n, 50, 75, 40, color.White, color.Black)
		cards = append(cards, c)
	}
	g.myHand = drawable.NewHand(cards, 5)
	cards = []*drawable.Card{}
	g.emNumbers = []int{9, 8, 7}
	for _, n := range g.emNumbers {
		c := drawable.NewNumberCard(n, 50, 75, 40, color.White, color.Black)
		cards = append(cards, c)
	}
	g.emHand = drawable.NewHand(cards, 5)
	w, h := g.Layout(420, 600)
	g.playerBoard = drawable.NewPlayerBoard(g.myPlayer, g.emPlayer, g.myHand, g.emHand, w, h/5, color.RGBA{0, 0, 255, 255})
	es := []*drawable.Estimate{
		drawable.NewEstimate(30, 30, []int{1, 2, 3}, 40, color.White, color.Black),
		drawable.NewEstimate(30, 30, []int{4, 5, 6}, 40, color.White, color.Black),
		drawable.NewEstimate(30, 30, []int{7, 8, 9}, 40, color.White, color.Black),
	}
	hs := []*drawable.Hint{
		drawable.NewHint(30, 30, 0, 0, 40, color.White, color.Black),
		drawable.NewHint(30, 30, 1, 0, 40, color.White, color.Black),
		drawable.NewHint(30, 30, 2, 0, 40, color.White, color.Black),
	}
	feedback := []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[0]),
		drawable.NewFeedback(es[1], hs[1]),
		drawable.NewFeedback(es[2], hs[2]),
	}
	g.myHistory = drawable.NewHistory(150, 30, "あなたの推理", color.White, feedback)
	feedback = []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[1]),
		drawable.NewFeedback(es[2], hs[0]),
		drawable.NewFeedback(es[1], hs[2]),
	}
	g.emHistory = drawable.NewHistory(150, 30, "相手の推理", color.White, feedback)
	g.historyBoard = drawable.NewHistoryBoard(g.myHistory, g.emHistory, w, h/2, color.RGBA{0, 0, 255, 255})

}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() error {
	ebiten.SetWindowSize(420, 600)
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
	if g.playerBoard == nil {
		g.init()
	}
	g.playerBoard.Draw(screen, 0, 0)
	g.historyBoard.Draw(screen, 0, screen.Bounds().Dy()/5)
}
