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

func main() {
	// create 3 deck of cards and shuffle it
	cards := deck.New(deck.Deck(3), deck.Shuffle)

	var card deck.Card

	// Take first 10 cards from the shuffled deck
	for i := 0; i < 10; i++ {
		card, cards = cards[0], cards[1:]
		fmt.Println(card)
	}

	var h Hand = cards[0:3]
	fmt.Println(h)
}
