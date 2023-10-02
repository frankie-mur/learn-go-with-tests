package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Frankie")

	want := "Hello, Frankie"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
