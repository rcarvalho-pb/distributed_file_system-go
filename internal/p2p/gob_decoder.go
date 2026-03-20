package p2p

import (
	"encoding/gob"
	"io"
)

func GobDecodeFunc(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}
