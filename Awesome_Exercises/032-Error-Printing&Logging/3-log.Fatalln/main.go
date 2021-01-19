// |****************************************|
// |				Errors                  |
// |****************************************|
// Las funciones de Fatal llaman a os.Exit(1) después de escribir el mensaje del log ...
// Fatalln es quivalente a Println() seguido por una llamada a os.Exit(1). Cuando se llama a esta funcion se cierra el
// programa matando a lo que este activo por esto la funcion foo por mas que este diferida no va a correr.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}
