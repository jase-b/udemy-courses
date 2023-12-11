package main

func main() {
	cards := newDeck()
	cards.saveToFile("52s")
	// hand, remainingDeck := deal(cards, 5)
	// hand.printDeck()
	// remainingDeck.printDeck()
}
