package lang

import (
	"fmt"
	"io"

	"github.com/pemistahl/lingua-go"
)

func Detect(source io.Reader) (string, error) {
	raw, err := io.ReadAll(source)
	if err != nil {
		return "", err
	}
	detector := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		Build()

	if language, exists := detector.DetectLanguageOf(string(raw)); exists {
		return language.String(), nil
	}

	return "unknown", fmt.Errorf("unknown language by lingua-go")

}
