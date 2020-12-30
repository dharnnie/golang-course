package main

func main() {
	cards := newDeck()
	// hand, remainingDeck := deal(cards, 5) // this is how you get the multiples values from deal(cards, 5). hand & remaindingDeck is being initialized to the return values of deck
	// hand.print()
	// remainingDeck.print()
	// fmt.Println("this simply converts a (Hi there) to a slice of bytes")
	// gretting := "Hi there"
	// fmt.Println([]byte(gretting))
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

}

// Notes
// Learnt about explicit type conversion
