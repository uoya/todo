package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := true
	if result != true {
		t.Fatal("failed test")
	}
}
