package main

import (
	"fmt"
)

func main() {
	cards := []string{newCard(), "Ace of spades"}
	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
