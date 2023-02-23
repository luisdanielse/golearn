/*
package main
Estamos nombrando a que este script pertenece al package "main" que es lo
mas parecido a un namespace o un modulo.
Package main sera ejecutable

Los comandos usados hasta el momento son:
go run main.go  (Debe ser package main)
go build main.go (Solo compila..Si es package main también creará el exe y automaticamente
lo ejecutara como si hubieramos escrito run)
*/
package main

/* fmt es como el package que contiene las librerias y lo necesario que Go necesita */
import "fmt"

/* Siempre que tenemos un main package, debe existir un
metodo llamado "main" */
func main() {
	fmt.Println("Hi there!")
}
