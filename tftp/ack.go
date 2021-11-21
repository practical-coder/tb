package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Ack uint16

func (ack Ack) MarshalBinary() ([]byte, error) {
	cap := 2 + 2 // operation code + block number

	b := new(bytes.Buffer)
	b.Grow(cap)

	err := binary.Write(b, binary.BigEndian, OpAck) // Acknowledgement operation code
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, ack) // write block number
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (ack *Ack) UnmarshalBinary(p []byte) error {
	var opcode OpCode

	r := bytes.NewReader(p)

	err := binary.Read(r, binary.BigEndian, &opcode) // read operation code
	if err != nil {
		return err
	}

	if opcode != OpAck {
		return errors.New("invalid operation code")
	}

	return binary.Read(r, binary.BigEndian, ack) // read block number
}
