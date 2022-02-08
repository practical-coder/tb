package txt

import (
	"bytes"
	"testing"
)

func TestLines(t *testing.T) {
	buff := bytes.NewBufferString("line1\nline2\nline3\nline4\nline5")
	expected := 5
	actual := countLines(buff)
	if actual != expected {
		t.Errorf("countLines Error actual: %d expected: %d", actual, expected)
	}
}

func TestWords(t *testing.T) {
	buff := bytes.NewBufferString("word1 word2 word3 word4 word5 word6\n")
	expected := 6
	actual := countWords(buff)
	if actual != expected {
		t.Errorf("countWords Error actual: %d expected: %d", actual, expected)
	}
}

func TestBytes(t *testing.T) {
	buff := bytes.NewBufferString("0123456789")
	expected := 10
	actual := countBytes(buff)
	if actual != expected {
		t.Errorf("countBytes Error actual: %d expected: %d", actual, expected)
	}
}

func TestRunes(t *testing.T) {
	buff := []byte("gźegźółka")
	expected := 9
	actual := countRunes(buff)
	if actual != expected {
		t.Errorf("countRunes Error actual: %d expected: %d", actual, expected)
	}
}
