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
	name         string
	rate         int
	icon         *drawable.Icon
	myPlayer     *drawable.Player
	emPlayer     *drawable.Player
	myNumbers    []int
	emNumbers    []int
	myHand       *drawable.Cards
	emHand       *drawable.Cards
	playerBoard  *drawable.PlayerBoard
	myHistory    *drawable.History
	emHistory    *drawable.History
	historyBoard *drawable.HistoryBoard
	inputBoard   *drawable.InputBoard
}

func (g *Game) init() {
	g.name = "NoName"
	g.rate = 1500
	// file, _ := os.Open("sample.png")
	// buf := new(bytes.Buffer)
	// io.Copy(buf, file)
	// defer file.Close()

	// image, _, err := image.Decode(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// g.icon = drawable.NewIcon(50, 50, image)
	img := image.NewRGBA(image.Rect(0, 0, 50, 50))
	// 短形に色を追加
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{255, 255, 0, 0})
		}
	}
	g.icon = drawable.NewIcon(50, 50, img)
	g.myPlayer = drawable.NewPlayer(150, 50, 10, g.icon, g.name, g.rate, color.White)
	g.emPlayer = drawable.NewPlayer(150, 50, 10, g.icon, g.name, g.rate, color.White)
	cards := []*drawable.Card{}
	g.myNumbers = []int{0, 1, 2}
	for _, n := range g.myNumbers {
		c := drawable.NewNumberCard(n, 50, 75, 40, color.White, color.Black)
		cards = append(cards, c)
	}

	g.myHand = drawable.NewNumberCards(g.myNumbers, 50, 75, 40, 5, color.White, color.Black)
	// g.myHand = drawable.NewHand(myHands, 5)
	cards = []*drawable.Card{}
	g.emNumbers = []int{9, 8, 7}
	for _, n := range g.emNumbers {
		c := drawable.NewNumberCard(n, 50, 75, 40, color.White, color.Black)
		cards = append(cards, c)
	}
	g.emHand = drawable.NewNumberCards(g.emNumbers, 50, 75, 40, 5, color.White, color.Black)
	w, h := g.Layout(420, 600)
	g.playerBoard = drawable.NewPlayerBoard(g.myPlayer, g.emPlayer, g.myHand, g.emHand, w, h/5, color.RGBA{0, 0, 255, 255})
	es := []*drawable.Cards{
		drawable.NewNumberCards([]int{1, 2, 3}, 30, 30, 40, 0, color.White, color.Black),
		drawable.NewNumberCards([]int{4, 5, 6}, 30, 30, 40, 0, color.White, color.Black),
		drawable.NewNumberCards([]int{7, 8, 9}, 30, 30, 40, 0, color.White, color.Black),
	}
	hs := []*drawable.Cards{
		drawable.NewNumberCards([]int{0, 0}, 30, 30, 40, 0, color.White, color.Black),
		drawable.NewNumberCards([]int{1, 0}, 30, 30, 40, 0, color.White, color.Black),
		drawable.NewNumberCards([]int{2, 0}, 30, 30, 40, 0, color.White, color.Black),
	}
	feedback := []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[0]),
		drawable.NewFeedback(es[1], hs[1]),
		drawable.NewFeedback(es[2], hs[2]),
	}
	g.myHistory = drawable.NewHistory(feedback, 150, 30, "あなたの推理", color.White, color.Black)
	feedback = []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[1]),
		drawable.NewFeedback(es[2], hs[0]),
		drawable.NewFeedback(es[1], hs[2]),
	}
	g.emHistory = drawable.NewHistory(feedback, 150, 30, "相手の推理", color.White, color.Black)
	g.historyBoard = drawable.NewHistoryBoard(g.myHistory, g.emHistory, w, h/2, color.RGBA{0, 0, 255, 255})
	timer := drawable.NewTimer(50, 10, 10, color.White, color.Black)
	// drawable.NewInputField(50, 30, 30, color.White, color.Black, []int{0, 1, 2})
	inputField := drawable.NewNumberCards([]int{0, 1, 2}, 50, 30, 30, 0, color.White, color.Black)
	buttons := []*drawable.Button{
		drawable.NewNumberButton(0, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(1, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(2, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(3, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(4, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(5, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(6, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(7, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(8, 30, 30, 30, color.White, color.Black),
		drawable.NewNumberButton(9, 30, 30, 30, color.White, color.Black),
	}
	tenkey := drawable.NewTenkey(buttons, 5)
	g.inputBoard = drawable.NewInputBoard(w, h/4, "相手は考えています...", timer, inputField, tenkey, color.RGBA{0, 0, 255, 255})
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
	if g.inputBoard == nil {
		g.init()
	}
	g.playerBoard.Draw(screen, 0, 0)
	g.historyBoard.Draw(screen, 0, screen.Bounds().Dy()/5)
	g.inputBoard.Draw(screen, 0, screen.Bounds().Dy()/5*3)
}
