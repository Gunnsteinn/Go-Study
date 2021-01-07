// |****************************************|
// |				EJERCICIO 3				|
// |****************************************|
// Crea un struct “errorPer” el cual implemente la interfaz builtin.error. Crea una función “foo” que
// tenga un valor de tipo error como parámetro. Crea un valor de tipo “errorPer” y pásalo a “foo”.
package main

import "fmt"

type errorPer struct {
	info string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("Error: %v", ep.info)
}

func main() {
	errAux := errorPer{info: "Algo"}

	foo(errAux)

}

func foo(e error) {
	fmt.Println("Foo run - ", e)
}
