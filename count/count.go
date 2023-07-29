package count

import (
	"io"
)

func FromReader(r io.Reader) (int, error) {
	buf := make([]byte, 16)
	counter := 0
	var err error
	for {
		c, err := r.Read(buf)
		counter += c
		if err != nil {
			break
		}
		if err == io.EOF {
			err = nil
			break
		}
	}
	return counter, err
}
