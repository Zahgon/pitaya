// Copyright (c) nano Author and TFG Co. All Rights Reserved.
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

package session

import (
	"context"
	"net"
	"sync"

	nats "github.com/nats-io/nats.go"
	"github.com/topfreegames/pitaya/v3/pkg/logger/interfaces"
	"github.com/topfreegames/pitaya/v3/pkg/networkentity"
)

type sessionPoolImpl struct {
	sessionBindCallbacks []func(ctx context.Context, s Session) error
	afterBindCallbacks   []func(ctx context.Context, s Session) error
	handshakeValidators  map[string]func(data *HandshakeData) error

	// SessionCloseCallbacks contains global session close callbacks
	SessionCloseCallbacks []func(s Session)
	sessionsByUID         sync.Map
	sessionsByID          sync.Map
	sessionIDSvc          *sessionIDService
	// SessionCount keeps the current number of sessions
	SessionCount int64
}

// SessionPool centralizes all sessions within a Pitaya app
type SessionPool interface {
	NewSession(entity networkentity.NetworkEntity, frontend bool, UID ...string) Session
	GetSessionCount() int64
	GetSessionCloseCallbacks() []func(s Session)
	GetSessionByUID(uid string) Session
	GetSessionByID(id int64) Session
	OnSessionBind(f func(ctx context.Context, s Session) error)
	OnAfterSessionBind(f func(ctx context.Context, s Session) error)
	OnSessionClose(f func(s Session))
	CloseAll()
	AddHandshakeValidator(name string, f func(data *HandshakeData) error)
	GetNumberOfConnectedClients() int64
	ForEachSession(f func(s Session))
}

// HandshakeClientData represents information about the client sent on the handshake.
type HandshakeClientData struct {
	Platform    string `json:"platform"`
	LibVersion  string `json:"libVersion"`
	BuildNumber string `json:"clientBuildNumber"`
	Version     string `json:"clientVersion"`
}

// HandshakeData represents information about the handshake sent by the client.
// `sys` corresponds to information independent from the app and `user` information
// that depends on the app and is customized by the user.
type HandshakeData struct {
	Sys  HandshakeClientData    `json:"sys"`
	User map[string]interface{} `json:"user,omitempty"`
}

type sessionImpl struct {
	sync.RWMutex                                              // protect data
	id                  int64                                 // session global unique id
	uid                 string                                // binding user id
	lastTime            int64                                 // last heartbeat time
	entity              networkentity.NetworkEntity           // low-level network entity
	data                map[string]interface{}                // session data store
	handshakeData       *HandshakeData                        // handshake data received by the client
	handshakeValidators map[string]func(*HandshakeData) error // validations to run on handshake
	encodedData         []byte                                // session data encoded as a byte array
	OnCloseCallbacks    []func()                              //onClose callbacks
	IsFrontend          bool                                  // if session is a frontend session
	frontendID          string                                // the id of the frontend that owns the session
	frontendSessionID   int64                                 // the id of the session on the frontend server
	Subscriptions       []*nats.Subscription                  // subscription created on bind when using nats rpc server
	requestsInFlight    ReqInFlight                           // whether the session is waiting from a response from a remote
	pool                *sessionPoolImpl
	logger              interfaces.Logger // logger instance for this session
}

type ReqInFlight struct {
	m  map[string]string
	mu sync.RWMutex
}

// Session represents a client session, which can store data during the connection.
// All data is released when the low-level connection is broken.
// Session instance related to the client will be passed to Handler method in the
// context parameter.
type Session interface {
	GetOnCloseCallbacks() []func()
	GetIsFrontend() bool
	GetSubscriptions() []*nats.Subscription
	SetOnCloseCallbacks(callbacks []func())
	SetIsFrontend(isFrontend bool)
	SetSubscriptions(subscriptions []*nats.Subscription)
	HasRequestsInFlight() bool
	GetRequestsInFlight() ReqInFlight
	SetRequestInFlight(reqID string, reqData string, inFlight bool)

	Push(route string, v interface{}) error
	ResponseMID(ctx context.Context, mid uint, v interface{}, err ...bool) error
	ID() int64
	UID() string
	GetData() map[string]interface{}
	SetData(data map[string]interface{}) error
	GetDataEncoded() []byte
	SetDataEncoded(encodedData []byte) error
	SetFrontendData(frontendID string, frontendSessionID int64)
	Bind(ctx context.Context, uid string) error
	Kick(ctx context.Context) error
	OnClose(c func()) error
	Close()
	RemoteAddr() net.Addr
	Remove(key string) error
	Set(key string, value interface{}) error
	HasKey(key string) bool
	Get(key string) interface{}
	Int(key string) int
	Int8(key string) int8
	Int16(key string) int16
	Int32(key string) int32
	Int64(key string) int64
	Uint(key string) uint
	Uint8(key string) uint8
	Uint16(key string) uint16
	Uint32(key string) uint32
	Uint64(key string) uint64
	Float32(key string) float32
	Float64(key string) float64
	String(key string) string
	Value(key string) interface{}
	PushToFront(ctx context.Context) error
	Clear()
	SetHandshakeData(data *HandshakeData)
	GetHandshakeData() *HandshakeData
	ValidateHandshake(data *HandshakeData) error
	GetHandshakeValidators() map[string]func(data *HandshakeData) error
}

