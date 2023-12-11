package main

import "testing"

func TestNewDeck(t *testing.T) {
	expectedDeckLength := 52
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "Jack of Clubs"
	d := newDeck()

	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v.", expectedDeckLength, len(d))
	}

	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card of '%v', but got '%v'.", expectedFirstCard, d[0])
	}

	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card of '%v', but got '%v'.", expectedLastCard, d[len(d)-1])
	}
}
