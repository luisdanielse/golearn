package main

import "fmt"

func main() {
	// Forma de declarar arreglos
	cards := []string{"first card", "Second card", newCard()}

	/* Como hacer un push al arreglo
	Observar que el push como tal no existe, sino lo que hace es crear un
	nuevo arreglo a partir del existente le hace el push del nuevo y
	eso es el valor de retorno
	*/
	cards = append(cards, "Six of spades")
	//fmt.Println(cards)

	// Iteracion del arreglo
	/*  Importante notar que en
	la definición del ciclo estamos creando las variables "i" y "card"
	Si no usaramos la variable "i" o alguna otra variable, nos generaría un error al momento
	de querer ejecutar (no se trata solamente de un warning)
	*/
	for i, card := range cards {
		fmt.Println(i, card)
		// La siguiente linea no jalaria debido a que la "i" no se esta usando
		//fmt.Println(card)
	}
}

// Cuando declaramos funciones debemos especificar el tipo de retorno que
// la función tendra
func newCard() string {
	return "Five of diamonds"
}
