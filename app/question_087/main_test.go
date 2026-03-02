package main

import "testing"

func Test_Add(t *testing.T) {
	t.Run("Positive numbers", func(t *testing.T) {
		got := add(2, 4)
		want := 6
		if got != want {
			t.Errorf("add(2, 4) = %d; want %d", got, want)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		got := add(-2, -3)
		want := -5
		if got != want {
			t.Errorf("add(-2, -3) = %d; want %d", got, want)
		}
	})
}
