package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello with name", func(t *testing.T) {
		got := Hello("Frankie", "")

		want := "Hello, Frankie"

		assertCorrectMsg(t, got, want)
	})
	t.Run("saying Hello, World when empty string provided", func(t *testing.T) {
		got := Hello("", "")

		want := "Hello, World"

		assertCorrectMsg(t, got, want)
	})
	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Francisco", "Spanish")

		want := "Hola, Francisco"

		assertCorrectMsg(t, got, want)
	})
	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Frank", "French")

		want := "Bonjur, Frank"

		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
