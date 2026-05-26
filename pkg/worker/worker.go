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

package worker

import (
	"github.com/golang/protobuf/proto"
	workers "github.com/topfreegames/go-workers"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/logger/interfaces"
)

// Worker executes RPCs with retry and backoff time
type Worker struct {
	concurrency int
	registered  bool
	opts        *config.EnqueueOpts
	started     bool
}

// NewWorker configures and returns a *Worker
func NewWorker(config config.WorkerConfig, opts config.EnqueueOpts) (*Worker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetLogger overwrites worker logger
func (w *Worker) SetLogger(logger interfaces.Logger) { _ = "STUB: not implemented"; return }

// Start starts worker in another gorotine
func (w *Worker) Start() { _ = "STUB: not implemented"; return }

// Started returns true if worker was started
func (w *Worker) Started() bool { _ = "STUB: not implemented"; return false }

// EnqueueRPC enqueues rpc job to worker
func (w *Worker) EnqueueRPC(
	routeStr string,
	metadata map[string]interface{},
	reply, arg proto.Message,
) (jid string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// EnqueueRPCWithOptions enqueues rpc job to worker
func (w *Worker) EnqueueRPCWithOptions(
	routeStr string,
	metadata map[string]interface{},
	reply, arg proto.Message,
	opts *config.EnqueueOpts,
) (jid string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RegisterRPCJob registers a RPC job
func (w *Worker) RegisterRPCJob(rpcJob RPCJob) error { _ = "STUB: not implemented"; return nil }

func (w *Worker) parsedRPCJob(rpcJob RPCJob) func(*workers.Msg) {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worker) enqueueOptions(
	opts *config.EnqueueOpts,
) workers.EnqueueOptions {
	_ = "STUB: not implemented"
	return *new(workers.EnqueueOptions)
}

func (w *Worker) unmarshalRouteMetadata(
	jobArg *workers.Msg,
) ([]byte, *rpcRoute, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
