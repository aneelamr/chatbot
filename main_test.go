package main

import "testing"

func TestGetHi(t *testing.T) {
	expected := "hi"
	actual := getHi()
	if actual != expected {
		t.Errorf("getHi() = %v, want %v", actual, expected)
	}
}
