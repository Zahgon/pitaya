package services

import (
	"context"

	"github.com/golang/protobuf/proto"
	pitaya "github.com/topfreegames/pitaya/v3/pkg"
	"github.com/topfreegames/pitaya/v3/pkg/component"
)

// Worker server
type Worker struct {
	component.Base
}

// Configure starts workers and register rpc job
func (w *Worker) Configure(app pitaya.Pitaya) { _ = "STUB: not implemented"; return }

// RPCJob implements worker.RPCJob
type RPCJob struct {
	app pitaya.Pitaya
}

// ServerDiscovery returns a serverID="", meaning any server
// is ok
func (r *RPCJob) ServerDiscovery(
	route string,
	rpcMetadata map[string]interface{},
) (serverID string, err error) {
	_ = "STUB: not implemented"

	// RPC calls pitaya's rpc
	return "", nil
}

func (r *RPCJob) RPC(
	ctx context.Context,
	serverID, routeStr string,
	reply, arg proto.Message,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetArgReply returns reply and arg of LogRemote,
// since we have no other methods in this example
func (r *RPCJob) GetArgReply(
	route string,
) (arg, reply proto.Message, err error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), *new(proto.Message), nil
}
