// |****************************************|
// |				Errors                  |
// |****************************************|
// Ejemplo en el cual disparamos un error para poder utilizar el recover y poder recuperar el control de una gorutina
// en panico. Importante aclarar que el recover captura el valor que se le pasa a panic y se lo puede utilizar desde
// este mismo.
package main

import "fmt"

func main() {
	bar()
	fmt.Println("Returned normally from f.")
}

func bar() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	foo(3)
	fmt.Println("Returned normally from g.")
}

func foo(num int) {
	if num < 0 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", num))
	}
	defer fmt.Println("Defer in g", num)
	fmt.Println("Printing in g", num)
	foo(num - 1)
}
