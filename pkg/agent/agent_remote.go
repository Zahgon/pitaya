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

package agent

import (
	"context"
	"net"

	"github.com/topfreegames/pitaya/v3/pkg/cluster"
	"github.com/topfreegames/pitaya/v3/pkg/conn/codec"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
	"github.com/topfreegames/pitaya/v3/pkg/session"
)

// Remote corresponding to another server
type Remote struct {
	Session          session.Session // session
	chDie            chan struct{}   // wait for close
	messageEncoder   message.Encoder
	encoder          codec.PacketEncoder      // binary encoder
	frontendID       string                   // the frontend that sent the request
	reply            string                   // nats reply topic
	rpcClient        cluster.RPCClient        // rpc client
	serializer       serialize.Serializer     // message serializer
	serviceDiscovery cluster.ServiceDiscovery // service discovery
}

// NewRemote create new Remote instance
func NewRemote(
	sess *protos.Session,
	reply string,
	rpcClient cluster.RPCClient,
	encoder codec.PacketEncoder,
	serializer serialize.Serializer,
	serviceDiscovery cluster.ServiceDiscovery,
	frontendID string,
	messageEncoder message.Encoder,
	sessionPool session.SessionPool,
) (*Remote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO this is totally coupled with NATS

// binding session

// Kick kicks the user
func (a *Remote) Kick(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Push pushes the message to the user
func (a *Remote) Push(route string, v interface{}) error { _ = "STUB: not implemented"; return nil }

// ResponseMID reponds the message with mid to the user
func (a *Remote) ResponseMID(ctx context.Context, mid uint, v interface{}, isError ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the remote
func (a *Remote) Close() error {
	_ = "STUB: not implemented"

	// RemoteAddr returns the remote address of the user
	return nil
}

func (a *Remote) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (a *Remote) serialize(m pendingMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// construct message and encode

// packet encode

func (a *Remote) send(m pendingMessage, to string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (a *Remote) sendPush(m pendingMessage, userID string, sv *cluster.Server) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// SendRequest sends a request to a server
func (a *Remote) SendRequest(ctx context.Context, serverID, reqRoute string, v interface{}) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
