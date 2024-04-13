package p2p

import (
	"net"
)

// Message represents a data packet or message.
type Message struct {
	origin net.Addr
	// Data contains the payload of the message.
	Data []byte
}
