package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}

	// Returns a new slices & does not modify the existing cards slice
	cards = append(cards, "Ten of Spades")

	// i - index,
	// card - current value of index i,
	// range from 0 < cards
	// NOTE: Every time we iterate through the array,
	// i & cards are deleted and initialised again
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

/**
Go has 2 data structures to handle lists:
 1. Array
   - fixed length list of things
 2. Slice
   - an array that can grow or shrink
   - all of the elements must be of the same type
 */