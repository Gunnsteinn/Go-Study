// |****************************************|
// |				Errors                  |
// |****************************************|
// leer sobre el uso structs con tipo error en la biblioteca estándard:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type customError struct {
	Date time.Time
	err  error
}

func (c customError) Error() string {
	return fmt.Sprintf("[ERROR][%v] - %v", c.Date, c.err)
}

func main() {
	_, err := openFile("sin-archivo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func openFile(dir string) (*os.File, error) {
	f, err := os.Open(dir)
	if err != nil {
		nme := fmt.Errorf("%v", err)
		return nil, customError{Date: time.Now().Local(), err: nme}
	}
	return f, nil
}
