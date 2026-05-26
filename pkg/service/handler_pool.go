package service

import (
	"context"

	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/pipeline"
	"github.com/topfreegames/pitaya/v3/pkg/route"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
	"github.com/topfreegames/pitaya/v3/pkg/session"
)

// HandlerPool ...
type HandlerPool struct {
	handlers map[string]*component.Handler // all handler method
}

// NewHandlerPool ...
func NewHandlerPool() *HandlerPool { _ = "STUB: not implemented"; return nil }

// Register ...
func (h *HandlerPool) Register(serviceName string, name string, handler *component.Handler) {
	_ = "STUB: not implemented"
	return
}

// GetHandlers ...
func (h *HandlerPool) GetHandlers() map[string]*component.Handler {
	_ = "STUB: not implemented"

	// ProcessHandlerMessage ...
	return nil
}

func (h *HandlerPool) ProcessHandlerMessage(
	ctx context.Context,
	rt *route.Route,
	serializer serialize.Serializer,
	handlerHooks *pipeline.HandlerHooks,
	session session.Session,
	data []byte,
	msgTypeIface interface{},
	remote bool,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First unmarshal the handler arg that will be passed to
// both handler and pipeline functions

// This is a special case and should only happen with nats rpc client
// because we used nats request we have to answer to it or else a timeout
// will happen in the caller server and will be returned to the client
// the reason why we don't just Publish is to keep track of failed rpc requests
// with timeouts, maybe we can improve this flow

func (h *HandlerPool) getHandler(rt *route.Route) (*component.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
