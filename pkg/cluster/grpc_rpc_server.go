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
	"google.golang.org/grpc"

	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
)

// GRPCServer rpc server struct
type GRPCServer struct {
	server           *Server
	port             int
	metricsReporters []metrics.Reporter
	grpcSv           *grpc.Server
	pitayaServer     protos.PitayaServer
}

// NewGRPCServer constructor
func NewGRPCServer(config config.GRPCServerConfig, server *Server, metricsReporters []metrics.Reporter) (*GRPCServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Init inits grpc rpc server
func (gs *GRPCServer) Init() error { _ = "STUB: not implemented"; return nil }

// SetPitayaServer sets the pitaya server
func (gs *GRPCServer) SetPitayaServer(ps protos.PitayaServer) { _ = "STUB: not implemented"; return }

// AfterInit runs after initialization
func (gs *GRPCServer) AfterInit() {
	_ = "STUB: not implemented"

	// BeforeShutdown runs before shutdown
	return
}

func (gs *GRPCServer) BeforeShutdown() {
	_ = "STUB: not implemented"

	// Shutdown stops grpc rpc server
	return
}

func (gs *GRPCServer) Shutdown() error {
	_ = "STUB: not implemented"
	// graceful: stops the server from accepting new connections and RPCs and
	// blocks until all the pending RPCs are finished.
	// source: https://godoc.org/google.golang.org/grpc#Server.GracefulStop
	return nil
}
