// |****************************************|
// |				Testing                 |
// |****************************************|
package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	x := Sum(2, 3, 4)
	if x != 9 {
		t.Error("Expected: ", 9, "Got: ", x)
	}
}

func TestMul(t *testing.T) {
	x := Mul(2, 3)
	if x != 6 {
		t.Error("Expected: ", 6, "Got: ", x)
	}
}

func TestDiv(t *testing.T) {
	x := Div(8, 4)
	if x != 2 {
		t.Error("Expected: ", 2, "Got: ", x)
	}
}