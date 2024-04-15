package p2p

import (
	"net"
)

// RPC represents a data packet or message.
type RPC struct {
	//address of the sender
	Origin net.Addr
	// Data contains the payload of the RPC.
	Data []byte
}
