package main

import "fmt"

func main() {
	// Returns a new slices & does not modify the existing cards slice
	cards := newDeck()
	_ = cards.saveToFile("my_cards")

	secondDeckOfCards := newDeckFromFile("my_cards")
	secondDeckOfCards.print()

	fmt.Println("----------------")

	secondDeckOfCards.shuffle()
	secondDeckOfCards.print()
}

