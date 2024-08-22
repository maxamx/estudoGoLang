package main

func main() {
	cards := newDeck()

	hand, reminingDeck := deal(cards, 5)

	hand.print()
	reminingDeck.print()

}
