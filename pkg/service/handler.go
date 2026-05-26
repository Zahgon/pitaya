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

package service

import (
	"context"

	"github.com/topfreegames/pitaya/v3/pkg/acceptor"
	"github.com/topfreegames/pitaya/v3/pkg/pipeline"

	"github.com/topfreegames/pitaya/v3/pkg/agent"
	"github.com/topfreegames/pitaya/v3/pkg/cluster"
	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/conn/codec"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/conn/packet"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/route"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
)

var (
	handlerType = "handler"
)

type (
	// HandlerService service
	HandlerService struct {
		baseService
		chLocalProcess   chan unhandledMessage // channel of messages that will be processed locally
		chRemoteProcess  chan unhandledMessage // channel of messages that will be processed remotely
		decoder          codec.PacketDecoder   // binary decoder
		remoteService    *RemoteService
		serializer       serialize.Serializer          // message serializer
		server           *cluster.Server               // server obj
		services         map[string]*component.Service // all registered service
		metricsReporters []metrics.Reporter
		agentFactory     agent.AgentFactory
		handlerPool      *HandlerPool
		handlers         map[string]*component.Handler // all handler method
	}

	unhandledMessage struct {
		ctx   context.Context
		agent agent.Agent
		route *route.Route
		msg   *message.Message
	}
)

// NewHandlerService creates and returns a new handler service
func NewHandlerService(
	packetDecoder codec.PacketDecoder,
	serializer serialize.Serializer,
	localProcessBufferSize int,
	remoteProcessBufferSize int,
	server *cluster.Server,
	remoteService *RemoteService,
	agentFactory agent.AgentFactory,
	metricsReporters []metrics.Reporter,
	handlerHooks *pipeline.HandlerHooks,
	handlerPool *HandlerPool,
) *HandlerService {
	_ = "STUB: not implemented"
	return nil
}

// Dispatch message to corresponding logic handler
func (h *HandlerService) Dispatch(thread int) {
	_ = "STUB: not implemented"
	// TODO: This timer is being stopped multiple times, it probably doesn't need to be stopped here
	return
}

// Calls to remote servers block calls to local server

// execute cron task

// new Timers

// closing Timers

// Register registers components
func (h *HandlerService) Register(comp component.Component, opts []component.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// register all handlers

// Handle handles messages from a conn
func (h *HandlerService) Handle(conn acceptor.PlayerConn) {
	_ = "STUB: not implemented"
	// create a client agent and startup write goroutine
	return
}

// startup agent goroutine

// guarantee agent related resource is destroyed

// Check if this is an expected error due to connection being closed

// Differentiate errors for valid sessions, to avoid noise from load balancer healthchecks and other internet noise

// process all packet

func (h *HandlerService) processPacket(a agent.Agent, p *packet.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse the json sent with the handshake by the client

// expected

func (h *HandlerService) processMessage(a agent.Agent, msg *message.Message) {
	_ = "STUB: not implemented"
	return
}

func (h *HandlerService) localProcess(ctx context.Context, a agent.Agent, route *route.Route, msg *message.Message) {
	_ = "STUB: not implemented"
	return
}

// DumpServices outputs all registered services
func (h *HandlerService) DumpServices() { _ = "STUB: not implemented"; return }

// Docs returns documentation for handlers
func (h *HandlerService) Docs(getPtrNames bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
