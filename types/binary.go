package types

import (
	"encoding/binary"
	"io"
)

type Binary []byte

func (b Binary) Bytes() []byte {
	return b
}

func (b Binary) String() string {
	return string(b)
}

func (b Binary) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, BinaryType)
	if err != nil {
		return 0, err
	}
	var n int64 = 1

	err = binary.Write(w, binary.BigEndian, uint32(len(b)))
	if err != nil {
		return n, err
	}
	n += 4

	o, err := w.Write(b)
	return n + int64(o), err
}
