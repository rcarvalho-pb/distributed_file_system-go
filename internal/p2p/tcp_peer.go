package p2p

import "net"

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbount bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbount,
	}
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

func OnPeerFunc(peer Peer) error {
	peer.Close()
	return nil
}
