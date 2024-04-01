package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	g := Hello()
	want := "Hello with Go!"

	if g != want {
		t.Errorf("Нужно %q, Получили %q", g, want)
	}
}
