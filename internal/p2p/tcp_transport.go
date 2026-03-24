package p2p

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

var ErrInvalidPayload error = errors.New("invalid payload to decode")

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
	Decoder       DecodeFunc
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
	t.listener, err = net.Listen("tcp", t.ListenAddr)

	go t.startAcceptLoop()

	return err
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error in accept loop: %s\n", err)
		}
		fmt.Printf("new incoming connection: %+v\n", conn)
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	defer conn.Close()

	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Printf("TCP handshake error: %s\n", ErrInvalidHandshake)
		return
	}

	msg := &Message{}
	for {
		if err := t.Decoder(conn, msg); err != nil {
			if err == io.EOF {
				fmt.Printf("TCP connection ended in remote side\n")
				return
			}
			fmt.Printf("TCP decode error: %s\n", ErrInvalidPayload)
			return
		}
		msg.From = conn.RemoteAddr()
		fmt.Printf("message: %+v\n", msg)
	}
}
