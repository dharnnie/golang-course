package main

import (
	"fmt"
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

// Notes (Good practice)
// 1. Only use receivers where necesary e.g when creating a new deck you do not need to use a receiver
