package p2p

import (
	"encoding/gob"
	"fmt"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct {
}

func (d GOBDecoder) Decode(r io.Reader, msg *RPC) error {
	return gob.NewDecoder(r).Decode(msg)
}

type NoDecoderNeeded struct {
}

func (d NoDecoderNeeded) Decode(r io.Reader, msg *RPC) error {
	buffer := make([]byte, 1028)
	n, err := r.Read(buffer)
	if err != nil {
		return err
	}
	msg.Data = buffer[:n]
	fmt.Printf(string(buffer)[:n])
	return nil
}
