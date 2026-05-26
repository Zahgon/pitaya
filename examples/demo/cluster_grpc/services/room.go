package services

import (
	"context"

	"github.com/topfreegames/pitaya/v3/examples/demo/protos"
	pitaya "github.com/topfreegames/pitaya/v3/pkg"
	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/timer"
)

type (
	// Room represents a component that contains a bundle of room related handler
	// like Join/Message
	Room struct {
		component.Base
		timer *timer.Timer
		app   pitaya.Pitaya
		Stats *Stats
	}

	// UserMessage represents a message that user sent
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	// Stats exports the room status
	Stats struct {
		outboundBytes int
		inboundBytes  int
	}

	// SendRPCMsg represents a rpc message
	SendRPCMsg struct {
		ServerID string `json:"serverId"`
		Route    string `json:"route"`
		Msg      string `json:"msg"`
	}

	// NewUser message will be received when new user join room
	NewUser struct {
		Content string `json:"content"`
	}

	// AllMembers contains all members uid
	AllMembers struct {
		Members []string `json:"members"`
	}

	// JoinResponse represents the result of joining room
	JoinResponse struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}
)

// Outbound gets the outbound status
func (Stats *Stats) Outbound(ctx context.Context, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Inbound gets the inbound status
func (Stats *Stats) Inbound(ctx context.Context, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewRoom returns a new room
func NewRoom(app pitaya.Pitaya) *Room { _ = "STUB: not implemented"; return nil }

// Init runs on service initialization
func (r *Room) Init() { _ = "STUB: not implemented"; return }

// It is necessary to register all structs that will be used in RPC calls
// This must be done both in the caller and callee servers

// AfterInit component lifetime callback
func (r *Room) AfterInit() { _ = "STUB: not implemented"; return }

// Entry is the entrypoint
func (r *Room) Entry(ctx context.Context, msg []byte) (*JoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSessionData gets the session data
func (r *Room) GetSessionData(ctx context.Context) (*SessionData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetSessionData sets the session data
func (r *Room) SetSessionData(ctx context.Context, data *SessionData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Join room
func (r *Room) Join(ctx context.Context) (*JoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Message sync last message to all members
func (r *Room) Message(ctx context.Context, msg *UserMessage) { _ = "STUB: not implemented"; return }

// SendRPC sends rpc
func (r *Room) SendRPC(ctx context.Context, msg *protos.RPCMsg) (*protos.RPCRes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MessageRemote just echoes the given message
func (r *Room) MessageRemote(ctx context.Context, msg *UserMessage, b bool, s string) (*UserMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
