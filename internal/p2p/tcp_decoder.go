package p2p

import (
	"encoding/gob"
	"io"
)

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(reader io.Reader, v any) error {
	return gob.NewDecoder(reader).Decode(v)
}
