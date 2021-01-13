// |****************************************|
// |				Errors                  |
// |****************************************|
// El paquete log implementa un paquete simple de logging ... escribe a standard error e imprime la fecha y hora de cada
// mensaje logueado...
// log.Println llama a Output a imprimir a el logger estándar. Los argumentos son manejados de la misma manera que en
// fmt.Println.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("C:/Users/u631170/go/src/awesomeProject/032-Error-Printing&Logging/2-log.Println/1-WithLog/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Println("un error ocurrió", err)
	}
	defer f2.Close()

	fmt.Println("Revisa el archivo log.txt en el directorio.")

}
