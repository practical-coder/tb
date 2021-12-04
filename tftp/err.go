package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Err struct {
	Error   ErrCode
	Message string
}

func (e Err) MarshalBinary() ([]byte, error) {
	// operation code + error code + message + null byte
	cap := 2 + 2 + len(e.Message) + 1
	b := new(bytes.Buffer)
	b.Grow(cap)

	err := binary.Write(b, binary.BigEndian, OpErr) // write operation code
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, e.Error) // write error code
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(e.Message) // write error message
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0) // write null byte
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (e *Err) UnmarshalBinary(p []byte) error {
	r := bytes.NewBuffer(p)

	var opcode OpCode

	err := binary.Read(r, binary.BigEndian, &opcode) // read operation code
	if err != nil {
		return err
	}

	if opcode != OpErr {
		return errors.New("invalid operation code!")
	}

	err = binary.Read(r, binary.BigEndian, &e.Error) // read error code
	if err != nil {
		return err
	}

	e.Message, err = r.ReadString(0)
	e.Message = RemoveNullByte(e.Message)
	if len(e.Message) == 0 {
		return errors.New("0 length error message")
	}

	return err
}
