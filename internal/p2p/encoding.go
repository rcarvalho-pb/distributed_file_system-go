package p2p

import (
	"encoding/gob"
	"io"
)

type DecodeFunc func(io.Reader, *Message) error

func GOBDecodeFunc(r io.Reader, m *Message) error {
	return gob.NewDecoder(r).Decode(m)
}

func DefaultDecodeFunc(r io.Reader, m *Message) error {
	buf := make([]byte, 1028)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	m.Payload = buf[:n]
	return nil
}
