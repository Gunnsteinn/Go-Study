// |****************************************|
// |				EJERCICIO 3				|
// |****************************************|
//	● Usando gorutinas, crea un programa incremento
//		○ Haz que una variable guarde el valor del incremento
//		○ Lanza varias gorutinas
//			■ cada gorutina deberá
//				● Leer el valor del incremento
//					○ Almacenarlo en una nueva variable
//				● Ceder el procesador con runtime.Gosched()
//				● Incrementa la nueva variable
//				● Escribe el valor en la nueva variable de vuelta a la variable incremento
//	● Usa waitgroups para esperar que todas las gorutinas finalicen
//	● Lo anterior generará una race condition.
//	● Comprueba que es una race condition usando el flag -race
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t\t", runtime.GOARCH)
	fmt.Println("NumCPU:\t\t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine:\t", runtime.NumGoroutine())

	count := 0
	const gm = 100
	var wg sync.WaitGroup
	wg.Add(gm)

	for i := 0; i < gm; i++ {
		go func() {
			aux := count
			runtime.Gosched()
			aux++
			count = aux
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", count)
}
