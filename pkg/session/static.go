package session

import (
	"context"
)

var DefaultSessionPool SessionPool

// GetSessionByUID return a session bound to an user id
func GetSessionByUID(uid string) Session { _ = "STUB: not implemented"; return *new(Session) }

// GetSessionByID return a session bound to a frontend server id
func GetSessionByID(id int64) Session { _ = "STUB: not implemented"; return *new(Session) }

// OnSessionBind adds a method to be called when a session is bound
// same function cannot be added twice!
func OnSessionBind(f func(ctx context.Context, s Session) error) { _ = "STUB: not implemented"; return }

// OnAfterSessionBind adds a method to be called when session is bound and after all sessionBind callbacks
func OnAfterSessionBind(f func(ctx context.Context, s Session) error) {
	_ = "STUB: not implemented"
	return
}

// OnSessionClose adds a method that will be called when every session closes
func OnSessionClose(f func(s Session)) { _ = "STUB: not implemented"; return }

// CloseAll calls Close on all sessions
func CloseAll() { _ = "STUB: not implemented"; return }

// GetNumberOfConnectedClients returns the number of connected clients
func GetNumberOfConnectedClients() int64 { _ = "STUB: not implemented"; return 0 }
