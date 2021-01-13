// |****************************************|
// |				Errors                  |
// |****************************************|
// Println formatea utilizando los formatos predeterminados para sus operandos y escribe a la salida estándar.
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("sin-archivo.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
