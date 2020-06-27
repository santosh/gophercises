package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Heart})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Hearts
	// Jack of Clubs
	// Queen of Diamonds
	// King of Spades
	// Joker
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

func TestDefaultSort(t *testing.T) {
	t.Run("test if first card is Ace of Spades", func(t *testing.T) {
		cards := New(DefaultSort)
		want := Card{Rank: Ace, Suit: Spade}
		got := cards[0]

		if got != want {
			t.Error("Expected Ace of Spades as first card. Got:", got)
		}
	})

	t.Run("test if last card is King of Hearts", func(t *testing.T) {
		cards := New(DefaultSort)
		want := Card{Rank: King, Suit: Heart}
		got := cards[len(cards)-1]

		if got != want {
			t.Error("Expected King of Hearts as last card. Got:", got)
		}
	})

}

func TestSort(t *testing.T) {
	t.Run("test if first card is Ace of Spades", func(t *testing.T) {
		cards := New(Sort(Less))
		want := Card{Rank: Ace, Suit: Spade}
		got := cards[0]

		if got != want {
			t.Error("Expected Ace of Spades as first card. Got:", got)
		}
	})

	t.Run("test if last card is King of Hearts", func(t *testing.T) {
		cards := New(Sort(Less))
		want := Card{Rank: King, Suit: Heart}
		got := cards[len(cards)-1]

		if got != want {
			t.Error("Expected King of Hearts as last card. Got:", got)
		}
	})

}

func TestShuffle(t *testing.T) {
	cards := New(Sort(Less))
	shuffledCard := Shuffle(cards)

	if reflect.DeepEqual(cards, shuffledCard) {
		t.Error("Shuffled cards are same as unshuffled cards.")
	}
}

func TestJokers(t *testing.T) {
	noWant := 3
	cards := New(Jokers(noWant))

	got := 0
	for _, c := range cards {
		if c.Suit == Joker {
			got++
		}
	}
	if got != noWant {
		t.Errorf("Expected %d Jokers, got %d", noWant, got)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == 2 || card.Rank == 3
	}

	cards := New(Filter(filter))

	for _, c := range cards {
		if c.Rank == 2 || c.Rank == 3 {
			t.Error("Expected all 2s & 3s to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, received %d cards.", 13*4*3, len(cards))
	}
}
