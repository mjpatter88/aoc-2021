package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	got := Solve(input)
	want := 7

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
