package p2p

import (
	"encoding/gob"
	"io"
)

type DecodeFunc func(io.Reader, *RPC) error

func GOBDecodeFunc(r io.Reader, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

func DefaultDecodeFunc(r io.Reader, rpc *RPC) error {
	buf := make([]byte, 1028)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	rpc.Payload = buf[:n]
	return nil
}
