// |****************************************|
// |				EJERCICIO 4				|
// |****************************************|
//	Arregla la race condition que creaste en el ejercicio anterior usando un mutex
//		● Tiene sentido eliminar runtime.Gosched(), porque en este caso si vamos a bloquear la
// gorutine en esa parte del codigo no tiene sentido ceder el procesador.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t\t", runtime.GOARCH)
	fmt.Println("NumCPU:\t\t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine:\t", runtime.NumGoroutine())

	const ct = 100
	counter := 0

	wg.Add(ct)

	for i := 0; i < ct; i++ {
		go func() {
			mu.Lock()
			aux := counter
			aux++
			counter = aux
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("NumGoroutine:\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter:\t", counter)
}
