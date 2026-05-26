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
	"sync"
	"time"

	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/interfaces"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/route"
	"github.com/topfreegames/pitaya/v3/pkg/session"
	"google.golang.org/grpc"
)

// GRPCClient rpc client struct
type GRPCClient struct {
	bindingStorage   interfaces.BindingStorage
	clientMap        sync.Map
	dialTimeout      time.Duration
	infoRetriever    InfoRetriever
	lazy             bool
	metricsReporters []metrics.Reporter
	reqTimeout       time.Duration
	server           *Server
}

// NewGRPCClient returns a new instance of GRPCClient
func NewGRPCClient(
	config config.GRPCClientConfig,
	server *Server,
	metricsReporters []metrics.Reporter,
	bindingStorage interfaces.BindingStorage,
	infoRetriever InfoRetriever,
) (*GRPCClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type grpcClient struct {
	address   string
	cli       protos.PitayaClient
	conn      *grpc.ClientConn
	connected bool
	lock      sync.Mutex
}

// Init inits grpc rpc client
func (gs *GRPCClient) Init() error {
	_ = "STUB: not implemented"

	// Call makes a RPC Call
	return nil
}

func (gs *GRPCClient) Call(
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

// Send not implemented in grpc client
func (gs *GRPCClient) Send(uid string, d []byte) error { _ = "STUB: not implemented"; return nil }

// BroadcastSessionBind sends the binding information to other servers that may be interested in this info
func (gs *GRPCClient) BroadcastSessionBind(uid string) error { _ = "STUB: not implemented"; return nil }

// SendKick sends a kick to an user
func (gs *GRPCClient) SendKick(userID string, serverType string, kick *protos.KickMsg) error {
	_ = "STUB: not implemented"
	return nil
}

// SendPush sends a message to an user, if you dont know the serverID that the user is connected to, you need to set a BindingStorage when creating the client
// TODO: Jaeger?
func (gs *GRPCClient) SendPush(userID string, frontendSv *Server, push *protos.Push) error {
	_ = "STUB: not implemented"
	return nil
}

// AddServer is called when a new server is discovered
func (gs *GRPCClient) AddServer(sv *Server) { _ = "STUB: not implemented"; return }

// RemoveServer is called when a server is removed
func (gs *GRPCClient) RemoveServer(sv *Server) { _ = "STUB: not implemented"; return }

// AfterInit runs after initialization
func (gs *GRPCClient) AfterInit() {
	_ = "STUB: not implemented"

	// BeforeShutdown runs before shutdown
	return
}

func (gs *GRPCClient) BeforeShutdown() {
	_ = "STUB: not implemented"

	// Shutdown stops grpc rpc server
	return
}

func (gs *GRPCClient) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (gs *GRPCClient) getServerHost(sv *Server) (host, portKey string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (gc *grpcClient) connect() error { _ = "STUB: not implemented"; return nil }

func (gc *grpcClient) disconnect() { _ = "STUB: not implemented"; return }

func (gc *grpcClient) pushToUser(ctx context.Context, push *protos.Push) error {
	_ = "STUB: not implemented"
	return nil
}

func (gc *grpcClient) call(ctx context.Context, req *protos.Request) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gc *grpcClient) sessionBindRemote(ctx context.Context, req *protos.BindMsg) error {
	_ = "STUB: not implemented"
	return nil
}

func (gc *grpcClient) sendKick(ctx context.Context, req *protos.KickMsg) error {
	_ = "STUB: not implemented"
	return nil
}
