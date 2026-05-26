package services

import (
	"context"

	"github.com/topfreegames/pitaya/v3/pkg/component"
)

// Room represents a component that contains a bundle of room related handler
type Room struct {
	component.Base
}

// NewRoom returns a new room
func NewRoom() *Room {
	_ = "STUB: not implemented"

	// Ping returns a pong
	return nil
}

func (r *Room) Ping(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
