// |****************************************|
// |				Errors                  |
// |****************************************|
// Panicln es quivalente a Println() seguido por una llamada a panic().
// Fatalln es equivalente a Println() seguido por una llamada a os.Exit(1).
// Aqui a diferencia de Fatalln las funciones diferidas si corren.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("C:/Users/u631170/go/src/awesomeProject/032-Error-Printing&Logging/4_fmt-errorf_var-log.Panicln/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	defer foo("Revisa el archivo log.txt en el directorio.")

	_, err1 := os.Open("sin-archivo.txt")
	if err1 != nil {
		log.Panicln(err1)
	}
}

func foo(txt string) {
	fmt.Println(txt)
}
