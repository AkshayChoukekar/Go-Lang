package main

import (
	"log"
	"testing"
)

func TestSayHello(t *testing.T) {
	if SayHello("Akshay") != "hello mahi" {
		t.Errorf("got %q want %q", got, want)
	}
}
