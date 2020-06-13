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
	Joker
)

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
