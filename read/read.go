package read

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func FromWeb(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func FromFile(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func From(source string) (io.ReadCloser, error) {
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return FromWeb(source)
	}

	return FromFile(source)
}
