package services

import (
	"context"

	"github.com/topfreegames/pitaya/v3/examples/demo/protos"
	pitaya "github.com/topfreegames/pitaya/v3/pkg"
	"github.com/topfreegames/pitaya/v3/pkg/component"
)

// ConnectorRemote is a remote that will receive rpc's
type ConnectorRemote struct {
	component.Base
}

// Connector struct
type Connector struct {
	component.Base
	app pitaya.Pitaya
}

// SessionData struct
type SessionData struct {
	Data map[string]interface{}
}

// Response struct
type Response struct {
	Code int32
	Msg  string
}

// NewConnector ctor
func NewConnector(app pitaya.Pitaya) *Connector { _ = "STUB: not implemented"; return nil }

func reply(code int32, msg string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// GetSessionData gets the session data
func (c *Connector) GetSessionData(ctx context.Context) (*SessionData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetSessionData sets the session data
func (c *Connector) SetSessionData(ctx context.Context, data *SessionData) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NotifySessionData sets the session data
func (c *Connector) NotifySessionData(ctx context.Context, data *SessionData) {
	_ = "STUB: not implemented"
	return
}

// SendPushToUser sends a push to a user
func (c *Connector) SendPushToUser(ctx context.Context, msg *UserMessage) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoteFunc is a function that will be called remotely
func (c *ConnectorRemote) RemoteFunc(ctx context.Context, msg *protos.RPCMsg) (*protos.RPCRes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
