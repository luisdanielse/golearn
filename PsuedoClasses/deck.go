package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Con la siguiente sintaxis estamos especificando que
la variable deck "heredara" todo lo que tenemos dentro de un
slice de tipo string
*/
type deck []string

/*
	La siguiente funcion va a pasar a ser propiedad

del tipo "deck"
La variable "d" podria tomarse similarmente a cuando decimos "this" en js o self en python
Normalmente por convencion se usa la primer letra del tipo (en este caso Deck 'd' )
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
	La siguiente funcion regresara un deck como retorno

En esta funcion no necesitamos un receiver dado que  ES LO QUE QUEREMOS REGRESAR (Parecido
a un constructor)
*/
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	/* Con el caracter "_" estamos indicando que "sabemos que existe una variable pero
	no la estamos usando" */
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/*
Recibe como parametro un deck y un entero
(Recordar que en argumentos la segunda palabra es la que define el "tipo".
En este caso los tipos son "deck" e "int")

En esta funcion el segundo parentesis indica que vamos a regresar 2 variables de tipo deck
*/
func deal(d deck, handSize int) (deck, deck) {
	/* El primer retorno sera desde el comienzo hasta el tamano del hand size
	   El segundo retorno sera desde el tamano del hand size hasta el final
	   Practicamente estamos dividiendo el deck en dos partes (esas dos partes son las que
	estamos retornando)

	Recordar que en la sintaxis el ultimo numero NO ESTA INCLUIDO. Ejemplo si handSize = 5,
	la rebanada contendra 0,1,2,3,4
	Y el remaining contendra 5,6,7,8, HASTA EL FINAL
	*/
	return d[:handSize], d[handSize:]
}

/*
type conversion en Go se refiere a pasar de un tipo a otro. Esto viene en la leccion
25 "Deck to String"

En este caso la sintaxis es "Type we want" "Value we already have"
Tenemos un string y lo queremos convertir a un slice de byte
[]byte ("Hi there!")


En este caso para poder convetir el tipo "deck" a un slice de byte
debemos de seguir el siguiente orden:
deck -> []string -> string -> []byte
*/

/* La siguiente sera una funcion "auxiliar" que nos permita convertir un deck a un byte
Rercordar que debemos seguir el siguiente flujo:
deck -> []string -> string -> []byte
*/

func (d deck) toString() string {
	/* Hacemos type conversion de un deck hacia un slice de string */
	//[]string(d)

	/* Una vez que tenemos el arreglo de strings queremos pasar a un solo string
	(concatenando todos los valores) Ejemplo:
	["red", "yellow", "blue"]  a  ["red, yellow, blue"] */

	/* Hay un package llamada "strings"  que nos permite hacer la concatenacion
	EL metodo es "Join" el cual pertenece al package "strings"
	Al contrario al metodo "join" tenemos el metodo split que nos permite
	aislar cada componente, esto se utiliza en el metodo que nos permite crear un
	deck a partir de un archivo
	*/
	return strings.Join([]string(d), ",")
}

/*
Funcion que nos permite guardar en un archivo el contenido de deck y regresa
un error (en caso de que exista)
*/
func (d deck) saveToFile(filename string) error {
	/* La funcion WriteFile nos permite convertir un deck a una sola string (concatenando
	todo en un solo string) sin embargo el metodo "WriteFile" necesita un slice de bytes
	como parametro

	Es por eso que se hace "doble type conversion"
	El 0666 indica que cualquier entidad puede escribir o leer el archivo */
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

/*
	Al igual que el constructor (newDeck) no tenemos ningún receiver dado que

lo que queremos es precisamente construir un deck
*/
func newDeckFromFile(filename string) deck {
	ioutil.ReadFile(filename)
	/* bs contendra el slice de bytes
	err contendra informacion referente al error al momento de lectura (en caso de que se presente)
	Si todo sale "bien" la variable err contendrá un null (nil)
	*/
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	/*
		Type conversion a tipo string

		Parecido al metodo "join" aqui tenemos el metodo "split" en el que dado una string
		y un separador, nos va a aislar cada componentes

		Los parametros que recibe son:
		1.- la string (Es por eso que estamos haciendo un type conversion)
		2.- El separador (En este caso es la , )
	*/
	s := strings.Split(string(bs), ",")
	return deck(s)
}

/*
Para hacer el shuffle, implementaremos un for en todo el deck
Por cada elemento (o iteracion)
generaremos un numero aleatorio entre 0 y el len-1 de forma que esto nos arroje un indice
Una vez que tenemos el indice, haremos un "swap" entre la currenPosition y el indice obtenido
*/
func (d deck) shuffle() {

	/*
		Con la siguiente linea personalizamos el "seed" con la cual
		se genera el numero aleatorio (ya que la funcion que Go ofrece "rand.Intn" es un poco limitada a
			los "numeros aleatorios" que genera
			Mas detalle puede encontrarse en leccion 31  */
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		//newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1)
		/* El swap se hace para ambos elementos INLINE */
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
