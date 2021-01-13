// |****************************************|
// |				Errors                  |
// |****************************************|
// http://godoc.org/builtin#panic
// Panic es un error que detendrá la ejecución de nuestro programa a menos que lo manejemos.
// Los errores nos permitían comprobar su existencia y tomar las acciones que consideremos adecuadas según la situación.
// En el caso de producirse un panic el programa se detendrá desde la función donde se originó hacia arriba,
// en el caso de que no sea esta nuestra intención, debemos tomar medidas para manejar el panic.
package main

import (
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		panic(err)
	}
}


