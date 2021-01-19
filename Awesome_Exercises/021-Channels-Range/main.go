// |****************************************|
// |				CHANNELS				|
// |****************************************|
// Se utliza el range para poder extraer los valores del canal pero necesita la condicion del "close" para que cierre el
// mismo y no se produzca un deadlook
package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 1; i < 11; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}

