package tftp

import (
	"strings"
)

const (
	HeaderSize   = 4
	BlockSize    = 512
	DatagramSize = HeaderSize + BlockSize
)

type OpCode uint16

const (
	OpRRQ OpCode = iota + 1
	_            // WRQ unsupported
	OpData
	OpAck
	OpErr
)

type ErrCode uint16

const (
	ErrUnknown ErrCode = iota
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrUnknownID
	ErrFileExists
	ErrNoUser
)

func RemoveNullByte(name string) string {
	return strings.TrimRight(name, "\x00")
}
