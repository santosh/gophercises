//go:generate stringer -type=Suit,Rank

// Package deck provides API for a representation of cards. A Card has a Suit and a Rank.
package deck

import (
	"fmt"
	"sort"
)

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
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	// This is completely different from how we would have done it in Python.
	// Here we are taking input, and taking options, and passing the
	// input to all the functions which takes cards
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// DefaultSort sorts in a manner like:
// Spade, Ace - King
// Diamond, Ace - King
// Club, Ace - King
// Heart, Ace - King
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Less matches the the signature for https://golang.org/pkg/sort/#Slice
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// absRank returns absolute rank of a card.
func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
