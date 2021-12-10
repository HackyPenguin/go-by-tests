package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("in french", func(t *testing.T) {
		got := Hello("Nicolas", "french")
		want := "Bonjour, Nicolas"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in german", func(t *testing.T) {
		got := Hello("Ben", "german")
		want := "Hallo, Ben"
		assertCorrectMessage(t, got, want)

	})
	t.Run("Say Hello to People", func(t *testing.T) {
		got := Hello("Barry", "english")
		want := "Hello, Barry"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
