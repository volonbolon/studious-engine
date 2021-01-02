package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
type deck []string

func newDeck() deck {
	deck := deck{}

	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Clubs",
		"Hearts",
	}

	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit

			deck = append(deck, card)
		}
	}

	return deck
}

func (d deck) print() { // (d deck) is the receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	hand := d[:handSize]
	return hand, d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	asString := d.toString()
	bytes := []byte(asString)
	err := ioutil.WriteFile(filename, bytes, 0644)
	return err
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func readFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	asString := strings.Split(string(data), ", ")
	deck := deck(asString)

	return deck
}
