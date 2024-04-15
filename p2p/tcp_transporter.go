package p2p

import (
	"fmt"
	"net"
)

/*
TCPPeer represents a remote node over a TCP connection.

It encapsulates information about the TCP connection to a remote peer,
including the connection itself and whether it's an outbound connection
initiated by the local node or an inbound connection received from the remote peer.

Example usage:
- Creating a TCPPeer instance to represent a remote node connected over TCP.
- Storing information about the connection and its directionality (outbound or inbound).

Fields:
- conn: The underlying TCP connection to the remote peer.
- outbound: A boolean indicating whether the connection is outbound (initiated by the local node).

Note: TCPPeer instances can be used in various network communication scenarios, such as peer-to-peer
networks, client-server applications, or distributed systems using TCP as the transport protocol.
*/
type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransporterOptions struct {
	ListenAddr           string
	HandShaker           ShakeHandsFunc
	Decoder              Decoder
	OnNewConnectionAdded func(Peer) error
}

type TCPTransporter struct {
	TCPTransporterOptions              // The network address (host:port) on which this transporter listens for incoming connections.
	listener              net.Listener // The net.Listener object responsible for accepting incoming connections.
	rpcChannel            chan RPC
}

func NewTCPTransporter(options TCPTransporterOptions) *TCPTransporter {
	return &TCPTransporter{
		TCPTransporterOptions: options,
		rpcChannel:            make(chan RPC),
	}
}

// Listen attempts to create a TCP listener on the specified address.
// It sets up the underlying listener without starting the accept loop.
// Returns an error if there was a problem setting up the listener.
func (t *TCPTransporter) Listen() error {
	var err error

	// Attempt to create a TCP listener on the specified address.
	t.listener, err = net.Listen("tcp", t.ListenAddr)

	// If an error occurs during listener creation, return the error.
	if err != nil {
		return err
	}

	// Return nil to indicate successful listener setup.
	return nil
}

// ListenAndAccept listens on the specified address for incoming TCP connections and accepts them.
// It sets up the underlying listener and starts an accept loop in a separate goroutine.
// Returns an error if there was a problem setting up the listener or starting the accept loop.
func (t *TCPTransporter) ListenAndAccept() error {
	var err error

	// Attempt to create a TCP listener on the specified address.
	t.listener, err = net.Listen("tcp", t.ListenAddr)

	// If an error occurs during listener creation, log the error and return it.
	if err != nil {
		fmt.Printf("Error listening to %s: %v", t.ListenAddr, err)
		return err
	}

	// Start an accept loop in a separate goroutine to handle incoming connections.
	go t.acceptLoop()

	// Return nil to indicate successful setup and start of the listener.
	return nil
}

// serveLoop continuously accepts incoming connections on the TCPTransporter's listener.
// It spawns a new goroutine to handle each accepted connection.
// This function runs indefinitely until an error occurs or the listener is closed.
func (t *TCPTransporter) acceptLoop() error {
	for {
		// Accept a new connection from the listener
		conn, err := t.listener.Accept()

		// Check if there was an error accepting the connection
		if err != nil {
			// Log the error and return it
			fmt.Printf("Error accepting TCP: %v\n", err)
			return err
		}
		fmt.Printf("new incoming connection %v\n", conn)

		// Handle the accepted connection in a separate goroutine
		go t.handleConnection(conn)
	}
}

/*
Consume returns a receive-only channel for consuming Remote Procedure Call (RPC) messages.
It provides a means to receive RPC messages sent over the TCP connection.
*/
func (t *TCPTransporter) Consume() <-chan RPC {
	return t.rpcChannel
}

func (t *TCPTransporter) handleConnection(conn net.Conn) {
	tcpPeer := NewTCPPeer(conn, false)

	var err error
	defer func() {
		fmt.Printf("error connecting %s", err)
		fmt.Printf("Dropping connection ...\n")
		conn.Close()
	}()
	if err = t.HandShaker(tcpPeer); err != nil {
		return
	}

	if t.OnNewConnectionAdded != nil {
		if err := t.OnNewConnectionAdded(tcpPeer); err != nil {
			return
		}
	}
	message := RPC{}
	//buffer := make([]byte, 12000)
	for {

		if err := t.Decoder.Decode(conn, &message); err != nil {
			fmt.Printf("There was an error decoding message: %v\n connection %v", err, conn)
			continue
		}

		message.Origin = conn.RemoteAddr()
		t.rpcChannel <- message

	}

}
