package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	g := Hello("Mark")
	want := "Hello, Mark"

	if g != want {
		t.Errorf("Нужно %q, Получили %q", g, want)
	}
}