type sessionIDService struct {
	sid int64
}

func newSessionIDService() *sessionIDService { _ = "STUB: not implemented"; return nil }

// SessionID returns the session id
func (c *sessionIDService) sessionID() int64 { _ = "STUB: not implemented"; return 0 }

// NewSession returns a new session instance
// a networkentity.NetworkEntity is a low-level network instance
func (pool *sessionPoolImpl) NewSession(entity networkentity.NetworkEntity, frontend bool, UID ...string) Session {
	_ = "STUB: not implemented"
	return *new(Session)
}

// NewSessionPool returns a new session pool instance
func NewSessionPool() SessionPool { _ = "STUB: not implemented"; return *new(SessionPool) }

func (pool *sessionPoolImpl) GetSessionCount() int64 { _ = "STUB: not implemented"; return 0 }

func (pool *sessionPoolImpl) GetSessionCloseCallbacks() []func(s Session) {
	_ = "STUB: not implemented"
	return nil
}

// GetSessionByUID return a session bound to an user id
func (pool *sessionPoolImpl) GetSessionByUID(uid string) Session {
	_ = "STUB: not implemented"
	// TODO: Block this operation in backend servers
	return *new(Session)
}

// GetSessionByID return a session bound to a frontend server id
func (pool *sessionPoolImpl) GetSessionByID(id int64) Session {
	_ = "STUB: not implemented"
	// TODO: Block this operation in backend servers
	return *new(Session)
}

// OnSessionBind adds a method to be called when a session is bound
// same function cannot be added twice!
func (pool *sessionPoolImpl) OnSessionBind(f func(ctx context.Context, s Session) error) {
	_ = "STUB: not implemented"
	// Prevents the same function to be added twice in onSessionBind
	return
}

// OnAfterSessionBind adds a method to be called when session is bound and after all sessionBind callbacks
func (pool *sessionPoolImpl) OnAfterSessionBind(f func(ctx context.Context, s Session) error) {
	_ = "STUB: not implemented"
	// Prevents the same function to be added twice in onSessionBind
	return
}

// OnSessionClose adds a method that will be called when every session closes
func (pool *sessionPoolImpl) OnSessionClose(f func(s Session)) { _ = "STUB: not implemented"; return }

// CloseAll calls Close on all sessions
func (pool *sessionPoolImpl) CloseAll() { _ = "STUB: not implemented"; return }

// AddHandshakeValidator allows adds validation functions that will run when
// handshake packets are processed. Errors will be raised with the given name.
func (pool *sessionPoolImpl) AddHandshakeValidator(name string, f func(data *HandshakeData) error) {
	_ = "STUB: not implemented"
	return
}

// GetNumberOfConnectedClients returns the number of connected clients
func (pool *sessionPoolImpl) GetNumberOfConnectedClients() int64 {
	_ = "STUB: not implemented"
	return 0
}

// ForEachSession iterates through all sessions in the pool and calls f for each one
func (pool *sessionPoolImpl) ForEachSession(f func(s Session)) { _ = "STUB: not implemented"; return }

func (s *sessionImpl) updateEncodedData() error { _ = "STUB: not implemented"; return nil }

// GetOnCloseCallbacks ...
func (s *sessionImpl) GetOnCloseCallbacks() []func() { _ = "STUB: not implemented"; return nil }

// GetIsFrontend ...
func (s *sessionImpl) GetIsFrontend() bool { _ = "STUB: not implemented"; return false }

// GetSubscriptions ...
func (s *sessionImpl) GetSubscriptions() []*nats.Subscription {
	_ = "STUB: not implemented"
	return nil

	// SetOnCloseCallbacks ...
}

func (s *sessionImpl) SetOnCloseCallbacks(callbacks []func()) { _ = "STUB: not implemented"; return }

// SetIsFrontend ...
func (s *sessionImpl) SetIsFrontend(isFrontend bool) { _ = "STUB: not implemented"; return }

// SetSubscriptions ...
func (s *sessionImpl) SetSubscriptions(subscriptions []*nats.Subscription) {
	_ = "STUB: not implemented"
	return
}

