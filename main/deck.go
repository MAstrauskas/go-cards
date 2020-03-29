package main

import "fmt"

// Deck borrows all the behaviours
// of slice of string
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}