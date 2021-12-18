package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld("asd") != "Hello World!!asd" {
		t.Fatal("Test fail")
	}
}
