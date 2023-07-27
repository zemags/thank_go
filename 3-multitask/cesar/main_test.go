package main

import (
	"testing"
)

func TestEncode(t *testing.T) {
	src := "abc"
	res := encode(src)
	if res != "nop" {
		t.Error("Expected 'nop', got ", res)
	}
}
