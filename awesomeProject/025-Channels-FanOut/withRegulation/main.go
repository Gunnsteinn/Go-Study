// |****************************************|
// |				CHANNELS				|
// |****************************************|
// La idea de FanOut es generar un trabajo(secuencial) donde se lo pueda dividir en multiples gorutinas para ir
// completando por cada una parte del mismo. En este caso se regula la cantidad de Gorutinas que se lanzan y se puede
// ver en la impresion de la linea 55 que nos deja 13:
// 10 gorutinas del proceso regulado
// 1 gorutina padre del proceso regulado
// 1 gorutina para el que agrega valores
// 1 gorutina del main

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go adder(c1)

	go fanOutIn(c1, c2)

	for value := range c2 {
		fmt.Println(value)
	}

	fmt.Println("--THE END--")
}

func adder(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const maxOfgoRutines = 10

	for i := 0; i < maxOfgoRutines; i++ {
		wg.Add(maxOfgoRutines)
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeComsumerWork(v2)
					wg.Done()
				}(v)
			}
		}()
	}
	fmt.Printf("Go rutines Number: %v\n", runtime.NumGoroutine())
	wg.Wait()
	close(c2)
}

func timeComsumerWork(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}
