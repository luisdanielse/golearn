package main

import "fmt"

func main() {
	/*
		type conversion en Go se refiere a pasar de un tipo a otro. Esto viene en la leccion
		25 "Deck to String"

		En este caso la sintaxis es "Type we want" "Value we already have"
		Tenemos un string y lo queremos convertir a un slice de byte
		[]byte ("Hi there!")
	*/
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))
}
