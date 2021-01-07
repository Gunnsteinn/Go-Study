// |****************************************|
// |				CHANNELS				|
// |****************************************|
// Ejemplo de un canal que no tiene bufer y necesita si o si ser consumido por otra gorutina
// Es decir que se necesita al menos 2 gorutimas para el uso de un canal sin bufer, una que reciba y una
// que lea.
package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	go func() { ch <- 23 }()

	fmt.Println(<-ch)
}
