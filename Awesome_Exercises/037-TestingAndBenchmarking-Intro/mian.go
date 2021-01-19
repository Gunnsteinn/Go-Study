// |****************************************|
// |				Testing                 |
// |****************************************|
package main

import "fmt"

func main() {
	fmt.Println("Sum: ", Sum(2,3,4))
	fmt.Println("Mul: ", Mul(2,3))
	fmt.Println("Div: ", Div(8,4))
}

// Sum add several numbers.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// Mul multiply two numbers.
func Mul(x, y float64) float64 {
	return x * y
}

// Div divide two numbers.
func Div(x, y float64) float64 {
	return x / y
}
