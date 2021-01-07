// |****************************************|
// |				CHANNELS				|
// |****************************************|
package main

import (
	"fmt"
	"sync"
)

func main() {
	pair := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	// enviar
	go sender(pair, odd)

	// recibir
	go receiver(pair, odd, fanIn)

	for i := range fanIn {
		fmt.Println(i)
	}
	fmt.Println("--THE END--")
}

func sender(p, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			o <- i
		}
	}
	close(p)
	close(o)
}

func receiver(p, o <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range p {
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanIn <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)
}
