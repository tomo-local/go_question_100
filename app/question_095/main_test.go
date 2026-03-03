package main

import "testing"

type MockDatastore struct {
	Users map[int]string
}

func (m *MockDatastore) GetUser(id int) string {
	return m.Users[id]
}

func TestApp_ShowUser(t *testing.T) {
	app := &App{store: &MockDatastore{Users: map[int]string{1: "Alice"}}}
	got := app.ShowUser(1)
	if got != "Alice" {
		t.Errorf("ShowUser(1) = %q; want Alice", got)
	}
}
