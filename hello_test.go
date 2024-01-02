package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Del!")
	want := "Hello, Del!"

	if got != want {
		t.Errorf("go %q want %q", got, want)
	}
}
