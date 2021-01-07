// |****************************************|
// |				EJERCICIO 1				|
// |****************************************|
//	● Haz que este código funcione:
//		○ Usando un canal con búfer
package main

import "fmt"

func main() {
	ch := make(chan int,1)

	ch <- 42

	fmt.Println(<-ch)
}
