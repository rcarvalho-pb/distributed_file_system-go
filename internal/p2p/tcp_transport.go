package p2p

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

var ErrInvalidHandshake error = errors.New("invalid handshake")

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

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.TCPTransportOpts.ListenAddr)

	go t.startAcceptLoop()

	return err
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error in accept loop: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

type Tmp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) error {
	peer := NewTCPPeer(conn, true)

	if err := t.TCPTransportOpts.HandshakeFunc(peer); err != nil {
		conn.Close()
		return ErrInvalidHandshake
	}

	msg := &Tmp{}
	for {
		if err := t.TCPTransportOpts.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error from decoder in handleConn: %s\n", err)
			continue
		}
		break
	}
	fmt.Printf("TCP accept from: %s\n", t.listeningAddress)
	return nil
}
