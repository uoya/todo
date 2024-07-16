package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello("John")
	if result != "Hello John" {
		t.Fatal("failed test")
	}
}
