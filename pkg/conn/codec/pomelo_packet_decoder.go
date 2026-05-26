// Copyright (c) nano Author and TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package codec

import (
	"bytes"

	"github.com/topfreegames/pitaya/v3/pkg/conn/packet"
)

// PomeloPacketDecoder reads and decodes network data slice following pomelo's protocol
type PomeloPacketDecoder struct{}

// NewPomeloPacketDecoder returns a new decoder that used for decode network bytes slice.
func NewPomeloPacketDecoder() *PomeloPacketDecoder { _ = "STUB: not implemented"; return nil }

func (c *PomeloPacketDecoder) forward(buf *bytes.Buffer) (int, packet.Type, error) {
	_ = "STUB: not implemented"
	return 0, *new(packet.Type), nil
}

// Decode decode the network bytes slice to packet.Packet(s)
func (c *PomeloPacketDecoder) Decode(data []byte) ([]*packet.Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check length

// first time

// if no more packets, break
