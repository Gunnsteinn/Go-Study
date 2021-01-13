// |****************************************|
// |				Errors                  |
// |****************************************|
package main

import (
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
		ErrOpen := fmt.Errorf("Error:  %v", err)
		return nil, ErrOpen
	}
	return f, nil
}