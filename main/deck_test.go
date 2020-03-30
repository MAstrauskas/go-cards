package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected deck length of 40, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 1st card of 'Ace of Spades', but got '%v'", d[0])
	}

	if d[len(d) - 1] != "Ten of Clubs" {
		t.Errorf("Expected the last card of 'Ten of Clubs', but got '%v'", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	_ = os.Remove("_decktesting")

	deck := newDeck()
	_ = deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Expected 40 cards in deck, but got %v", len(deck))
	}

	_ = os.Remove("_decktesting")
}