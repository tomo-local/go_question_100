package main

import "testing"

type MockUserFetcher struct {
	UserName string
	Err      error
}

func (m *MockUserFetcher) Fetch(userID int) (string, error) {
	return m.UserName, m.Err
}

func Test_GetUser(t *testing.T) {
	mock := &MockUserFetcher{UserName: "Alice", Err: nil}
	got, err := GetUser(mock, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got != "Alice" {
		t.Errorf("GetUser(mock, 1) = %q; want %q", got, "Alice")
	}

}
