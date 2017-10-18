package main

import (
	"testing"
)

func TestTravis(t *testing.T) {
	if Travis(2, 2) != 4 {
		t.Error("Expected 4")
	}
}
