package p2p

/*
Peer is defines a node in the network
*/
type Peer interface {
}

/*
Transport defines a channel off communication (anything that handles communication)
example : - TCP, UDP, Websocket
*/
type Transaction interface {
	// ListenAndAccept listens for incoming connections and accepts them.
	// It blocks until an error occurs or the listener is closed.
	// Returns an error if there was a problem setting up the listener.
	ListenAndAccept() error
}
