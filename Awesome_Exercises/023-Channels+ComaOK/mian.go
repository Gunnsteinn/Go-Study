// |****************************************|
// |				CHANNELS				|
// |****************************************|
// Utilizando el ejercicio de select podemos ver que en el sender solo cerramos el canal del exit y desde el receiver
// utilizamos el Coma OK para poder determinar cuando el canal esta en false a consecuencia de estar cerrado
package main

import (
	"fmt"
)

func main() {
	pair := make(chan int)
	odd := make(chan int)
	exit := make(chan bool)

	// enviar
	go sender(pair, odd, exit)

	// recibir
	receiver(pair, odd, exit)

	fmt.Println("-- THE END--")
}

func sender(p, o chan<- int, e chan<- bool) {
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			o <- i
		}
	}
	close(e)
}

func receiver(p, o <-chan int, e <-chan bool) {
	for {
		select {
		case v, ok := <-p:
			fmt.Println(v, ok)
		case v := <-o:
			fmt.Println(v)
		case t, ok := <-e:
			if !ok {
				fmt.Println(t, ok)
				return
			}
			fmt.Println(t, ok)
		}
	}
}
