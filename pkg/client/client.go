// Copyright (c) TFG Co. All Rights Reserved.
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

package client

import (
	"bytes"
	"crypto/tls"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v3/pkg/conn/codec"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/conn/packet"
	"github.com/topfreegames/pitaya/v3/pkg/session"
)

// HandshakeSys struct
type HandshakeSys struct {
	Dict       map[string]uint16 `json:"dict"`
	Heartbeat  int               `json:"heartbeat"`
	Serializer string            `json:"serializer"`
}

// HandshakeData struct
type HandshakeData struct {
	Code int          `json:"code"`
	Sys  HandshakeSys `json:"sys"`
}

type pendingRequest struct {
	msg    *message.Message
	sentAt time.Time
}

// Client struct
type Client struct {
	conn                net.Conn
	Connected           bool
	packetEncoder       codec.PacketEncoder
	packetDecoder       codec.PacketDecoder
	packetChan          chan *packet.Packet
	IncomingMsgChan     chan *message.Message
	pendingChan         chan bool
	pendingRequests     map[uint]*pendingRequest
	pendingReqMutex     sync.Mutex
	requestTimeout      time.Duration
	closeChan           chan struct{}
	nextID              uint32
	messageEncoder      message.Encoder
	clientHandshakeData *session.HandshakeData
}

// MsgChannel return the incoming message channel
func (c *Client) MsgChannel() chan *message.Message { _ = "STUB: not implemented"; return nil }

// ConnectedStatus return the connection status
func (c *Client) ConnectedStatus() bool {
	_ = "STUB: not implemented"

	// New returns a new client
	return false
}

func New(logLevel logrus.Level, requestTimeout ...time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// 30 here is the limit of inflight messages
// TODO this should probably be configurable

// SetClientHandshakeData sets the data to send inside handshake
func (c *Client) SetClientHandshakeData(data *session.HandshakeData) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) sendHandshakeRequest() error { _ = "STUB: not implemented"; return nil }

func (c *Client) handleHandshakeResponse() error { _ = "STUB: not implemented"; return nil }

// pendingRequestsReaper delete timedout requests
func (c *Client) pendingRequestsReaper() { _ = "STUB: not implemented"; return }

// send a timeout to incoming msg chan

func (c *Client) handlePackets() { _ = "STUB: not implemented"; return }

//handle data

// do not process msg for already timedout request

func (c *Client) readPackets(buf *bytes.Buffer) ([]*packet.Packet, error) {
	_ = "STUB: not implemented"
	// listen for sv messages
	return nil, nil
}

func (c *Client) handleServerMessages() { _ = "STUB: not implemented"; return }

func (c *Client) sendHeartbeats(interval int) { _ = "STUB: not implemented"; return }

// Disconnect disconnects the client
func (c *Client) Disconnect() { _ = "STUB: not implemented"; return }

// ConnectTo connects to the server at addr, for now the only supported protocol is tcp
// if tlsConfig is sent, it connects using TLS
func (c *Client) ConnectTo(addr string, tlsConfig ...*tls.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// ConnectToWS connects using webshocket protocol
func (c *Client) ConnectToWS(addr string, path string, tlsConfig ...*tls.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) handleHandshake() error { _ = "STUB: not implemented"; return nil }

// SendRequest sends a request to the server
func (c *Client) SendRequest(route string, data []byte) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SendNotify sends a notify to the server
func (c *Client) SendNotify(route string, data []byte) error { _ = "STUB: not implemented"; return nil }

func (c *Client) buildPacket(msg message.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sendMsg sends the request to the server
func (c *Client) sendMsg(msgType message.Type, route string, data []byte) (uint, error) {
	_ = "STUB: not implemented"
	// TODO mount msg and encode
	return 0, nil
}
