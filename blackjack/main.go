package main

import (
	"fmt"
	"strings"

	"github.com/santosh/gophercises/deck"
)

// Hand is alias for slice of deck of cards
type Hand []deck.Card

// String from Hand pretty prints the slice of deck of cards
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// DealerString hides the second card
func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

// Score returns max score, in contrast to MinScore func.
func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		// we can't increase 12 to 22, because that will be a bust
		return minScore
	}

	for _, c := range h {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, and we are changing it to be worth 11
			// 11 - 1 = 10
			return minScore + 10
		}
	}
	return minScore
}

// MinScore takes A as 1, not 11.
func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
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

// draw fetches a card from deck of cards and returns fetched
// card and the remaining deck
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// Shuffle shuffles the card from game state and return a new state
func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	// CODE SMELL: Remove magic number
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

// Deal deals Hand to the Player and Dealer
func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(gs)
	}
	return ret
}

// Stand changes the turn from Dealer to Player
func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

// EndHand print final score with appripriate text
func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", ret.Player, "\nScore:", pScore)
	fmt.Println("Dealer:", ret.Dealer, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted!")
	case dScore > 21:
		fmt.Println("Dealer busted!")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose!")
	case pScore == dScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	var gs GameState
	// create 3 deck of cards and shuffle it
	gs = Shuffle(gs)
	gs = Deal(gs)
	var input string
	for gs.State == StatePlayerTurn {
		fmt.Println("Player:", gs.Player)
		fmt.Println("Dealer:", gs.Dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("Invalid option:", input)
		}
	}

	for gs.State == StateDealerTurn {
		// If dealer score is <= 16, dealer hit
		// If dealer has a soft 17, then dealer hit
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

	gs = EndHand(gs)
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

// GameState binds all component together
type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

// CurrentPlayer returns player's hand
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It isn't currently any player's turn")
	}
}

// clone clones a game state and returns it
func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}
