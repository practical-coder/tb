package which

import "testing"

func TestFind(t *testing.T) {
	input := "tldr"
	expected := []string{"/usr/local/bin/tldr"}
	results := Find(input)
	if expected[0] != results[0] {
		t.Errorf("Expected %s, got %s", expected[0], results[0])
	}
}
