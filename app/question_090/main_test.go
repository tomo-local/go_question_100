package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("Setup: running tests...")
	code := m.Run()
	println("Teardown: tests finished.")
	os.Exit(code)
}

func TestSomething(t *testing.T) {
	if 1+1 != 2 {
		t.Error("1+1 should be 2")
	}
}
