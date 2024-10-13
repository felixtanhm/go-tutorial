package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	cards.readFromFile("my_cards")
	cards.shuffle()
}
