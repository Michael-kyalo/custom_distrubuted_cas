package main

import (
	"log"

	"github.com/Michael-kyalo/custom_distrubuted_cas/p2p"
)

func main() {

	tcpOptions := p2p.TCPTransporterOptions{
		ListenAddr: ":8103",
		HandShaker: p2p.NOShakeNeeded,
		Decoder:    p2p.NoDecoderNeeded{},
	}

	// Create a new TCPTransporter instance to handle peer-to-peer communication over TCP on port 3000.
	tcpTransporter := p2p.NewTCPTransporter(tcpOptions)

	// Listen for incoming connections and accept them.
	// This call blocks until an error occurs or the listener is closed.
	if err := tcpTransporter.ListenAndAccept(); err != nil {
		// If an error occurs during listening and accepting connections, log the error and exit the program.
		log.Printf("Error listening to connection \n%v", err)
	}

	// The following select statement blocks indefinitely, ensuring that the program continues to run
	// and handle incoming connections or events. Since there are no cases provided, it effectively
	// serves as a way to prevent the main function from returning immediately.
	select {}

}
