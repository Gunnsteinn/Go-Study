// |****************************************|
// |				Testing                 |
// |****************************************|
package main

import (

	"fmt"
	"github.com/Go-Study/awesomeProject/035-Documentation-GoDoc/mathme"
)

func main() {
	fmt.Println("Sum: ", mathme.Sum(2, 3, 4))
	fmt.Println("Mul: ", mathme.Mul(2, 3))
	fmt.Println("Div: ", mathme.Div(8, 4))
}
