package codec

import "github.com/topfreegames/pitaya/v3/pkg/conn/packet"

// ParseHeader parses a packet header and returns its dataLen and packetType or an error
func ParseHeader(header []byte) (int, packet.Type, error) {
	_ = "STUB: not implemented"
	return 0, *new(packet.Type), nil
}

// BytesToInt decode packet data length byte to int(Big end)
func BytesToInt(b []byte) int { _ = "STUB: not implemented"; return 0 }

// IntToBytes encode packet data length to bytes(Big end)
func IntToBytes(n int) []byte { _ = "STUB: not implemented"; return nil }
