package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "2 of Spades")

	cards.printDeck()
}

func newCard() string {
	return "5 of Diamonds"
}
