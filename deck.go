package main

import "fmt"

// new type of deck
// it is a slice of string
type deck []string

func (d deck) print() { // (d deck) is a reciver
	for i, card := range d {
		fmt.Println(i, card)
	}
}
