package main

import "testing"

// Writing useful tests/deciding what to test
// testing newDeck
// 1. Make sure a deck is created with x number of cards
// 2. Make sure the first card is an Ace of Spades
// 3 Make sure the last card is a Four of CLubs

// Conditions
// 1. total number of cards expected => suits * values

// testing saveToFile() & newDeckFromFile()

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v: ", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v: ", len(d))
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v: ", len(d))
	}
}
