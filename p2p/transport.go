package p2p

/*
Peer represents a node within the network.
*/
type Peer interface {
	// Close closes the connection associated with the peer.
	Close() error
}

/*
Transaction defines a communication channel, encompassing various protocols such as TCP, UDP, or WebSockets.
*/
type Transaction interface {
	// ListenAndAccept sets up a listener for incoming connections and accepts them.
	// It blocks until an error occurs or the listener is closed.
	// Returns an error if there was a problem setting up the listener.
	ListenAndAccept() error

	// Consume returns a channel for receiving Remote Procedure Call (RPC) messages.
	Consume() <-chan RPC
}
