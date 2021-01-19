// Package trigonometric provides math solutions.
package trigonometric

import (
	"fmt"
	"testing"
)

type oneIn struct {
	InA      float64
	Expected float64
}

type twoIn struct {
	InA      float64
	InB      float64
	Expected float64
}

type threeIn struct {
	InA      float64
	InB      float64
	InC      float64
	Expected float64
}

var twoIns1 = []twoIn{
	{InA: 1, InB: 2, Expected: 1.00},
	{InA: 2, InB: 2, Expected: 2.00},
	{InA: 3, InB: 2, Expected: 3.00},
	{InA: 4, InB: 2, Expected: 4.00},
	{InA: 5, InB: 2, Expected: 5.00},
}

var twoIns2 = []twoIn{
	{InA: 1, InB: 2, Expected: 2.00},
	{InA: 2, InB: 2, Expected: 4.00},
	{InA: 3, InB: 2, Expected: 6.00},
	{InA: 4, InB: 2, Expected: 8.00},
	{InA: 5, InB: 2, Expected: 10.00},
}

var threeIns = []threeIn{
	{InA: 1, InB: 5, InC: 5, Expected: 15.00},
	{InA: 2, InB: 5, InC: 4, Expected: 14.00},
	{InA: 3, InB: 5, InC: 3, Expected: 12.00},
	{InA: 4, InB: 5, InC: 2, Expected: 9.00},
	{InA: 5, InB: 5, InC: 1, Expected: 5.00},
}

func TestTriangleArea(t *testing.T) {
	for _, value := range twoIns1 {
		got := TriangleArea(value.InA, value.InB)
		if got != value.Expected {
			t.Errorf("got %.2f expected %.2f", got, value.Expected)
		}
	}
}

func ExampleTriangleArea() {
	// {InA: 1, InB: 2}
	// {InA: 2, InB: 2}
	// {InA: 3, InB: 2}
	// {InA: 4, InB: 2}
	// {InA: 5, InB: 2}
	for _, value := range twoIns1 {
		fmt.Println(TriangleArea(value.InA, value.InB))
	}
	//Output: 
	// 1
	// 2
	// 3
	// 4
	// 5
}

func BenchmarkTriangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range twoIns1 {
			TriangleArea(value.InA, value.InB)
		}
	}
}

func TestSquareArea(t *testing.T) {

	oneIns := []oneIn{
		{InA: 1, Expected: 1},
		{InA: 2, Expected: 4},
		{InA: 3, Expected: 9},
		{InA: 4, Expected: 16},
		{InA: 5, Expected: 25},
	}

	for _, value := range oneIns {
		got := SquareArea(value.InA)
		if got != value.Expected {
			t.Errorf("got %.2f expected %.2f", got, value.Expected)
		}
	}
}

func ExampleSquareArea() {
	for _, value := range []float64{1, 2, 3, 4, 5} {
		fmt.Println(SquareArea(value))
	}
	//Output:
	// 1
	// 4
	// 9
	// 16
	// 25
}

func BenchmarkSquareArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range []float64{1, 2, 3, 4, 5} {
			SquareArea(value)
		}
	}
}

func TestRectangleArea(t *testing.T) {
	for _, value := range twoIns2 {
		got := RectangleArea(value.InA, value.InB)
		if got != value.Expected {
			t.Errorf("got %.2f expected %.2f", got, value.Expected)
		}
	}
}

func ExampleRectangleArea() {
	// {InA: 1, InB: 2}
	// {InA: 2, InB: 2}
	// {InA: 3, InB: 2}
	// {InA: 4, InB: 2}
	// {InA: 5, InB: 2}
	for _, value := range twoIns2 {
		fmt.Println(RectangleArea(value.InA, value.InB))
	}
	//Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}

func BenchmarkRectangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range twoIns2 {
			RectangleArea(value.InA, value.InB)
		}
	}
}

func TestDiamondArea(t *testing.T) {
	for _, value := range twoIns1 {
		got := DiamondArea(value.InA, value.InB)
		if got != value.Expected {
			t.Errorf("got %.2f expected %.2f", got, value.Expected)
		}
	}
}

func ExampleDiamondArea() {
	// {InA: 1, InB: 2}
	// {InA: 2, InB: 2}
	// {InA: 3, InB: 2}
	// {InA: 4, InB: 2}
	// {InA: 5, InB: 2}
	for _, value := range twoIns1 {
		fmt.Println(DiamondArea(value.InA, value.InB))
	}
	//Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func BenchmarkDiamondArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range twoIns1 {
			DiamondArea(value.InA, value.InB)
		}
	}
}

