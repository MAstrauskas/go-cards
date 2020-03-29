package main

func main() {
	// Returns a new slices & does not modify the existing cards slice
	cards := newDeck()
	_ = cards.saveToFile("my_cards")
}

