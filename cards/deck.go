package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of 'deck',
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deckFromString(serialized string) deck {
	return deck(strings.Split(serialized, ","))
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(path string) error {
	err := ioutil.WriteFile(path, []byte(d.toString()), 0666)

	return err
}

func (d deck) shuffle() deck {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}

func newDeckFromFile(path string) (deck, error) {
	bs, err := ioutil.ReadFile(path)
	return deckFromString(string(bs)), err
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
