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
		OnPeer:        p2p.OnPeerFunc,
	}
	tr := p2p.NewTCPTransport(opts)

	go func() {
		for {
			msg := <-tr.Consume()
			log.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
