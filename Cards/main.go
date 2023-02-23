package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// Lo siguiente es otra forma de declarar la variable
	card := "Ace of spades"
	fmt.Println(card)
}

// Cuando declaramos funciones debemos especificar el tipo de retorno que
// la función tendra
func newCard() string {
	return "Five of diamonds"
}
