package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	if len(testDeck) != 52 {
		t.Errorf("Expected deck length of 52 but got %v.", len(testDeck))
	}

	if testDeck[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card to be Ace of Diamonds but got %v.", testDeck[0])
	}

	if testDeck[len(testDeck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs but got %v.", testDeck[len(testDeck)-1])
	}
}

func TestSaveAndLoadDeckFile(t *testing.T) {
	os.Remove("_testing_deck")

	testDeck := newDeck()
	testDeck.saveToFile("_testing_deck")
	loadedDeck := testDeck.readFromFile("_testing_deck")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck to be of length 52 but got %v.", len(loadedDeck))
	}

	os.Remove("_testing_deck")
}
