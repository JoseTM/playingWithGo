package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// new type deck
// wich is a slice of cards
type deck []card
type card struct {
	suit  string
	value string
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards.addCard(suit, value)
		}
	}

	return cards
}

func newCard(suit string, value string) card {
	c := card{suit: suit, value: value}
	return c
}

func (d deck) addCard(suit string, value string) {
	c := card{suit: suit, value: value}
	d = append(d, c)
}

func newDeckFromFile(filename string) deck {
	fileContent, errorReading := ioutil.ReadFile(filename)
	if errorReading != nil {
		fmt.Println("Error: ", errorReading)
		os.Exit(1)
	}

	stringFromFile := strings.Split(string(fileContent), ",")
	return deck(stringFromFile)
}

//receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}
