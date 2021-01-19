// |****************************************|
// |				EJERCICIO 3				|
// |****************************************|
// ● Comenzando con este código , extrae los valores del canal usando un ciclo for range
package main

import (
	"fmt"
)

func main() {
	c := gen()
	recibir(c)

	fmt.Println("A punto de finalizar.")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func recibir(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
