package main

import "testing"

func TestCreateFile(t *testing.T) {
	got := CreateFile("C:\\Users\\SRS\\Desktop\\")
	want := "file created successfully"

	if got != want {
		t.Errorf("file creation failed")
	}
}

func TestReadFile(t *testing.T) {
	got := ReadFile("C:\\Users\\SRS\\Desktop\\NewFile.txt")
	want := " "
	if got != want {
		t.Errorf("file reading failed")
	}

}
