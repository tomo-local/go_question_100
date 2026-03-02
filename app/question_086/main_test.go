package main

import "testing"

func Test_Add(t *testing.T) {
	got := add(2, 4)
	want := 6
	if got != want {
		t.Errorf("add(2, 4) = %d; want %d", got, want)
	}
}
