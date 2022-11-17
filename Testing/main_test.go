package main

import (
	"log"
	"testing"
)

func TestSayHello(t *testing.T) {
	want := 124
	got := SayHello("akshay")

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}

}
