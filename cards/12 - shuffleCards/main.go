package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

}

// Notes
// Learnt about explicit type conversion
