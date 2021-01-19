// |****************************************|
// |				Errors                  |
// |****************************************|
// errors.New nos da la posibilidad de poder customizar el error de salida.
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := openFile("sin-archivo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func openFile(dir string) (*os.File, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: %v", err))
	}
	return f, nil
}
