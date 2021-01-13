// |****************************************|
// |				Errors                  |
// |****************************************|
// lee sobre el uso de errors.New en la biblioteca estándar:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go
// En este caso el error se toma desde una creacion de una variable global
package main

import (
	"errors"
	"log"
	"os"
)

var errOpenFile = errors.New("Se produce un error al intentar abrir el archivo.")

func main() {
	_, err := openFile("sin-archivo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func openFile(dir string) (*os.File, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, errOpenFile
	}
	return f, nil
}
