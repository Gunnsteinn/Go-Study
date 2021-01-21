/*Package trigonometric provides geometric calculations.

-------------------------------[go tool cover -h]--------------------------------

Given a coverage profile produced by 'go test':
        go test -coverprofile=c.out

Open a web browser displaying annotated source code:


Write out an HTML file instead of launching a web browser:
        go tool cover -html=c.out -o coverage.html

Display coverage percentages to stdout for each function:
        go tool cover -func=c.out

Finally, to generate modified source code with coverage annotations
(what go test -cover does):
        go tool cover -mode=set -var=CoverageVariableName program.go


----------------------------------godoc-----------------------------------

Godoc extracts and generates documentation for Go programs.

It runs as a web server and presents the documentation as a web page.
 godoc -http : 8081

----------------------------------Gofmt-----------------------------------

	Gofmt reformats Go source code.

----------------------------------Govet-----------------------------------

	Govet is concerned with correctness.

----------------------------------Golint-----------------------------------

	Golint is concerned with coding style.
	Golint prints out style mistakes.
	Golint is in use at Google, and it seeks to match the accepted style of the open source Go project.
*/
package trigonometric

import "math"

// TriangleArea calculate the area of a triangle
func TriangleArea(base, height float64) float64 {
	return (base * height) / 2
}

// SquareArea calculate the area of a Square
func SquareArea(side float64) float64 {
	return side * side
}

// RectangleArea calculate the area of a Rectangle
func RectangleArea(base, height float64) float64 {
	return base * height
}

// DiamondArea calculate the area of a Diamond
func DiamondArea(majorDiagonal, minorDiagonal float64) float64 {
	return (majorDiagonal * minorDiagonal) / 2
}

// TrapezeArea calculate the area of a Trapeze
func TrapezeArea(majorBase, minorBase, height float64) float64 {
	return ((majorBase + minorBase) * height) / 2
}

// ParallelogramArea calculate the area of a Parallelogram
func ParallelogramArea(base, height float64) float64 {
	return base * height
}

// PentagonArea calculate the area of a Pentagon
func PentagonArea(perimeter, apothem float64) float64 {
	return (perimeter * apothem) / 2
}

// HexagonArea calculate the area of a Hexagon
func HexagonArea(perimeter, apothem float64) float64 {
	return (perimeter * apothem) / 2
}

// CircleArea calculate the area of a Circle
func CircleArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}
