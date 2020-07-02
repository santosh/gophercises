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

func main() {
	// create 3 deck of cards and shuffle it
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			// take the first card from cards
			card, cards = draw(cards)
			// append it to one of the hand (either player or dealer)
			*hand = append(*hand, card)
			// by the end of the loop,
			// both player and dealer will have one-one card
		}
		// In outer loop, we just repeat the appending of cards n number of times.
	}

	var input string
	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", player, "\nScore:", player.Score())
	fmt.Println("Dealer:", dealer, "\nScore:", dealer.Score())
}
