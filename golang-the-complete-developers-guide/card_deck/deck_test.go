package main

import (
	"os"
	"testing"
)

const expectedDeckLength int = 52

func TestNewDeck(t *testing.T) {
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "Jack of Clubs"
	deck := newDeck()

	if len(deck) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v.", expectedDeckLength, len(deck))
	}

	if deck[0] != expectedFirstCard {
		t.Errorf("Expected first card of '%v', but got '%v'.", expectedFirstCard, deck[0])
	}

	if deck[len(deck)-1] != expectedLastCard {
		t.Errorf("Expected last card of '%v', but got '%v'.", expectedLastCard, deck[len(deck)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"

	os.Remove(testFileName)

	deck := newDeck()
	deck.saveToFile(testFileName)

	loadedDeck := newDeckFromeFile(testFileName)

	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v.", expectedDeckLength, len(loadedDeck))
	}

	os.Remove(testFileName)
}
