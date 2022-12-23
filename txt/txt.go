package txt

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func Count(path string) {
	var err error
	var data []byte

	if path == "-" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(path)
	}

	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("Error reading file: %s", path)
	}

	b := bytes.NewReader(data)
	w := bytes.NewReader(data)
	l := bytes.NewReader(data)

	log.Info().
		Str("file", path).
		Int("lines", countLines(l)).
		Int("words", countWords(w)).
		Int("bytes", countBytes(b)).
		Int("runes", countRunes(data)).
		Msg("")

}

func count(r io.Reader, splitFunc bufio.SplitFunc) int {
	s := bufio.NewScanner(r)
	s.Split(splitFunc)
	counter := 0
	for s.Scan() {
		counter++
	}

	return counter
}

func countBytes(r io.Reader) int {
	return count(r, bufio.ScanBytes)
}

func countWords(r io.Reader) int {
	return count(r, bufio.ScanWords)
}

func countLines(r io.Reader) int {
	return count(r, bufio.ScanLines)
}

func countRunes(data []byte) int {
	runes := []rune(string(data))
	return len(runes)
}
