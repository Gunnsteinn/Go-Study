// |****************************************|
// |     Benchmarking and Documentation     |
// |****************************************|
package main

import (
	"fmt"
	"github.com/Gunnsteinn/Go-Study/awesomeProject/040-TestingAndBenchmarking-Benchmark/trigonometricFunctions"
)

func main() {
	fmt.Println(">>> Calculation of basic geometric figures <<<")
	fmt.Println("Triangle: ", trigonometric.TriangleArea(2, 2))
	fmt.Println("Square: ", trigonometric.SquareArea(2))
	fmt.Println("Rectangle: ", trigonometric.RectangleArea(2, 2))
	fmt.Println("Diamond: ", trigonometric.DiamondArea(2, 1))
	fmt.Println("Trapeze: ", trigonometric.TrapezeArea(2, 1, 3))
	fmt.Println("Parallelogram: ", trigonometric.ParallelogramArea(2, 3))
	fmt.Println("Pentagon: ", trigonometric.PentagonArea(2, 3))
	fmt.Println("Hexagon: ", trigonometric.HexagonArea(2, 3))
	fmt.Println("Circle: ", trigonometric.CircleArea(2))
}
