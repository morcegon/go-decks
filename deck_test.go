package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of 16, but got %v", len(d))
	}

	expectedFirstCard := "Ace of Spades"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first deck card be %v but got %v", expectedFirstCard, d[0])
	}

	expectedLastCard := "Four of Clubs"
	deckLength := len(d) - 1
	if d[deckLength] != expectedLastCard {
		t.Errorf("Expected last deck card be %v but got %v", expectedLastCard, d[deckLength])
	}
}
