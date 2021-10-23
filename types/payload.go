package types

import (
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
