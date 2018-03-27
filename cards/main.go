package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	hand.saveToFile("foo")

	cards, _ = newDeckFromFile("foo")

	hand.print()
	hand = hand.shuffle()
	hand.print()
}
