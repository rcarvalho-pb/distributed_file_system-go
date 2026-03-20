package main

import (
	"log"

	"github.com/rcarvalho-pb/distributed_file_system-go/internal/p2p"
)

func main() {
	println("test")
	port := ":3000"
	decoder := &p2p.TCPDecoder{}
	opts := p2p.TCPTransportOpts{
		ListenAddr:    port,
		HandshakeFunc: p2p.NOOPHandshake,
		Decoder:       decoder,
	}
	tr := p2p.NewTCPTransport(opts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
