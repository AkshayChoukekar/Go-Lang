package main

import (
	"testing"
)

func TestPostData(t *testing.T) {
	want := "200 OK"
	got := PostData()

	if got != want {
		t.Errorf("failed to post data")
	}
}

func TestGetData(t *testing.T) {
	want := "200 OK"
	got := GetData()

	if got != want {
		t.Errorf("failed to get data")
	}
}

func TestPutData(t *testing.T) {
	want := "200 OK"
	got := PutData()

	if got != want {
		t.Errorf("failed to put data")
	}
}

func TestDeleteData(t *testing.T) {
	want := "200 OK"
	got := DeleteData()

	if got != want {
		t.Errorf("failed to delete data")
	}
}
