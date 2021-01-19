// |****************************************|
// |				CHANNELS				|
// |****************************************|
// Para el caso de que se utilice bufer sobre la creacion del canal, no es necesario generar 2 gorutinas(el main y otra)
// Es decir que solo basta con especificar el tamaño del bufer para que se pueda acceder al channel desde la
// gorutina principal.
// Es importante aclarar que si se desea agregar o leer mas elementos de los que propone el bufer, el prgrama lanzara un
// deadlock!
package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 23
	ch <- 19

	fmt.Println(<-ch, <-ch)
}
