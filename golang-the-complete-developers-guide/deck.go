package main

import "fmt"

type deck []string

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
