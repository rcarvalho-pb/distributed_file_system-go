package p2p

import "errors"

var ErrInvalidHandshake error = errors.New("invalid handshake")

type HandshakeFunc func(Peer) error

func NOOPHandshakeFunc(_ Peer) error { return nil }
