package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	cards.readFromFile("my_cards")
	cards.shuffle()
	cards.print()
	// hand, remainingDeck := cards.deal(3)
	// fmt.Println(hand)
	// fmt.Println("-----------")
	// fmt.Println(remainingDeck)

}
