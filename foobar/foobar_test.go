package foobar

import "testing"

// ... ...
func TestHelloWorld(t *testing.T) {
	if Method("asd") != "foobar: asd" {
		t.Fatal("Test fail")
	}
}
