package main

import (
	"fmt"

	"github.com/santosh/gophercises/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
