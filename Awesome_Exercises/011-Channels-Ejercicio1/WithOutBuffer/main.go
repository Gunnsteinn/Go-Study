// |****************************************|
// |				EJERCICIO 1				|
// |****************************************|
//	● Haz que este código funcione:
//		○ Usando una función literal, también conocida como, función anónima autoejecutable
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}
