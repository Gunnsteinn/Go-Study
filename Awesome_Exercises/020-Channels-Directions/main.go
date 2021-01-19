// |****************************************|
// |				CHANNELS				|
// |****************************************|
// Utilizando el canal bidireccional c es posible poder generar funciones pasandole la direccionalidad sea send o
// receive.
package main

import "fmt"

func main() {
	c := make(chan int)

	// enviar
	go send(c)

	// recibir
	receive(c)

	fmt.Println("-- THE END --")
}

func send(c chan<- int) {
	c <- 23
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}