// Push message to client
func (s *sessionImpl) Push(route string, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// ResponseMID responses message to client, mid is
// request message ID
func (s *sessionImpl) ResponseMID(ctx context.Context, mid uint, v interface{}, err ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ID returns the session id
func (s *sessionImpl) ID() int64 {
	_ = "STUB: not implemented"

	// UID returns uid that bind to current session
	return 0
}

func (s *sessionImpl) UID() string {
	_ = "STUB: not implemented"

	// GetData gets the data
	return ""
}

func (s *sessionImpl) GetData() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// SetData sets the whole session data
func (s *sessionImpl) SetData(data map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDataEncoded returns the session data as an encoded value
func (s *sessionImpl) GetDataEncoded() []byte { _ = "STUB: not implemented"; return nil }

// SetDataEncoded sets the whole session data from an encoded value
func (s *sessionImpl) SetDataEncoded(encodedData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SetFrontendData sets frontend id and session id
func (s *sessionImpl) SetFrontendData(frontendID string, frontendSessionID int64) {
	_ = "STUB: not implemented"
	return
}

// Bind bind UID to current session
func (s *sessionImpl) Bind(ctx context.Context, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

// if code running on frontend server

// If a session with the same UID already exists in this frontend server, close it

// If frontentID is set this means it is a remote call and the current server
// is not the frontend server that received the user request

// invoke after callbacks on session bound

// Kick kicks the user
func (s *sessionImpl) Kick(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// OnClose adds the function it receives to the callbacks that will be called
// when the session is closed
func (s *sessionImpl) OnClose(c func()) error { _ = "STUB: not implemented"; return nil }

// Close terminates current session, session related data will not be released,
// all related data should be cleared explicitly in Session closed callback
func (s *sessionImpl) Close() { _ = "STUB: not implemented"; return }

// Only remove session by UID if the session ID matches the one being closed. This avoids problems with removing a valid session after the user has already reconnected before this session's heartbeat times out

// TODO: this logic should be moved to nats rpc server

// if the user is bound to an userid and nats rpc server is being used we need to unsubscribe

// RemoteAddr returns the remote network address.
func (s *sessionImpl) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// Remove delete data associated with the key from session storage
func (s *sessionImpl) Remove(key string) error { _ = "STUB: not implemented"; return nil }

// Set associates value with the key in session storage
func (s *sessionImpl) Set(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// HasKey decides whether a key has associated value
func (s *sessionImpl) HasKey(key string) bool { _ = "STUB: not implemented"; return false }

// Get returns a key value
func (s *sessionImpl) Get(key string) interface{} { _ = "STUB: not implemented"; return nil }

// Int returns the value associated with the key as a int.
func (s *sessionImpl) Int(key string) int { _ = "STUB: not implemented"; return 0 }

// Int8 returns the value associated with the key as a int8.
func (s *sessionImpl) Int8(key string) int8 { _ = "STUB: not implemented"; return 0 }

// Int16 returns the value associated with the key as a int16.
func (s *sessionImpl) Int16(key string) int16 { _ = "STUB: not implemented"; return 0 }

// Int32 returns the value associated with the key as a int32.
func (s *sessionImpl) Int32(key string) int32 { _ = "STUB: not implemented"; return 0 }

// Int64 returns the value associated with the key as a int64.
func (s *sessionImpl) Int64(key string) int64 { _ = "STUB: not implemented"; return 0 }

// Uint returns the value associated with the key as a uint.
func (s *sessionImpl) Uint(key string) uint { _ = "STUB: not implemented"; return 0 }

// Uint8 returns the value associated with the key as a uint8.
func (s *sessionImpl) Uint8(key string) uint8 { _ = "STUB: not implemented"; return 0 }

// Uint16 returns the value associated with the key as a uint16.
func (s *sessionImpl) Uint16(key string) uint16 { _ = "STUB: not implemented"; return 0 }

// Uint32 returns the value associated with the key as a uint32.
func (s *sessionImpl) Uint32(key string) uint32 { _ = "STUB: not implemented"; return 0 }

// Uint64 returns the value associated with the key as a uint64.
func (s *sessionImpl) Uint64(key string) uint64 { _ = "STUB: not implemented"; return 0 }

// Float32 returns the value associated with the key as a float32.
func (s *sessionImpl) Float32(key string) float32 { _ = "STUB: not implemented"; return 0 }

// Float64 returns the value associated with the key as a float64.
func (s *sessionImpl) Float64(key string) float64 { _ = "STUB: not implemented"; return 0 }

// String returns the value associated with the key as a string.
func (s *sessionImpl) String(key string) string { _ = "STUB: not implemented"; return "" }

// Value returns the value associated with the key as a interface{}.
func (s *sessionImpl) Value(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (s *sessionImpl) bindInFront(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// PushToFront updates the session in the frontend
func (s *sessionImpl) PushToFront(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Clear releases all data related to current session
func (s *sessionImpl) Clear() { _ = "STUB: not implemented"; return }

// SetHandshakeData sets the handshake data received by the client.
func (s *sessionImpl) SetHandshakeData(data *HandshakeData) { _ = "STUB: not implemented"; return }

// GetHandshakeData gets the handshake data received by the client.
func (s *sessionImpl) GetHandshakeData() *HandshakeData { _ = "STUB: not implemented"; return nil }

// GetHandshakeValidators return the handshake validators associated with the session.
func (s *sessionImpl) GetHandshakeValidators() map[string]func(data *HandshakeData) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionImpl) ValidateHandshake(data *HandshakeData) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionImpl) sendRequestToFront(ctx context.Context, route string, includeData bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionImpl) HasRequestsInFlight() bool { _ = "STUB: not implemented"; return false }

func (s *sessionImpl) GetRequestsInFlight() ReqInFlight {
	_ = "STUB: not implemented"
	return *new(ReqInFlight)
}

func (s *sessionImpl) SetRequestInFlight(reqID string, reqData string, inFlight bool) {
	_ = "STUB: not implemented"
	return
}
