package main

import "fmt"

var deckSize int

func main() {
	cards := newDeck()

	cards.print()

	printEmptyLine()

	cards.shuffle()

	cards.print()

}

/*
What is the max value and min value of int in Go ?
*/

func printEmptyLine() {
	fmt.Println("")
}
