package main

import (
	"testing"
)

func TestEncode(t *testing.T) {
	funcs := []func() any{squared(5), squared(2), squared(1), squared(3), squared(4)}
	nums := gather(funcs)
	expected := []any{25, 4, 1, 9, 16}
	for i, v := range nums {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], v)
		}
	}
}
