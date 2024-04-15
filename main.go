package main

import (
	"fmt"

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

	go func() {
		for {
			message := <-tcpTransporter.Consume()
			fmt.Printf("Received message from %v: %v\n", message.Origin, message.Data)
		}
	}()
	// Listen for incoming connections and accept them.
	// This call blocks until an error occurs or the listener is closed.
	if err := tcpTransporter.ListenAndAccept(); err != nil {
		// If an error occurs during listening and accepting connections, log the error and exit the program.
		fmt.Printf("Error listening to connection \n%v", err)
	}

	// The following select statement blocks indefinitely, ensuring that the program continues to run
	// and handle incoming connections or events. Since there are no cases provided, it effectively
	// serves as a way to prevent the main function from returning immediately.
	select {}

}
