package main

import "fmt"

func evenOrOdd() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, number := range numbers {
		if number%2 != 0 {
			fmt.Printf("The number '%v' at position %v is odd\n", number, i)
		} else {
			fmt.Printf("The number '%v' at position %v is even\n", number, i)
		}
	}
}