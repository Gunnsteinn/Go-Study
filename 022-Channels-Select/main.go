// |****************************************|
// |				CHANNELS				|
// |****************************************|
// Utliiza el select con la idea de extraer los valores generados de varios canales y sin necesidad de close.
package main

import "fmt"

func main() {
	pair := make(chan int)
	odd := make(chan int)
	exit := make(chan bool)

	// recibir
	go sender(pair, odd, exit)

	// enviar
	receiver(pair, odd, exit)

	fmt.Println("-- THE END --")
}

func sender(p, o chan<- int, exit chan<- bool) {
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			o <- i
		}
	}
	exit <- false
}

func receiver(p, o <-chan int,  exit <-chan bool) {
	for {
		select {
		case v := <-p:
			fmt.Println(v)
		case v := <-o:
			fmt.Println(v)
		case v := <-exit:
			fmt.Println(v)
			return
		}
	}
}