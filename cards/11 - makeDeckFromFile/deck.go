package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := suit + " of " + value
			cards = append(cards, card)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // to understand this, see deck.go in 06
}

func (d deck) toString() string {
	s := []string(d)
	return strings.Join(s, ",")
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1  - log the error and return a call to newDeck()
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s) // we are able to do this because deck is an underlying []string type
}

// Notes (Good practice)
// Make a deck from a file
