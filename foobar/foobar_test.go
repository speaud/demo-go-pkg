package foobar

import "testing"

func TestHelloWorld(t *testing.T) {
	if d("asd") != "Hello World!!asd" {
		t.Fatal("Test fail")
	}
}
