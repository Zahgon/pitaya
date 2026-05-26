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
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/session"
)

// NatsRPCServer struct
type NatsRPCServer struct {
	service                int
	connString             string
	connectionTimeout      time.Duration
	maxReconnectionRetries int
	server                 *Server
	conn                   *nats.Conn
	pushBufferSize         int
	messagesBufferSize     int
	stopChan               chan bool
	subChan                chan *nats.Msg // subChan is the channel used by the server to receive network messages addressed to itself
	bindingsChan           chan *nats.Msg // bindingsChan receives notify from other servers on every user bind to session
	unhandledReqCh         chan *protos.Request
	responses              []*protos.Response
	requests               []*protos.Request
	userPushCh             chan *protos.Push
	userKickCh             chan *protos.KickMsg
	sub                    *nats.Subscription
	dropped                int
	pitayaServer           protos.PitayaServer
	metricsReporters       []metrics.Reporter
	sessionPool            session.SessionPool
	appDieChan             chan bool
	websocketCompression   bool
	reconnectJitter        time.Duration
	reconnectJitterTLS     time.Duration
	reconnectWait          time.Duration
	pingInterval           time.Duration
	maxPingsOutstanding    int
}

// NewNatsRPCServer ctor
func NewNatsRPCServer(
	config config.NatsRPCServerConfig,
	server *Server,
	metricsReporters []metrics.Reporter,
	appDieChan chan bool,
	sessionPool session.SessionPool,
) (*NatsRPCServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *NatsRPCServer) configure(config config.NatsRPCServerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// the reason this channel is buffered is that we can achieve more performance by not
// blocking producers on a massive push

// GetBindingsChannel gets the channel that will receive all bindings
func (ns *NatsRPCServer) GetBindingsChannel() chan *nats.Msg { _ = "STUB: not implemented"; return nil }

// GetUserMessagesTopic get the topic for user
func GetUserMessagesTopic(uid string, svType string) string { _ = "STUB: not implemented"; return "" }

// GetUserKickTopic get the topic for kicking an user
func GetUserKickTopic(uid string, svType string) string { _ = "STUB: not implemented"; return "" }

// GetBindBroadcastTopic gets the topic on which bind events will be broadcasted
func GetBindBroadcastTopic(svType string) string { _ = "STUB: not implemented"; return "" }

// onSessionBind should be called on each session bind
func (ns *NatsRPCServer) onSessionBind(ctx context.Context, s session.Session) error {
	_ = "STUB: not implemented"
	return nil
}

// SetPitayaServer sets the pitaya server
func (ns *NatsRPCServer) SetPitayaServer(ps protos.PitayaServer) { _ = "STUB: not implemented"; return }

func (ns *NatsRPCServer) subscribeToBindingsChannel() error { _ = "STUB: not implemented"; return nil }

func (ns *NatsRPCServer) subscribeToUserKickChannel(uid string, svType string) (*nats.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *NatsRPCServer) subscribeToUserMessages(uid string, svType string) (*nats.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *NatsRPCServer) handleMessages() { _ = "STUB: not implemented"; return }

// Check if subscription is still valid before accessing it
// This can happen during connection replacement when the old subscription becomes invalid

// TODO: Add tracing here to report delay to start processing message in spans

// should answer rpc with an error

// GetUnhandledRequestsChannel gets the unhandled requests channel from nats rpc server
func (ns *NatsRPCServer) GetUnhandledRequestsChannel() chan *protos.Request {
	_ = "STUB: not implemented"
	return nil
}

func (ns *NatsRPCServer) getUserPushChannel() chan *protos.Push {
	_ = "STUB: not implemented"
	return nil
}

func (ns *NatsRPCServer) getUserKickChannel() chan *protos.KickMsg {
	_ = "STUB: not implemented"
	return nil
}

func (ns *NatsRPCServer) marshalResponse(res *protos.Response) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *NatsRPCServer) processMessages(threadID int) { _ = "STUB: not implemented"; return }

func (ns *NatsRPCServer) processSessionBindings() { _ = "STUB: not implemented"; return }

func (ns *NatsRPCServer) processPushes() { _ = "STUB: not implemented"; return }

func (ns *NatsRPCServer) processKick() { _ = "STUB: not implemented"; return }

// replaceConnection replaces the NATS connection, draining the old one and re-subscribing
func (ns *NatsRPCServer) replaceConnection() error { _ = "STUB: not implemented"; return nil }

// Init inits nats rpc server
func (ns *NatsRPCServer) Init() error { _ = "STUB: not implemented"; return nil }

// initConnection initializes or replaces the NATS connection
func (ns *NatsRPCServer) initConnection(isReplacement bool) error {
	_ = "STUB: not implemented"
	return nil

	// TODO should we have concurrency here? it feels like we should
}

// Re-subscribe to all session subscriptions if this is a replacement
// The onSessionBind callback is already set up, we just need to trigger it for existing sessions

// Re-use the same subscription logic as onSessionBind

// this handles remote messages

// this should be so fast that we shoudn't need concurrency

// AfterInit runs after initialization
func (ns *NatsRPCServer) AfterInit() {
	_ = "STUB: not implemented"

	// BeforeShutdown runs before shutdown
	return
}

func (ns *NatsRPCServer) BeforeShutdown() {
	_ = "STUB: not implemented"

	// Shutdown stops nats rpc server
	return
}

func (ns *NatsRPCServer) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (ns *NatsRPCServer) subscribe(topic string) (*nats.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ns *NatsRPCServer) stop() { _ = "STUB: not implemented"; return }

func (ns *NatsRPCServer) reportMetrics() { _ = "STUB: not implemented"; return }

// subchan

// bindingschan

// userpushch

// IsConnected returns true if NATS connection is established
func (ns *NatsRPCServer) IsConnected() bool { _ = "STUB: not implemented"; return false }
