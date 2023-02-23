/* Pruebas unitarias en GO
Para hacer test normalmente vamos a terminar el nombre del script con  "_test.go"
En este caso hariamos deck_test.go
Para ejecutar todas las pruebas en un pacakge, ejecutamos "go test"
*/
package main

import "testing"
import "os"

/* Recordar que t es la variable de tipo "*testing.T" esto segun se explicara más adelante
En la mayoria de lenguajes tenemos una prueba unitaria por cada metodo. S
in embargo en Go podemos tener multiples tests por cada metodo. Sin embargo, el tradeof es que
solo sera PASS OR FAIL (O todos pasan o si alguno falla ya marcara el FAIL)

*/
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %d", len(d))
	}

	/* Probando que la primer carta sea "Ace of Spades" */
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
