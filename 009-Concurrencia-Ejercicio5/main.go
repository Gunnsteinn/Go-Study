// |****************************************|
// |				EJERCICIO 4				|
// |****************************************|
//	Arreglar la race condition que creaste en el ejercicio #4 usando el paquete atomic
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const ONE = 1

var count int64
var wg sync.WaitGroup

func main() {

	const ct = 100
	wg.Add(ct)

	for i := 0; i < ct; i++ {
		go foo()
	}

	wg.Wait()
	fmt.Println("Counter:", count)
}

func foo() {
	atomic.AddInt64(&count, ONE)
	fmt.Println("Count:", atomic.LoadInt64(&count))
	wg.Done()
}
