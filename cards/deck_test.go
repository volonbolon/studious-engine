package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v instead", d[0])
	}

	lidx := len(d) - 1
	if d[lidx] != "King of Hearts" {
		t.Errorf("Expected last card to be 'King of Hearts', but got %v instead", d[lidx])
	}
}

func TestSaveToDeskAndReadFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	ld := readFromFile(filename)

	if len(ld) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	os.Remove(filename)
}
