package main

import "testing"

func TestCountDigitsInWords(t *testing.T) {
	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWords(phrase)
	if counts["0ne"] != 1 {
		t.Errorf("expected 1, got %d", counts["0ne"])
	}
	if counts["1wo"] != 1 {
		t.Errorf("expected 1, got %d", counts["1wo"])
	}
	if counts["thr33"] != 2 {
		t.Errorf("expected 2, got %d", counts["thr33"])
	}
	if counts["4068"] != 4 {
		t.Errorf("expected 4, got %d", counts["4068"])
	}
}
