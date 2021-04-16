package main

import "testing"

func TestNewDeck(t *testing.T) {
	lengthExpeted := 16
	cards := newDeck()

	if len(cards) != lengthExpeted {
		t.Fatalf(`len(cards) = %d, want be length = %v 0`, len(cards), lengthExpeted)
	}
}

func TestNewDeckFromFile(t *testing.T) {
	lengthExpeted := 16
	cards := newDeckFromFile("deckTest.dat")

	if len(cards) != lengthExpeted {
		t.Fatalf(`len(cards) = %d, want be length = %v 0`, len(cards), lengthExpeted)
	}
}

func TestFirstElementOfDeck(t *testing.T) {
	firstCardExpected := "Ace of Spades"
	cards := newDeck()
	if cards[0] != firstCardExpected {
		t.Fatalf(`cards[0] = %q, want be %q  0`, cards[0], firstCardExpected)
	}
}

func TestLastElementOfDeck(t *testing.T) {
	lastCardExpected := "Four of Clubs"
	cards := newDeck()
	if cards[len(cards)-1] != lastCardExpected {
		t.Fatalf(`cards[%v] = %q, want be %q  0`, len(cards)-1, cards[len(cards)-1], lastCardExpected)
	}
}
