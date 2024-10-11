package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// Creates a new Deck
func newDeck() deck {
	newDeck := []string{}
	var cardSuits = []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	var cardValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}
	return newDeck
}

// Prints out the contents of the Deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deals a hand from a provided Deck based on a provided handSize
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Randomly shuffles a Deck
func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(source)

	for i := range d {
		randNum := newRand.Intn(len(d) - 1)
		d[i], d[randNum] = d[randNum], d[i]
	}

	return d
}

// Converts a slice of type string to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Creates and saves a string to a file in the hardrive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Accesses and reads a file in the hardrive
func (d deck) readFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	stringVar := strings.Split(string(byteSlice), ",")
	fmt.Println(stringVar)
	return stringVar
}
