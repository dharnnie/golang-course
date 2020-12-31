package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// func (d deck) shuffle() {
// 	for i := range d {

// 		newPos := rand.Intn(len(d) - 1)
// 		d[i], d[newPos] = d[newPos], d[i]
// 	}
// }

func (d deck) shuffle() {
	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

// if you test deck.shuffle() it returns the same randomness everytime, so we'll need to add a little twist to this.
// We'll use the
// source := rand.NewSource()
// r := math.Rand(s Source)
// value := r.Intn(someInteger) // call the Intn function of reciever of type r -> Rand

// time
//
