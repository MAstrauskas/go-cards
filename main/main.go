package main

func main() {
	cards := deck{newCard(), newCard()}

	// Returns a new slices & does not modify the existing cards slice
	cards = append(cards, "Ten of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

