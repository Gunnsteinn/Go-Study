// |****************************************|
// |				EJERCICIO 3				|
// |****************************************|
// Crea un struct “errorPer” el cual implemente la interfaz builtin.error. Crea una función “foo” que
// tenga un valor de tipo error como parámetro. Crea un valor de tipo “errorPer” y pásalo a “foo”.
// La funcion foo muestra dos ejemplos de impresion donde podemos ver en el print que nos muestra toda la info del
// metodo Error(), pero el segundo print a traves de assertion podemos indexar la informacion que contiene info ya que
// este metodo implementa errorPer.
package main

import "fmt"

type errorPer struct {
	info string
}

func (e errorPer) Error() string {
	return fmt.Sprintf("[ERROR] - %v", e.info)
}

func main() {
	foo(errorPer{info: "Error by struct. "})
}

func foo(err error) {

	fmt.Println(err)

	//assertion
	fmt.Println(err.(errorPer).info)
}
