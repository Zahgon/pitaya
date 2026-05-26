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
		Stats *protos.Stats
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

	// RPCResponse represents a rpc message
	RPCResponse struct {
		Msg string `json:"msg"`
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

// NewRoom returns a new room
func NewRoom(app pitaya.Pitaya) *Room { _ = "STUB: not implemented"; return nil }

// Init runs on service initialization
func (r *Room) Init() { _ = "STUB: not implemented"; return }

// AfterInit component lifetime callback
func (r *Room) AfterInit() { _ = "STUB: not implemented"; return }

// Entry is the entrypoint
func (r *Room) Entry(ctx context.Context, msg []byte) (*protos.JoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The default logger contains a requestId, the route being executed and the sessionId

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

// Notify push is a notify route that triggers a push to a session
func (r *Room) NotifyPush(ctx context.Context) { _ = "STUB: not implemented"; return }

// Join room
func (r *Room) Join(ctx context.Context) (*protos.JoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Leave room
func (r *Room) Leave(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Message sync last message to all members
func (r *Room) Message(ctx context.Context, msg *protos.UserMessage) {
	_ = "STUB: not implemented"
	return
}

// SendRPC sends rpc
func (r *Room) SendRPC(ctx context.Context, msg *protos.SendRPCMsg) (*protos.RPCRes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MessageRemote just echoes the given message
func (r *Room) MessageRemote(ctx context.Context, msg *protos.UserMessage, b bool, s string) (*protos.UserMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
