package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	got := SayHello("Akshay")
	want := "Hello Akshay"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCalculateSqaure(t *testing.T) {
	want := 4
	got := CalculateSquare(2)

	if got != want {
		t.Errorf("Square calculations failed")
	}

}
