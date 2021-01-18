// |****************************************|
// |				EJERCICIO 1				|
// |****************************************|
// _Lazar 2 gorutinas ademas de la principal.
// _Cada Gorutina debe imprimir algo por pantalla.
// _Usar WaitGroup para asegurarte que cada gorutina finalice antes de que el programa termine.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Número de CPUs al inicio: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas al inicio: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la primera Gorutina.")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hola desde la segunda Gorutina.")
		wg.Done()
	}()

	fmt.Printf("Número de CPUs en el medio: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas en el medio: %v\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("A punto de finalizar main.")
	fmt.Printf("Número de CPUs al final: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas al final: %v\n", runtime.NumGoroutine())

}
