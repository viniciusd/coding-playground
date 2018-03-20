package main

import "fmt"

func main() {
	cards, i := []string{}, 4

	fmt.Println(i)

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
