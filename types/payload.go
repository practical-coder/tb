package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	BinaryType uint8 = iota + 1
	StringType
	MaxPayloadSize uint32 = 10 << 20 // 10MB
)

var ErrMaxPayloadSize = errors.New("MaxPayloadSize exceeded")

type Payload interface {
	fmt.Stringer
	io.ReaderFrom
	io.WriterTo
	Bytes() []byte
}

func decode(r io.Reader) (Payload, error) {
	var payloadType uint8
	err := binary.Read(r, binary.BigEndian, &payloadType)
	if err != nil {
		return nil, err
	}

	var payload Payload

	switch payloadType {
	case BinaryType:
		payload = new(Binary)
	case StringType:
		payload = new(String)
	default:
		return nil, errors.New("Unkown Type")
	}

	_, err = payload.ReadFrom(
		io.MultiReader(
			bytes.NewReader([]byte{payloadType}),
			r,
		),
	)

	if err != nil {
		return nil, err
	}

	return payload, nil
}
