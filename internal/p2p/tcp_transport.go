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

func (t *TCPTransport) handleConn(conn net.Conn) error {
	peer := NewTCPPeer(conn, true)
	defer conn.Close()

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		return ErrInvalidHandshake
	}

	// msg := &Message{}
	buf := make([]byte, 2000)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("TCP error: %s\n", err)
			break
		}

		fmt.Println("message:", string(buf[:n]))
		// if err := t.Decoder.Decode(conn, msg); err != nil {
		// 	fmt.Printf("TCP error from decoder in handleConn: %s\n", err)
		// 	continue
		// }
	}
	return nil
}