func TestTrapezeArea(t *testing.T) {

	for _, value := range threeIns {
		got := TrapezeArea(value.InA, value.InB, value.InC)
		if got != value.Expected {
			t.Errorf("got %.2f expected %.2f", got, value.Expected)
		}
	}
}

func ExampleTrapezeArea() {
	// 	{InA: 1, InB: 5, InC: 5}
	//	{InA: 2, InB: 5, InC: 4}
	//	{InA: 3, InB: 5, InC: 3}
	//	{InA: 4, InB: 5, InC: 2}
	//	{InA: 5, InB: 5, InC: 1}
	for _, value := range threeIns {
		fmt.Println(TrapezeArea(value.InA, value.InB, value.InC))
	}
	//Output:
	// 15
	// 14
	// 12
	// 9
	// 5
}

func BenchmarkTrapezeArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range threeIns {
			TrapezeArea(value.InA, value.InB, value.InC)
		}
	}
}

func TestParallelogramArea(t *testing.T) {

	for _, value := range twoIns2 {
		got := ParallelogramArea(value.InA, value.InB)
		if got != value.Expected {
			t.Errorf("got %.2f expected %.2f", got, value.Expected)
		}
	}
}

func ExampleParallelogramArea() {
	// {InA: 1, InB: 2}
	// {InA: 2, InB: 2}
	// {InA: 3, InB: 2}
	// {InA: 4, InB: 2}
	// {InA: 5, InB: 2}
	for _, value := range twoIns2 {
		fmt.Println(ParallelogramArea(value.InA, value.InB))
	}
	//Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}

func BenchmarkParallelogramArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range twoIns2 {
			ParallelogramArea(value.InA, value.InB)
		}
	}
}

func TestPentagonArea(t *testing.T) {

	for _, value := range twoIns1 {
		got := PentagonArea(value.InA, value.InB)
		if got != value.Expected {
			t.Errorf("got %.2f expected %.2f", got, value.Expected)
		}
	}
}

func ExamplePentagonArea() {
	// {InA: 1, InB: 2}
	// {InA: 2, InB: 2}
	// {InA: 3, InB: 2}
	// {InA: 4, InB: 2}
	// {InA: 5, InB: 2}
	for _, value := range twoIns1 {
		fmt.Println(PentagonArea(value.InA, value.InB))
	}
	//Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func BenchmarkPentagonArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range twoIns1 {
			PentagonArea(value.InA, value.InB)
		}
	}
}

func TestHexagonArea(t *testing.T) {

	for _, value := range twoIns1 {
		got := HexagonArea(value.InA, value.InB)
		if got != value.Expected {
			t.Errorf("got %.2f expected %.2f", got, value.Expected)
		}
	}
}

func ExampleHexagonArea() {
	// {InA: 1, InB: 2}
	// {InA: 2, InB: 2}
	// {InA: 3, InB: 2}
	// {InA: 4, InB: 2}
	// {InA: 5, InB: 2}
	for _, value := range twoIns1 {
		fmt.Println(HexagonArea(value.InA, value.InB))
	}
	//Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func BenchmarkHexagonArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range twoIns1 {
			HexagonArea(value.InA, value.InB)
		}
	}
}

func TestCircleArea(t *testing.T) {

	oneIns := []oneIn{
		{InA: 1, Expected: 3.141592653589793},
		{InA: 2, Expected: 12.566370614359172},
		{InA: 3, Expected: 28.274333882308138},
		{InA: 4, Expected: 50.26548245743669},
		{InA: 5, Expected: 78.53981633974483},
	}

	for _, value := range oneIns {
		got := CircleArea(value.InA)
		if got != value.Expected {
			t.Errorf("got %v expected %v", got, value.Expected)
		}
	}
}

func ExampleCircleArea() {
	for _, value := range []float64{1, 2, 3, 4, 5} {
		fmt.Println(CircleArea(value))
	}
	//Output:
	// 3.141592653589793
	// 12.566370614359172
	// 28.274333882308138
	// 50.26548245743669
	// 78.53981633974483
}

func BenchmarkCircleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range []float64{1, 2, 3, 4, 5} {
			CircleArea(value)
		}
	}
}
