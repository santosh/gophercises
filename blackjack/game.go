package blackjack

import (
	"fmt"

	"github.com/santosh/gophercises/deck"
)

type state int8

// Options is the main option struct for blackjack.New()
type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

// New returns a new Game object
func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}

	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 100
	}
	if opts.BlackjackPayout == 0.0 {
		opts.BlackjackPayout = 1.5
	}
	g.nDecks = opts.Decks
	g.nHands = opts.Hands
	g.blackjackPayout = opts.BlackjackPayout
	return g
}

// Game object is the core of the application.
// See methods.
type Game struct {
	// unexported fields
	nDecks          int
	nHands          int
	blackjackPayout float64

	state state
	deck  []deck.Card

	player    []deck.Card
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI AI
}

// CurrentHand returns player's hand
func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("It isn't currently any player's turn")
	}
}

func bet(g *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	g.playerBet = bet
}

// deal deals Hand to the Player and Dealer
func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

// Play is the main event loop of the game.
func (g *Game) Play(ai AI) int {
	g.deck = nil
	min := 52 * g.nDecks / 3

	// create 3 deck of cards and shuffle it
	for i := 0; i < g.nHands; i++ {
		shuffled := false
		// reshuffle the cards, or its easy to make a guess
		// when cards are low
		if len(g.deck) < min {
			g.deck = deck.New(deck.Deck(g.nDecks), deck.Shuffle)
			shuffled = true
		}
		bet(g, ai, shuffled)
		shuffled = false

		deal(g)

		// if its player turn
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}

		// if its dealer turn
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}
		endHand(g, ai)
	}
	return g.balance
}

//
type Move func(*Game)

// MoveHit draws a card from the deck and append it to hand.
func MoveHit(g *Game) {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		MoveStand(g)
	}
}

// MoveStand moves the turn to next player.
func MoveStand(g *Game) {
	g.state++
}

// draw fetches a card from deck of cards and returns fetched
// card and the remaining deck
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// endHand prints final score with appripriate text
func endHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)
	winnings := g.playerBet
	switch {
	case pScore > 21:
		fmt.Println("You busted!")
		winnings *= -1
	case dScore > 21:
		fmt.Println("Dealer busted!")
		g.balance++
	case pScore > dScore:
		fmt.Println("You win!")
		g.balance++
	case dScore > pScore:
		fmt.Println("You lose!")
		winnings *= -1
	case pScore == dScore:
		winnings = 0
	}
	g.balance += winnings
	fmt.Println()
	ai.Summary([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
}

// Score returns max score, in contrast to MinScore func.
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		// we can't increase 12 to 22, because that will be a bust
		return minScore
	}

	for _, c := range hand {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, and we are changing it to be worth 11
			// 11 - 1 = 10
			return minScore + 10
		}
	}
	return minScore
}

// Soft tells if A is counted as 11
func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

// MinScore takes A as 1, not 11.
func minScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		// because J, Q, K has rank 11, 12, 13..
		// we'll either add 10 or less than 10
		score += min(int(c.Rank), 10)
	}
	return score
}

// min is a helper function to MinScore
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
