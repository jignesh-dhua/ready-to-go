package main

func main() {

	cards := newDeck()
	cards.shuffle()
	// cards.saveToFile("my_cards.txt")

	//cards := newDeckFromFile("my_cards.txt")
	cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
}
