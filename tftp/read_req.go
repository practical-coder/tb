package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"
)

type ReadReq struct {
	Filename string
	Mode     string
}

func (rr ReadReq) MarshalBinary() ([]byte, error) {
	mode := "octet"
	if rr.Mode != "" {
		mode = rr.Mode
	}

	// operation code + filename-length + 1 null-byte + model-length + 1 null-byte
	cap := 2 + 2 + len(rr.Filename) + 1 + len(rr.Mode) + 1

	b := new(bytes.Buffer)
	b.Grow(cap)

	err := binary.Write(b, binary.BigEndian, OpRRQ) // write operation code
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(rr.Filename) // write filename
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0) // write separating null-byte
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(mode) // write mode
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0) // write closing null-byte
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (rr *ReadReq) UnmarshalBinary(p []byte) error {
	r := bytes.NewBuffer(p)

	var code OpCode

	err := binary.Read(r, binary.BigEndian, &code) // read operation code
	if err != nil {
		return err
	}

	if code != OpRRQ {
		return errors.New("invalid RRQ")
	}

	rr.Filename, err = r.ReadString(0) // read until null-byte separator
	if err != nil {
		return errors.New("Read filename up to and INCLUDING null-byte")
	}

	rr.Filename = RemoveNullByte(rr.Filename)
	if len(rr.Filename) == 0 {
		return errors.New("invalid filename: 0 length")
	}

	rr.Mode, err = r.ReadString(0) // read until ending null-byte
	if err != nil {
		return errors.New("Read mode up to and INCLUDING null-byte")
	}

	rr.Mode = RemoveNullByte(rr.Mode)
	if len(rr.Mode) == 0 {
		return errors.New("invalid mode: 0 length")
	}

	actualMode := strings.ToLower(rr.Mode)
	if actualMode != "octet" {
		return errors.New("only octet mode supported!")
	}

	return nil
}
