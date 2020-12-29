package main

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5) // this is how you get the multiples values from deal(cards, 5). hand & remaindingDeck is being initialized to the return values of deck
	hand.print()
	remainingDeck.print()

}
