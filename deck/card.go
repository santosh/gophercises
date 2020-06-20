//go:generate stringer -type=Suit,Rank

// Package deck provides API for a representation of cards. A Card has a Suit and a Rank.
package deck

import "fmt"

// Suit represents card's category.
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	// Joker is a special case and has no Suit, but we'll treat it as Suit.
	Joker
)

// To compensate for Joker, suits pools only real Suits
var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank represents a single value of card between A, 2, ..., Q, K
type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card belongs to a Suit and has a Rank.
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New creates new deck of cards with a total of
// 13 (ranks) * 4 (suits). Joker is not included.
func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	return cards
}
