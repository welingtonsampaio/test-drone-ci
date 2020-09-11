package main

import "testing"

func Test_hello(t *testing.T) {
	r := hello("World")
	if r != "Hello World!" {
		t.Errorf("Wrong result: %s", r)
	}
}
