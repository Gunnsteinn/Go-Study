// |****************************************|
// |				EJERCICIO 2				|
// |****************************************|
// Haz que este código funcione:
package main

import (
	"fmt"
)

func main() {
	// Error por declarar un canal de tipo send
	//     cs := make(chan<- int)
	// Error por declarar un canal de tipo recieve
	//     cs := make(<-chan int)
	// Se soluciona convirtiendo el canal de tipo bidireccional
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
