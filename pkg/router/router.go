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

package router

import (
	"context"

	"github.com/topfreegames/pitaya/v3/pkg/cluster"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/route"
)

// Router struct
type Router struct {
	serviceDiscovery cluster.ServiceDiscovery
	routesMap        map[string]RoutingFunc
}

// RoutingFunc defines a routing function
type RoutingFunc func(
	ctx context.Context,
	route *route.Route,
	payload []byte,
	servers map[string]*cluster.Server,
) (*cluster.Server, error)

// New returns the router
func New() *Router { _ = "STUB: not implemented"; return nil }

// SetServiceDiscovery sets the sd client
func (r *Router) SetServiceDiscovery(sd cluster.ServiceDiscovery) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) defaultRoute(
	servers map[string]*cluster.Server,
) *cluster.Server {
	_ = "STUB: not implemented"
	return nil
}

// Route gets the right server to use in the call
func (r *Router) Route(
	ctx context.Context,
	rpcType protos.RPCType,
	svType string,
	route *route.Route,
	msg *message.Message,
) (*cluster.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddRoute adds a routing function to a server type
func (r *Router) AddRoute(
	serverType string,
	routingFunction RoutingFunc,
) {
	_ = "STUB: not implemented"
	return
}
