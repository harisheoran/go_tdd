package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Aryan", "Eng")
		want := "Hello, Aryan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello world when no name is passed", func(t *testing.T) {
		got := Hello("", "Eng")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello in French", func(t *testing.T) {
		got := Hello("Felatu", "French")
		want := "Bonjour, Felatu"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
