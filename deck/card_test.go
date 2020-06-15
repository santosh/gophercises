package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Heart})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Spade})

	// Output:
	// Ace of Hearts
	// Two of Hearts
	// Jack of Clubs
	// Queen of Diamonds
	// King of Spades
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 * 4

	t.Run("test if number of cards = 13*4", func(t *testing.T) {
		if len(cards) != 13*4 {
			t.Error("Wrong number of cards in a new deck.")
		}
	})

	t.Run("test if all suits are present", func(t *testing.T) {
		var aces []Card
		for _, card := range cards {
			if card.Rank.String() == "Ace" {
				aces = append(aces, card)
			}
		}

		if len(aces) != 4 {
			t.Error("Wrong number of suits in new deck.")
		}
	})

}
