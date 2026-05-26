//
// Copyright (c) TFG Co. All Rights Reserved.
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

	"github.com/golang/protobuf/proto"

	"github.com/topfreegames/pitaya/v3/pkg/agent"
	"github.com/topfreegames/pitaya/v3/pkg/cluster"
	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/conn/codec"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/pipeline"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/route"
	"github.com/topfreegames/pitaya/v3/pkg/router"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
	"github.com/topfreegames/pitaya/v3/pkg/session"
)

// RemoteService struct
type RemoteService struct {
	baseService
	rpcServer              cluster.RPCServer
	serviceDiscovery       cluster.ServiceDiscovery
	serializer             serialize.Serializer
	encoder                codec.PacketEncoder
	rpcClient              cluster.RPCClient
	services               map[string]*component.Service // all registered service
	router                 *router.Router
	messageEncoder         message.Encoder
	server                 *cluster.Server // server obj
	remoteBindingListeners []cluster.RemoteBindingListener
	remoteHooks            *pipeline.RemoteHooks
	sessionPool            session.SessionPool
	handlerPool            *HandlerPool
	remotes                map[string]*component.Remote // all remote method
}

// NewRemoteService creates and return a new RemoteService
func NewRemoteService(
	rpcClient cluster.RPCClient,
	rpcServer cluster.RPCServer,
	sd cluster.ServiceDiscovery,
	encoder codec.PacketEncoder,
	serializer serialize.Serializer,
	router *router.Router,
	messageEncoder message.Encoder,
	server *cluster.Server,
	sessionPool session.SessionPool,
	remoteHooks *pipeline.RemoteHooks,
	handlerHooks *pipeline.HandlerHooks,
	handlerPool *HandlerPool,
) *RemoteService {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoteService) remoteProcess(
	ctx context.Context,
	server *cluster.Server,
	a agent.Agent,
	route *route.Route,
	msg *message.Message,
) {
	_ = "STUB: not implemented"
	return
}

// AddRemoteBindingListener adds a listener
func (r *RemoteService) AddRemoteBindingListener(bindingListener cluster.RemoteBindingListener) {
	_ = "STUB: not implemented"
	return
}

// Call processes a remote call
func (r *RemoteService) Call(ctx context.Context, req *protos.Request) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SessionBindRemote is called when a remote server binds a user session and want us to acknowledge it
func (r *RemoteService) SessionBindRemote(ctx context.Context, msg *protos.BindMsg) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushToUser sends a push to user
func (r *RemoteService) PushToUser(ctx context.Context, push *protos.Push) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KickUser sends a kick to user
func (r *RemoteService) KickUser(ctx context.Context, kick *protos.KickMsg) (*protos.KickAnswer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DoRPC do rpc and get answer
func (r *RemoteService) DoRPC(ctx context.Context, serverID string, route *route.Route, protoData []byte) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RPC makes rpcs
func (r *RemoteService) RPC(ctx context.Context, serverID string, route *route.Route, reply proto.Message, arg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Register registers components
func (r *RemoteService) Register(comp component.Component, opts []component.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// register all remotes

func processRemoteMessage(ctx context.Context, req *protos.Request, r *RemoteService) *protos.Response {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoteService) handleRPCUser(ctx context.Context, req *protos.Request, rt *route.Route) *protos.Response {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoteService) handleRPCSys(ctx context.Context, req *protos.Request, rt *route.Route) *protos.Response {
	_ = "STUB: not implemented"
	return nil
}

// (warning) a new agent is created for every new request

func (r *RemoteService) remoteCall(
	ctx context.Context,
	server *cluster.Server,
	rpcType protos.RPCType,
	route *route.Route,
	session session.Session,
	msg *message.Message,
) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DumpServices outputs all registered services
func (r *RemoteService) DumpServices() { _ = "STUB: not implemented"; return }

// Docs returns documentation for remotes
func (r *RemoteService) Docs(getPtrNames bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
