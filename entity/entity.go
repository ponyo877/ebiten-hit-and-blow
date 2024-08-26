package entity

import (
	"fmt"
	"log"
	"slices"

	"github.com/mowshon/iterium"
)

type State int

const numOfDigits = 3
const numOfAllHandsPatturn = 720

var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

const (
	InMenu State = iota
	Playing
	Finished
)

type Turn int

const (
	MyTurn Turn = iota
	OpTurn
)

func (t Turn) Reverse() Turn {
	return t ^ 1
}

func (t Turn) IsMyTurn() bool {
	return t == MyTurn
}

func NewTurnBySeed(seed int) Turn {
	return Turn(seed % 2)
}

type Hand []int

var allHands [numOfAllHandsPatturn]Hand

func init() {
	// generate all hands
	permutations := iterium.Permutations(numbers, numOfDigits)
	numbersList, _ := permutations.Slice()
	for i, ns := range numbersList {
		allHands[i] = Hand(ns)
	}
}

func NewHand(numbers []int) *Hand {
	hand := Hand(numbers)
	return &hand
}

func NewHandFromText(numStr string) *Hand {
	var numbers []int
	for _, r := range numStr {
		numbers = append(numbers, int(r-'0'))
	}
	return NewHand(numbers)
}

func NewHandBySeed(seed int) *Hand {
	return &allHands[seed%numOfAllHandsPatturn]
}

func (h *Hand) Msg() string {
	n := []int(*h)
	return fmt.Sprintf("%d%d%d", n[0], n[1], n[2])
}

type Guess Hand

func NewGuess(numbers []int) *Guess {
	guess := Guess(numbers)
	return &guess
}

func NewGuessFromText(numStr string) *Guess {
	var numbers []int
	for _, r := range numStr {
		numbers = append(numbers, int(r-'0'))
	}
	guess := Guess(numbers)
	return &guess
}

func (g *Guess) View() string {
	n := []int(*g)
	return fmt.Sprintf("%d %d %d", n[0], n[1], n[2])
}

func (g *Guess) Msg() string {
	n := []int(*g)
	return fmt.Sprintf("%d%d%d", n[0], n[1], n[2])
}

func (h *Hand) Answer(guess *Guess) *Answer {
	var hit, blow int
	for i, n := range *guess {
		if n == (*h)[i] {
			hit++
		} else if slices.Contains(*h, n) {
			blow++
		}
	}
	return NewAnswer(hit, blow)
}

func (h *Hand) QA(guess *Guess) *QA {
	return NewQA(guess, h.Answer(guess))
}

type Answer struct {
	hit  int
	blow int
}

func NewAnswer(hit, blow int) *Answer {
	return &Answer{hit, blow}
}

func (a *Answer) Hit() int {
	return a.hit
}

func (a *Answer) Blow() int {
	return a.blow
}

func (a *Answer) IsAllHit() bool {
	return a.hit == numOfDigits && a.blow == 0
}

func (a *Answer) Msg() string {
	return fmt.Sprintf("%d hit, %d blow", a.hit, a.blow)
}

type QA struct {
	guess  *Guess
	answer *Answer
}

func NewQA(guess *Guess, answer *Answer) *QA {
	return &QA{guess, answer}
}

func (qa *QA) GuessView() []int {
	return []int(*qa.guess)
}

func (qa *QA) HistoryView() []int {
	return []int{qa.answer.hit, qa.answer.blow}
}

type Board struct {
	state       State
	initTurn    Turn
	turn        Turn
	pNum        int
	myHand      *Hand
	myQA        []*QA
	opQA        []*QA
	myTurnCount int
	opTurnCount int
}

func NewBoard() *Board {
	return &Board{
		state: InMenu,
		myQA:  make([]*QA, 0),
		opQA:  make([]*QA, 0),
	}
}

func (b *Board) IsInMenu() bool {
	return b.state == InMenu
}

func (b *Board) IsPlaying() bool {
	return b.state == Playing
}

func (b *Board) IsMyTurn() bool {
	return b.turn == MyTurn
}

func (b *Board) IsOpTurn() bool {
	return b.turn == OpTurn
}

func (b *Board) IsMyTurnInit() bool {
	return b.initTurn == MyTurn
}

func (b *Board) IsOpTurnInit() bool {
	return b.initTurn == OpTurn
}

func (b *Board) ToggleTurn() {
	b.turn = b.turn.Reverse()
}

func (b *Board) CountTurn() {
	if b.IsMyTurn() {
		b.myTurnCount++
	}
	if b.IsOpTurn() {
		b.opTurnCount++
	}
}

func (b *Board) TurnCount() int {
	if b.IsMyTurn() {
		return b.myTurnCount
	}
	return b.opTurnCount
}

func (b *Board) Start(hand *Hand, initTurn Turn, pNum int) {
	b.state = Playing
	b.initTurn, b.turn = initTurn, initTurn
	b.myHand = hand
	b.pNum = pNum
}

type JudgeStatus int

const (
	NotYet JudgeStatus = iota
	Win
	Lose
	Draw
)
const maxTurnCount = 8

func (b *Board) Judge() JudgeStatus {
	var isMy3hit, isOp3hit bool
	if len(b.myQA) > 0 {
		isMy3hit = b.myQA[len(b.myQA)-1].answer.IsAllHit()
	}
	if len(b.opQA) > 0 {
		isOp3hit = b.opQA[len(b.opQA)-1].answer.IsAllHit()
	}
	// 自分のターン(answer送信) && 自分スタート
	if b.IsMyTurn() && b.IsMyTurnInit() {
		if !isMy3hit && isOp3hit {
			return Lose
		}
		if isMy3hit && !isOp3hit {
			return Win
		}
		if isMy3hit && isOp3hit {
			return Draw
		}
	}
	// 相手のターン(answer受取) && 相手スタート
	if b.IsOpTurn() && b.IsOpTurnInit() {
		if !isOp3hit && isMy3hit {
			return Win
		}
		if isOp3hit && !isMy3hit {
			return Lose
		}
		if isMy3hit && isOp3hit {
			return Draw
		}
	}
	if b.myTurnCount == b.opTurnCount && b.myTurnCount == maxTurnCount {
		return Draw
	}
	return NotYet
}

func (b *Board) Finish() {
	b.state = Finished
}

func (b *Board) CalcAnswer(guess *Guess) *Answer {
	return b.myHand.Answer(guess)
}

func (b *Board) AddMyQA(qa *QA) {
	log.Printf("AddMyQA: %v", qa)
	b.myQA = append(b.myQA, qa)
}

func (b *Board) AddOpQA(qa *QA) {
	log.Printf("AddOpQA: %v", qa)
	b.opQA = append(b.opQA, qa)
}

func (b *Board) WaitGuess(ch chan *Guess, cCh, toCh chan struct{}) (*Guess, bool) {
	select {
	case guess := <-ch:
		close(cCh)
		return guess, false
	case <-toCh:
		return nil, true
	}
}

func (b *Board) MyHandText() string {
	return b.myHand.Msg()
}

func (b *Board) PNum() int {
	return b.PNum()
}

func (b *Board) Result() string {
	if b.Judge() == Draw {
		return "0.5"
	}
	if b.Judge() == Win && b.pNum == 1 || b.Judge() == Lose && b.pNum == 2 {
		return "1"
	}
	return "0"
}

type Rating struct {
	id     string
	rating int
}

func NewRating(id string, rating int) *Rating {
	return &Rating{id, rating}
}

func (r *Rating) ID() string {
	return r.id
}

func (r *Rating) Rating() int {
	return r.rating
}
