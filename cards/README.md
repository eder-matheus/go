package main

import "fmt"

func main() {
	// creating new variable
	// var card string = newCard()

	// alternative way to declare and assign a variable
	// card := "Five of Diamonds"

	// declaring a list of strings
	var cards []string = []string{newCard(), newCard()}
	// alternative way to delcare a list of strings
	// cards := []string{"Ace of Spades", newCard()}

	// add new element to cards
	cards = append(cards, "Six of Spades")

	// for loop
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// new function declaration
// func functionName() return_type {...}
func newCard() string {
	return "Five of Diamonds"
}

// Array in go is a list of fixed length
// Slice is a list that can grow or shrink
// Arrays and Slices should have the same data type on each element
// Similar to C++ vector
