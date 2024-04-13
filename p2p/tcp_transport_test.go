package p2p

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransporter(t *testing.T) {
	tests := []struct {
		options       TCPTransporterOptions
		expectedError error
	}{
		{
			options: TCPTransporterOptions{
				ListenAddr: ":4000",
			},
			expectedError: nil,
		},
		{
			options: TCPTransporterOptions{
				ListenAddr: "4000",
			},
			expectedError: &net.OpError{Op: "listen", Net: "tcp", Source: nil, Addr: nil, Err: &net.AddrError{Err: "missing port in address", Addr: "4000"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.options.ListenAddr, func(t *testing.T) {
			tr := NewTCPTransporter(tt.options)

			assert.Equal(t, tr.ListenAddr, tt.options.ListenAddr)

			// Call ListenAndAccept and compare the returned error with the expected error.
			err := tr.ListenAndAccept()
			assert.Equal(t, err, tt.expectedError)
		})
	}
}
