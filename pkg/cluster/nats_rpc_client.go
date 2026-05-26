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

package cluster

import (
	"context"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/route"
	"github.com/topfreegames/pitaya/v3/pkg/session"
)

// NatsRPCClient struct
type NatsRPCClient struct {
	conn                   *nats.Conn
	connString             string
	connectionTimeout      time.Duration
	maxReconnectionRetries int
	reqTimeout             time.Duration
	running                bool
	server                 *Server
	metricsReporters       []metrics.Reporter
	appDieChan             chan bool
	websocketCompression   bool
	reconnectJitter        time.Duration
	reconnectJitterTLS     time.Duration
	reconnectWait          time.Duration
	pingInterval           time.Duration
	maxPingsOutstanding    int
}

// NewNatsRPCClient ctor
func NewNatsRPCClient(
	config config.NatsRPCClientConfig,
	server *Server,
	metricsReporters []metrics.Reporter,
	appDieChan chan bool,
) (*NatsRPCClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *NatsRPCClient) configure(config config.NatsRPCClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// BroadcastSessionBind sends the binding information to other servers that may be interested in this info
func (ns *NatsRPCClient) BroadcastSessionBind(uid string) error {
	_ = "STUB: not implemented"
	return nil
}

// Send publishes a message in a given topic
func (ns *NatsRPCClient) Send(topic string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SendPush sends a message to a user
func (ns *NatsRPCClient) SendPush(userID string, frontendSv *Server, push *protos.Push) error {
	_ = "STUB: not implemented"
	return nil
}

// SendKick kicks an user
func (ns *NatsRPCClient) SendKick(userID string, serverType string, kick *protos.KickMsg) error {
	_ = "STUB: not implemented"
	return nil
}

// Call calls a method remotely
func (ns *NatsRPCClient) Call(
	ctx context.Context,
	rpcType protos.RPCType,
	route *route.Route,
	session session.Session,
	msg *message.Message,
	server *Server,
) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// replaceConnection replaces the NATS connection, draining the old one
func (ns *NatsRPCClient) replaceConnection() error { _ = "STUB: not implemented"; return nil }

// client doesn't have subscriptions

// Init inits nats rpc client
func (ns *NatsRPCClient) Init() error { _ = "STUB: not implemented"; return nil }

// initConnection initializes or replaces the NATS connection
func (ns *NatsRPCClient) initConnection(isReplacement bool) error {
	_ = "STUB: not implemented"
	return nil
}

// AfterInit runs after initialization
func (ns *NatsRPCClient) AfterInit() {
	_ = "STUB: not implemented"

	// BeforeShutdown runs before shutdown
	return
}

func (ns *NatsRPCClient) BeforeShutdown() {
	_ = "STUB: not implemented"

	// Shutdown stops nats rpc server
	return
}

func (ns *NatsRPCClient) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (ns *NatsRPCClient) stop() { _ = "STUB: not implemented"; return }

func (ns *NatsRPCClient) getSubscribeChannel() string { _ = "STUB: not implemented"; return "" }

func (ns *NatsRPCClient) IsConnected() bool { _ = "STUB: not implemented"; return false }
