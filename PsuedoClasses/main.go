package main

import "fmt"

func main() {
	//cards := deck{"Ace of diamonds", "Another card"}

	cards := newDeck()

	/*
		for i, card := range cards {
			fmt.Println(i, card)
		}
	*/
	//cards.print()
	// fmt.Println("Ahora vamos a hacer el split")
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("Imprimiendo hand")
	// //hand.print()
	// fmt.Println("Imprimiendo remaining deck")
	// //remainingDeck.print()

	// /*A continuacion la funcionalidad de pasar de un deck a una sola string */
	// //fmt.Println(cards.toString())
	// remainingDeck.saveToFile("remainingDeck")

	// fmt.Println("Recuperando el archivo remainingDeck")
	// deckRecuperado := newDeckFromFile("remainingDeck")
	// //deckRecuperado.print()

	/* Ahora hagamos un shuffle de la cartas */
	fmt.Println("Baraja original")
	cards.print()
	fmt.Println("Baraja shuffleada")
	cards.shuffle()
	cards.print()
}
