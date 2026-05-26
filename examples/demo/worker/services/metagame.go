package services

import (
	"context"

	"github.com/topfreegames/pitaya/v3/examples/demo/worker/protos"
	"github.com/topfreegames/pitaya/v3/pkg/component"
)

// Metagame server
type Metagame struct {
	component.Base
}

// LogRemote logs argument when called
func (m *Metagame) LogRemote(ctx context.Context, arg *protos.Arg) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
