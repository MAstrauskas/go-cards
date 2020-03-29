package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Deck borrows all the behaviours
// of slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// Returns everything from 0 to handSize (not included)
	// and another one from handSize to the end of the slice
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666 ) // 0666 - anybody can read/write
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// log an error & quit the program
		fmt.Println("Error occurred while reading the file: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") // Convert byte slice to slice of strings

	return deck(s) // Return a deck of strings
}