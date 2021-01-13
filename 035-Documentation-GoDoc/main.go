/*----------------------------------godoc-----------------------------------

Corre como un servidor web y presenta la documentación como una página web
 godoc -http : 8081
--------------------------------------------------------------------------
 Package main solve some math operations. */
package main

import (
	"awesomeProject/035-Documentation-GoDoc/mathme"
	"fmt"
)

// main main program.
func main() {
	fmt.Println("2 + 3 =", mathme.Sum(2, 3, 4))
	fmt.Println("4 + 7 =", mathme.Mul(4, 7))
	fmt.Println("5 + 9 =", mathme.Div(5, 9))
}
