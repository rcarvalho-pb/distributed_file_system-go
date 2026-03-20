package p2p

type HandshakeFunc func(Peer) error

func NOOPHandshake(_ Peer) error { return nil }
