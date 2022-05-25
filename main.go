package main

func main() {
	// var card string = "new card"
	// card := newCard()
	// fmt.Println(card)

	// cards := deck{newCard(), newCard()}
	// cards = append(cards, "new card")

	// cards := newDeck()
	cards := newDeckFromFile("cards")
	cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// cards.toString()

	// cards.saveToFile("cards")
	// fmt.Println(cards.toString())
}

// func newCard() string {
// 	return "Ace of Spades"
// }
