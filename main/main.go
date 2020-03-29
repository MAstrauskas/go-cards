package main

import "fmt"

func main() {
	// Returns a new slices & does not modify the existing cards slice
	cards := newDeck()

	// 1st - current hand 0 to 5 (5 not included)
	// 2nd - handSize to the end (5 to end)
	currentHand, remainingDeck := deal(cards, 5)

    // cards.print()
	fmt.Println("------------------------")
    currentHand.print()
	fmt.Println("------------------------")
	remainingDeck.print()
	fmt.Println("------------------------")
}

