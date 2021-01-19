// |****************************************|
// |				Testing                 |
// |****************************************|
package main

import (
	"testing"
)

type dataTest1 struct {
	In  []int
	Out int
}

type dataTest2 struct {
	InA float64
	InB float64
	Out float64
}

func TestSum(t *testing.T) {

	dataTests := []dataTest1{
		{In: []int{2, 3, 4}, Out: 9},
		{In: []int{3, 3, 4}, Out: 10},
		{In: []int{4, 3, 4}, Out: 11},
		{In: []int{5, 3, 4}, Out: 12},
		{In: []int{6, 3, 4}, Out: 13},
	}

	for _, test := range dataTests {
		aux := Sum(test.In...)
		if aux != test.Out {
			t.Error("Expected: ", test.Out, "Got: ", aux)
		}
	}
}

func TestMul(t *testing.T) {

	dataTests := []dataTest2{
		{InA: 1, InB: 3, Out: 3},
		{InA: 2, InB: 3, Out: 6},
		{InA: 3, InB: 3, Out: 9},
		{InA: 4, InB: 3, Out: 12},
		{InA: 5, InB: 3, Out: 15},
	}

	for _, test := range dataTests {
		aux := Mul(test.InA, test.InB)
		if aux != test.Out {
			t.Error("Expected: ", test.Out, "Got: ", aux)
		}
	}
}

func TestDiv(t *testing.T) {

	dataTests := []dataTest2{
		{InA: 3, InB: 3, Out: 1},
		{InA: 6, InB: 3, Out: 2},
		{InA: 9, InB: 3, Out: 3},
		{InA: 12, InB: 3, Out: 4},
		{InA: 15, InB: 3, Out: 5},
	}

	for _, test := range dataTests {
		aux := Div(test.InA, test.InB)
		if aux != test.Out {
			t.Error("Expected: ", test.Out, "Got: ", aux)
		}
	}
}
