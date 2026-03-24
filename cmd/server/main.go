package main

import (
	"log"

	"github.com/rcarvalho-pb/distributed_file_system-go/internal/p2p"
)

func main() {
	log.Println("Starting on port 3000")
	opts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOOPHandshakeFunc,
		Decoder:       p2p.DefaultDecodeFunc,
	}
	tr := p2p.NewTCPTransport(opts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
