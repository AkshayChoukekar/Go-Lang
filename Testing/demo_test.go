package main2

import "testing"

func TestCalculateSqaure(t *testing.T) {
	want := 4
	got := CalculateSqaure(2)

	if got != want {
		t.Errorf("Square calculations failed")
	}

}
