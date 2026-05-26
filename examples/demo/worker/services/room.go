package services

import (
	"context"

	"github.com/topfreegames/pitaya/v3/examples/demo/worker/protos"
	pitaya "github.com/topfreegames/pitaya/v3/pkg"
	"github.com/topfreegames/pitaya/v3/pkg/component"
)

// Room server
type Room struct {
	component.Base
	app pitaya.Pitaya
}

// NewRoom ctor
func NewRoom(app pitaya.Pitaya) *Room { _ = "STUB: not implemented"; return nil }

// CallLog makes ReliableRPC to metagame LogRemote
func (r *Room) CallLog(ctx context.Context, arg *protos.Arg) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
