package p2p

import "io"

type DecodeFunc func(io.Reader, any) error
