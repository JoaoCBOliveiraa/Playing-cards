package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardFigures := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, figure := range cardFigures {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+figure)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
