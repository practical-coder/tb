package main

import (
	"testing"
	"unicode/utf8"
)

func TestRuneLen(t *testing.T) {
	city := "Kraków"
	length := len([]rune(city))
	t.Logf("length: %d", length)
	if length != 6 {
		t.Errorf("Expected 6, got: %d", length)
	}
	length = utf8.RuneCountInString(city)
	t.Logf("length: %d", length)
	if length != 6 {
		t.Errorf("Expected 6, got: %d", length)
	}
}
