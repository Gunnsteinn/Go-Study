// |****************************************|
// |				Errors                  |
// |****************************************|
// El paquete log implementa un paquete simple de logging ... escribe a standard error e imprime la fecha y hora de cada
// mensaje logueado...
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Println("un error ocurrió", err)
	}
	defer f.Close()
}
