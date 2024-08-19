package conn

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/pion/webrtc/v3"
	"github.com/ponyo877/ebiten-hit-and-blow/entity"
	"github.com/ponyo877/ebiten-hit-and-blow/go-ayame"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

var (
	wsScheme          string
	httpScheme        string
	matchmakingOrigin string
	signalingOrigin   string
	ratingOrigin      string
	solt              string
	recentGuess       *entity.Guess
)

type mmReqMsg struct {
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type mmResMsg struct {
	Type      string    `json:"type"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type getRatingResMsg struct {
	Player1 struct {
		ID   string `json:"id"`
		Rate int    `json:"rate"`
	} `json:"player1"`
	Player2 struct {
		ID   string `json:"id"`
		Rate int    `json:"rate"`
	} `json:"player2"`
}

type updateRatingResMsg struct {
	MatchID  string `json:"match_id"`
	PlayerID string `json:"player_id"`
	Number   int    `json:"number"`
	Hash     string `json:"hash"`
	Result   string `json:"result"`
}

func Matching(hch chan *entity.Hand, gch chan *entity.Guess) {
	mmURL := url.URL{Scheme: wsScheme, Host: matchmakingOrigin, Path: "/matchmaking"}
	signalingURL := url.URL{Scheme: wsScheme, Host: signalingOrigin, Path: "/signaling"}
	ratingURL := url.URL{Scheme: httpScheme, Host: ratingOrigin, Path: "/rating"}

	now := time.Now()
	// window := js.Global().Get("window")
	// localStorage := window.Get("localStorage")

	var userID, hash string
	// userIDjs := localStorage.Get("userID")
	// userID = userIDjs.String()
	// hash = localStorage.Get("hash").String()
	// if userIDjs.Equal(js.Undefined()) {
	userID = shortHash(now)
	// localStorage.Set("userID", userID)
	// localStorage.Set("hash", sha256Hash(solt+userID+solt))
	hash = sha256Hash(solt + userID + solt)
	// }
	reqMsg, err := json.Marshal(mmReqMsg{
		UserID:    userID,
		CreatedAt: now,
	})
	if err != nil {
		log.Fatal(err)
	}
	var resMsg mmResMsg
	var dc *webrtc.DataChannel
	// ch := make(chan *entity.Guess)
	defer func() {
		if dc != nil {
			dc.Close()
		}
	}()
	var conn *ayame.Connection
	connected := make(chan bool)
	board := entity.NewBoard()
	go func() {
		ws, _, err := websocket.Dial(context.Background(), mmURL.String(), nil)
		if err != nil {
			log.Fatal(err)
		}
		defer ws.Close(websocket.StatusNormalClosure, "close connection")

		if err := ws.Write(context.Background(), websocket.MessageText, reqMsg); err != nil {
			log.Fatal(err)
		}
		logElem("[Sys]: Waiting match...\n")
		for {
			if err := wsjson.Read(context.Background(), ws, &resMsg); err != nil {
				log.Fatal(err)
				break
			}
			if resMsg.Type == "MATCH" {
				break
			}
		}
		ws.Close(websocket.StatusNormalClosure, "close connection")
		if resMsg.Type == "MATCH" {

			conn = ayame.NewConnection(signalingURL.String(), resMsg.RoomID, ayame.DefaultOptions(), false, false)
			conn.OnOpen(func(metadata *interface{}) {
				var err error
				dc, err = conn.CreateDataChannel("matchmaking-hit-and-blow", nil)
				if errors.Is(err, ayame.ErrorClientDoesNotExist) {
					return
				}
				if err != nil {
					log.Printf("CreateDataChannel error: %v", err)
					return
				}

				log.Printf("CreateDataChannel: label=%s", dc.Label())
				go func(ch chan *entity.Hand) {
					rand.NewSource(time.Now().UnixNano())
					seed := rand.Int()

					initTurn := entity.NewTurnBySeed(seed)
					myHand := entity.NewHandBySeed(seed)
					// setHand(true, myHand)
					ch <- myHand
					log.Printf("myHand(opener): %v", myHand)
					board.Start(myHand, initTurn, 1)
					if board.IsMyTurnInit() {
						log.Printf("YOU FIRST !!!")
						// setTurn("It's Your Turn !")
					}
					turn := int(initTurn)
					// myRate, opRate, err := getRating(ratingURL, userID, resMsg.UserID)
					// if err != nil {
					// 	log.Printf("failed to get rating: %v", err)
					// 	return
					// }
					// setProfile(userID, resMsg.UserID, myRate, opRate)
					startMsg := Message{Type: "start", Turn: &turn}
					by, _ := json.Marshal(startMsg)
					time.Sleep(10 * time.Second)
					log.Printf("startMsg(opener): %v", string(by))
					if err := dc.SendText(string(by)); err != nil {
						log.Printf("failed to send startMsg: %v", err)
						return
					}
				}(hch)
				finChan := make(chan struct{})
				dc.OnMessage(onMessage(dc, hch, gch, finChan, board))
				go func() {
					select {
					case <-finChan:
						if err := updateRating(ratingURL, resMsg.RoomID, userID, hash, 1, board.Result()); err != nil {
							log.Printf("failed to update rating: %v", err)
							return
						}
					}
				}()
			})

			conn.OnConnect(func() {
				logElem("[Sys]: Matching! Start P2P chat not via server\n")
				conn.CloseWebSocketConnection()
				connected <- true
			})

			conn.OnDataChannel(func(c *webrtc.DataChannel) {
				log.Printf("OnDataChannel: label=%s", c.Label())
				if dc == nil {
					dc = c
				}
				log.Println("ready to recieve")
				// myRate, opRate, err := getRating(ratingURL, userID, resMsg.UserID)
				// if err != nil {
				// 	log.Printf("failed to get rating: %v", err)
				// 	return
				// }
				// setProfile(userID, resMsg.UserID, myRate, opRate)
				finChan := make(chan struct{})
				dc.OnMessage(onMessage(dc, hch, gch, finChan, board))
				go func() {
					select {
					case <-finChan:
						if err := updateRating(ratingURL, resMsg.RoomID, userID, hash, 2, board.Result()); err != nil {
							log.Printf("failed to update rating: %v", err)
							return
						}
					}
				}()
			})

			if err := conn.Connect(); err != nil {
				log.Fatal("failed to connect Ayame", err)
			}
			log.Printf("conn.Connect();")
			select {
			case <-connected:
				return
			}
		}
	}()
	select {}

	// js.Global().Set("SendGuess", js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
	// 	go func() {
	// 		el := getElementByID("input-number")
	// 		message := el.Get("value").String()
	// 		if message == "" {
	// 			js.Global().Call("alert", "Message must not be empty")
	// 			return
	// 		}
	// 		if dc == nil {
	// 			return
	// 		}

	// 		ch <- entity.NewGuessFromText(message)
	// 		logElem(fmt.Sprintf("[You]: %s\n", message))
	// 		el.Set("value", "")
	// 		for i := 0; i <= 9; i++ {
	// 			s := strconv.Itoa(i)
	// 			getElementByID("input-"+s).Set("disabled", false)
	// 		}
	// 	}()
	// 	return js.Undefined()
	// }))
}

func Send(ch chan *entity.Guess, numbers []int) {
	ch <- entity.NewGuess(numbers)
}

func shortHash(now time.Time) string {
	return sha256Hash(now.String())[:7]
}

func sha256Hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

type Message struct {
	Type   string `json:"type"`
	Turn   *int   `json:"turn,omitempty"`
	Hit    *int   `json:"hit,omitempty"`
	Blow   *int   `json:"blow,omitempty"`
	Guess  string `json:"guess,omitempty"`
	MyHand string `json:"my_hand,omitempty"`
}

func onMessage(dc *webrtc.DataChannel, hch chan *entity.Hand, gch chan *entity.Guess, finChan chan struct{}, board *entity.Board) func(webrtc.DataChannelMessage) {
	return func(msg webrtc.DataChannelMessage) {
		log.Printf("recieve msg.Data: %s", string(msg.Data))
		if !msg.IsString {
			return
		}
		var message Message
		if err := json.Unmarshal(msg.Data, &message); err != nil {
			log.Printf("failed to unmarshal: %v", err)
			return
		}
		// logElem(fmt.Sprintf("[Any]: %s\n", msg.Data))
		switch message.Type {
		case "start":
			// 非開室者Only: GameStart処理
			if board.IsInMenu() {
				log.Printf("message.Turn: %v", *message.Turn)
				initTurn := entity.Turn(*message.Turn).Reverse()
				rand.NewSource(time.Now().UnixNano())
				seed := rand.Int()
				myHand := entity.NewHandBySeed(seed)
				hch <- myHand
				log.Printf("myHand(unopener): %v", myHand)
				// setHand(true, myHand)
				board.Start(myHand, initTurn, 2)
				if board.IsMyTurnInit() {
					log.Printf("YOU FIRST!!!")
				}
			}
			// 非開室者Only: 初回が後攻のときに開室者を初回guess処理に誘導
			if board.IsOpTurn() {
				startMsg := Message{Type: "start"}
				by, _ := json.Marshal(startMsg)
				if err := dc.SendText(string(by)); err != nil {
					log.Printf("failed to send startMsg: %v", err)
					return
				}
				log.Printf("startMsg(unopener): %v", string(by))
				setTurn("It's Opponent's Turn, Waiting ...")
				return
			}
			setTurn("It's Your Turn !")
			// guess送信処理に続く
		case "guess":
			if board.IsMyTurn() {
				return
			}
			// 自分ターンへ遷移
			board.ToggleTurn()
			setTurn("It's Your Turn !")
			guess := entity.NewGuessFromText(message.Guess)
			ans := board.CalcAnswer(guess)
			hit, blow := ans.Hit(), ans.Blow()
			ansMsg := Message{Type: "answer", Hit: &hit, Blow: &blow}
			by, _ := json.Marshal(ansMsg)
			board.CountTurn()
			board.AddOpQA(entity.NewQA(guess, ans))
			setScore(board, guess.View(), hit, blow)
			j := board.Judge()
			setJudge(j)
			log.Printf("ansMsg: %v", string(by))
			if err := dc.SendText(string(by)); err != nil {
				log.Printf("failed to send ansMsg: %v", err)
				return
			}
			if j != entity.NotYet {
				finishProcess(dc, board, finChan)
				return
			}
			// guess送信処理に続く
		case "answer":
			if board.IsMyTurn() {
				return
			}
			ans := entity.NewAnswer(*message.Hit, *message.Blow)
			board.CountTurn()
			board.AddMyQA(entity.NewQA(recentGuess, ans))
			setScore(board, recentGuess.View(), ans.Hit(), ans.Blow())
			j := board.Judge()
			setJudge(j)
			if j != entity.NotYet {
				finishProcess(dc, board, finChan)
				return
			}
			return
		case "timeout":
			setJudge(entity.Win)
			finishProcess(dc, board, finChan)
			return
		case "expose":
			setHand(false, entity.NewHandFromText(message.MyHand))
			return
		default:
			return
		}
		if board.IsOpTurn() {
			return
		}

		// 60sの間にguessを送信する処理
		timeout := 60
		gracePeriod := 1
		toChan := make(chan struct{})
		go func(to int, ch chan struct{}) {
			for {
				select {
				case <-ch:
					log.Printf("catch guess!!!!")
					return
				default:
					to--
					setTimer(to)
					if to <= 0 {
						return
					}
					time.Sleep(1 * time.Second)
				}
			}
		}(timeout, toChan)
		myGuess, isTO := board.WaitGuess(gch, toChan, time.Duration(timeout+gracePeriod)*time.Second)
		recentGuess = myGuess
		if isTO {
			toMsg := Message{Type: "timeout"}
			by, _ := json.Marshal(toMsg)
			if err := dc.SendText(string(by)); err != nil {
				log.Printf("failed to send toMsg: %v", err)
				return
			}
			logElem("[Sys]: You Timeout! You Lose!\n")
			setJudge(entity.Lose)
			finishProcess(dc, board, finChan)
			return
		}
		guessMsg := Message{Type: "guess", Guess: myGuess.Msg()}
		by, _ := json.Marshal(guessMsg)
		// 相手ターンへ遷移
		board.ToggleTurn()
		setTurn("It's Opponent's Turn, Waiting...")
		if err := dc.SendText(string(by)); err != nil {
			log.Printf("failed to send guessMsg: %v", err)
			return
		}
	}
}

func logElem(msg string) {
	log.Printf(msg)
	// el := getElementByID("logs")
	// el.Set("innerHTML", el.Get("innerHTML").String()+msg)
}

// func getElementByID(id string) js.Value {
// 	return js.Global().Get("document").Call("getElementById", id)
// }

func setJudge(judge entity.JudgeStatus) {
	// myJudge := js.Global().Get("document").Call("getElementById", "my-judge")
	// switch judge {
	// case entity.Win:
	// 	myJudge.Set("id", "win")
	// 	myJudge.Set("innerHTML", "WIN")
	// case entity.Lose:
	// 	myJudge.Set("id", "lose")
	// 	myJudge.Set("innerHTML", "LOSE")
	// case entity.Draw:
	// 	myJudge.Set("innerHTML", "DRAW")
	// default:
	// 	return
	// }
}

func setScore(board *entity.Board, guess string, hit int, blow int) {
	// var scores js.Value
	// doc := js.Global().Get("document").Call("getElementsByClassName", "board")
	// scores = doc.Index(0).Call("querySelector", "table").Get("tBodies").Index(0).Get("rows")
	// if board.IsMyTurn() {
	// 	scores = doc.Index(1).Call("querySelector", "table").Get("tBodies").Index(0).Get("rows")
	// }
	// turnCount := board.TurnCount()
	// log.Printf("board.TurnCount() %d", board.TurnCount())
	// guessCell := scores.Index(turnCount).Get("cells").Index(0)
	// hitCell := scores.Index(turnCount).Get("cells").Index(1)
	// blowCell := scores.Index(turnCount).Get("cells").Index(2)
	// guessCell.Set("innerHTML", guess)
	// hitCell.Set("innerHTML", hit)
	// blowCell.Set("innerHTML", blow)
}

func setTimer(second int) {
	// timer := js.Global().Get("document").Call("getElementById", "timer")
	// timer.Set("innerHTML", second)
}

func setTurn(message string) {
	// turnElem := js.Global().Get("document").Call("getElementById", "display-turn")
	// turnElem.Set("innerHTML", message)
}

func setHand(isMyHand bool, hand *entity.Hand) {
	// handID := "op-hand"
	// if isMyHand {
	// 	handID = "my-hand"
	// }
	// for i, number := range *hand {
	// 	handElem := js.Global().Get("document").Call("getElementById", fmt.Sprintf("%s-%d", handID, i+1))
	// 	handElem.Set("innerHTML", number)
	// }
}

func setProfile(myID, opID string, myRate, opRate int) {
	// myProfile := js.Global().Get("document").Call("getElementById", "my-profile")
	// opProfile := js.Global().Get("document").Call("getElementById", "op-profile")
	// myProfile.Set("innerHTML", fmt.Sprintf("%s(r%d)", myID, myRate))
	// opProfile.Set("innerHTML", fmt.Sprintf("%s(r%d)", opID, opRate))
}

func finishProcess(dc *webrtc.DataChannel, board *entity.Board, finChan chan struct{}) {
	setTurn("Finish !!!")
	exposeMsg := Message{Type: "expose", MyHand: board.MyHandText()}
	by, _ := json.Marshal(exposeMsg)
	if err := dc.SendText(string(by)); err != nil {
		log.Printf("failed to send exposeMsg: %v", err)
		return
	}
	finChan <- struct{}{}
	board.Finish()
}

func getRating(ratingURL url.URL, myID, opID string) (int, int, error) {
	ratingURL.Path = path.Join(ratingURL.Path, "/start")
	q := ratingURL.Query()
	q.Set("p1", myID)
	q.Set("p2", opID)
	ratingURL.RawQuery = q.Encode()
	res, err := http.Get(ratingURL.String())
	if err != nil {
		return -1, -1, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return -1, -1, fmt.Errorf("failed to get rating: %v", res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return -1, -1, err
	}
	var resMsg getRatingResMsg
	if err := json.Unmarshal(body, &resMsg); err != nil {
		return -1, -1, err
	}
	return resMsg.Player1.Rate, resMsg.Player2.Rate, nil
}

func updateRating(ratingURL url.URL, roomID, myID, hash string, pNum int, result string) error {
	resMsg := updateRatingResMsg{
		MatchID:  roomID,
		PlayerID: myID,
		Number:   pNum,
		Hash:     hash,
		Result:   result,
	}
	ratingURL.Path = path.Join(ratingURL.Path, "/finish")
	body, err := json.Marshal(resMsg)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(body)
	res, err := http.Post(ratingURL.String(), "application/json", buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update rating: %v", res.Status)
	}
	return nil
}
