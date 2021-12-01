package main

import (
	"testing"
)

var input = []int{
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

func TestSolve(t *testing.T) {
	got := Solve(input)
	want := 7

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSolve2(t *testing.T) {
	got := Solve2(input)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCalcWindowSum(t *testing.T) {
	t.Run("Calc window at 0", func(t *testing.T) {
		got, err := CalcWindowSum(input, 0)
		want := input[0] + input[1] + input[2]
		if err != nil {
			t.Errorf("got an unexpected error")
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Calc window in middle", func(t *testing.T) {
		got, err := CalcWindowSum(input, 3)
		want := input[3] + input[4] + input[5]
		if err != nil {
			t.Errorf("got an unexpected error")
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("Calc window not enough elements", func(t *testing.T) {
		_, err := CalcWindowSum(input, 8)
		if err == nil {
			t.Errorf("expected an error but did not get one")
		}
	})
}
