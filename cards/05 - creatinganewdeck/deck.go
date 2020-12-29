package main

import "fmt"

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

// Syntax (Slices)
// fruits = ["apple","pawpaw","pineapple","apple"]
// fruits[0:2] ====> fruits[startIndexIncludingThisIndex:endAtThisIndexExcludingIt]
// fruits[:2] ====> start from the initial index: end at the 2 excluding 2
// fruits[1:] ====> start from the index 1: end at the end of the slice
// Notes (Good practice)
