package types

import (
	"encoding/binary"
	"fmt"
	"io"
)

type String string

func (s String) Bytes() []byte {
	return []byte(s)
}

func (s String) String() string {
	return string(s)
}

func (s String) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, StringType)
	if err != nil {
		return 0, err
	}
	var n int64 = 1

	err = binary.Write(w, binary.BigEndian, uint32(len(s)))
	if err != nil {
		return n, err
	}
	n += 4

	o, err := w.Write([]byte(s))

	return n + int64(o), err
}

func (s *String) ReadFrom(r io.Reader) (int64, error) {
	var payloadType uint8
	err := binary.Read(r, binary.BigEndian, &payloadType)
	if err != nil {
		return 0, err
	}
	var n int64 = 1

	if payloadType != StringType {
		return n, fmt.Errorf("invalid type of StringType")
	}

	var size uint32
	err = binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return n, err
	}
	n += 4

	if size > MaxPayloadSize {
		return n, ErrMaxPayloadSize
	}

	buf := make([]byte, size)
	o, err := r.Read(buf)
	*s = String(buf)

	return n + int64(o), err
}
