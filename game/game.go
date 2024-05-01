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
	name := "NoName"
	rate := 1500
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
	icon := drawable.NewIcon(50, 50, img)
	myPlayer := drawable.NewPlayer(150, 50, 10, icon, name, rate, color.White, color.Black)
	emPlayer := drawable.NewPlayer(150, 50, 10, icon, name, rate, color.White, color.Black)
	myHand := drawable.NewNumberCards([]int{0, 1, 2}, 50, 75, 40, 5, color.White, color.Black)
	emHand := drawable.NewNumberCards([]int{9, 8, 7}, 50, 75, 40, 5, color.White, color.Black)
	w, h := g.Layout(420, 600)
	g.playerBoard = drawable.NewPlayerBoard(myPlayer, emPlayer, myHand, emHand, w, h/5, color.RGBA{0, 0, 255, 255})
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
	myHistory := drawable.NewHistory(feedback, 150, 30, "あなたの推理", color.White, color.Black)
	feedback = []*drawable.Feedback{
		drawable.NewFeedback(es[0], hs[1]),
		drawable.NewFeedback(es[2], hs[0]),
		drawable.NewFeedback(es[1], hs[2]),
	}
	emHistory := drawable.NewHistory(feedback, 150, 30, "相手の推理", color.White, color.Black)
	g.historyBoard = drawable.NewHistoryBoard(myHistory, emHistory, w, h/2, color.RGBA{0, 0, 255, 255})
	timer := drawable.NewTimer(50, 10, 10, color.White, color.Black)
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
