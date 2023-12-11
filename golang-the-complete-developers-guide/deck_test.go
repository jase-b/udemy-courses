package main

import "testing"

func TestNewDeck(t *testing.T) {
	expectedDeckLength := 52
	d := newDeck()

	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v.", expectedDeckLength, len(d))
	}
}
