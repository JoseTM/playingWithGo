package main

import "log"

func main() {
	cards := newDeck()

	hand, remainingDeck := cards.deal(5)

	hand.print()
	remainingDeck.print()
	//cards.print()
	errorSaving := cards.saveToFile("deck.dat")
	if errorSaving != nil {
		log.Fatal(errorSaving)
	}
	deckImported := newDeckFromFile("deck.dat")
	//deckImported.print()
	deckImported.shuffle()
	deckImported.print()
}
