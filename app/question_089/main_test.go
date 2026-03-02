package main

import "testing"

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Test_Greet(t *testing.T) {
	got := greet("World")
	want := "Hello, World!"
	assertCorrectMessage(t, got, want)
}
