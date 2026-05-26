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

package agent

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/topfreegames/pitaya/v3/pkg/conn/codec"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/logger/interfaces"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
	"github.com/topfreegames/pitaya/v3/pkg/session"

	"go.opentelemetry.io/otel/trace"
)

var (
	// hbd contains the heartbeat packet data
	hbd []byte
	// hrd contains the handshake response data
	hrd []byte
	// herd contains the handshake error response data
	herd []byte
	once sync.Once
)

const handlerType = "handler"

type (
	agentImpl struct {
		Session            session.Session // session
		sessionPool        session.SessionPool
		appDieChan         chan bool         // app die channel
		chDie              chan struct{}     // wait for close
		chSend             chan pendingWrite // push message queue
		chStopHeartbeat    chan struct{}     // stop heartbeats
		chStopWrite        chan struct{}     // stop writing messages
		closeMutex         sync.Mutex
		conn               net.Conn            // low-level conn fd
		decoder            codec.PacketDecoder // binary decoder
		encoder            codec.PacketEncoder // binary encoder
		heartbeatTimeout   time.Duration
		writeTimeout       time.Duration
		lastAt             int64 // last heartbeat unix time stamp
		messageEncoder     message.Encoder
		messagesBufferSize int // size of the pending messages buffer
		metricsReporters   []metrics.Reporter
		serializer         serialize.Serializer // message serializer
		state              int32                // current agent state
		logger             interfaces.Logger
	}

	pendingMessage struct {
		ctx     context.Context
		typ     message.Type // message type
		route   string       // message route (push)
		mid     uint         // response message id (response)
		payload interface{}  // payload
		err     bool         // if its an error message
	}

	pendingWrite struct {
		ctx  context.Context
		data []byte
		err  error
	}

	// Agent corresponds to a user and is used for storing raw Conn information
	Agent interface {
		GetSession() session.Session
		Push(route string, v interface{}) error
		ResponseMID(ctx context.Context, mid uint, v interface{}, isError ...bool) error
		Close() error
		RemoteAddr() net.Addr
		String() string
		GetStatus() int32
		Kick(ctx context.Context) error
		SetLastAt()
		SetStatus(state int32)
		Handle()
		IPVersion() string
		SendHandshakeResponse() error
		SendHandshakeErrorResponse() error
		SendRequest(ctx context.Context, serverID, route string, v interface{}) (*protos.Response, error)
		AnswerWithError(ctx context.Context, mid uint, err error)
	}

	// AgentFactory factory for creating Agent instances
	AgentFactory interface {
		CreateAgent(conn net.Conn) Agent
	}

	agentFactoryImpl struct {
		sessionPool        session.SessionPool
		appDieChan         chan bool           // app die channel
		decoder            codec.PacketDecoder // binary decoder
		encoder            codec.PacketEncoder // binary encoder
		heartbeatTimeout   time.Duration
		writeTimeout       time.Duration
		messageEncoder     message.Encoder
		messagesBufferSize int // size of the pending messages buffer
		metricsReporters   []metrics.Reporter
		serializer         serialize.Serializer // message serializer
	}
)

// NewAgentFactory ctor
func NewAgentFactory(
	appDieChan chan bool,
	decoder codec.PacketDecoder,
	encoder codec.PacketEncoder,
	serializer serialize.Serializer,
	heartbeatTimeout time.Duration,
	writeTimeout time.Duration,
	messageEncoder message.Encoder,
	messagesBufferSize int,
	sessionPool session.SessionPool,
	metricsReporters []metrics.Reporter,
) AgentFactory {
	_ = "STUB: not implemented"
	return *new(AgentFactory)
}

// CreateAgent returns a new agent
func (f *agentFactoryImpl) CreateAgent(conn net.Conn) Agent {
	_ = "STUB: not implemented"
	return *new(Agent)
}

// NewAgent create new agent instance
func newAgent(
	conn net.Conn,
	packetDecoder codec.PacketDecoder,
	packetEncoder codec.PacketEncoder,
	serializer serialize.Serializer,
	heartbeatTime time.Duration,
	writeTimeout time.Duration,
	messagesBufferSize int,
	dieChan chan bool,
	messageEncoder message.Encoder,
	metricsReporters []metrics.Reporter,
	sessionPool session.SessionPool,
) Agent {
	_ = "STUB: not implemented"
	// initialize heartbeat and handshake data on first user connection
	return *new(Agent)
}

// binding session

func (a *agentImpl) getMessageFromPendingMessage(pm pendingMessage) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// construct message and encode

func (a *agentImpl) packetEncodeMessage(m *message.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// packet encode

func (a *agentImpl) send(pendingMsg pendingMessage) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Capture the error code from the payload BEFORE packetEncodeMessage
// runs, because MessagesEncoder.Encode mutates m.Data in place when
// compression produces smaller output. If we read m.Data after encoding,
// the payload may be deflated bytes and GetErrorFromPayload's silent
// Unmarshal failure causes the metric code tag to fall back to the
// default ErrUnknownCode (PIT-000), masking the real status code for
// any error whose JSON payload is large enough to compress.

// packet encode

// chSend is never closed so we need this to don't block if agent is already closed

// GetSession returns the agent session
func (a *agentImpl) GetSession() session.Session {
	_ = "STUB: not implemented"

	// Push implementation for NetworkEntity interface
	return *new(session.Session)
}

func (a *agentImpl) Push(route string, v interface{}) error { _ = "STUB: not implemented"; return nil }

// ResponseMID implementation for NetworkEntity interface
// Respond message to session
func (a *agentImpl) ResponseMID(ctx context.Context, mid uint, v interface{}, isError ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the agent, cleans inner state and closes low-level connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (a *agentImpl) Close() error { _ = "STUB: not implemented"; return nil }

// prevent closing closed channel

// expect

// RemoteAddr implementation for NetworkEntity interface
// returns the remote network address.
func (a *agentImpl) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// String, implementation for Stringer interface
func (a *agentImpl) String() string { _ = "STUB: not implemented"; return "" }

// GetStatus gets the status
func (a *agentImpl) GetStatus() int32 { _ = "STUB: not implemented"; return 0 }

// Kick sends a kick packet to a client
func (a *agentImpl) Kick(ctx context.Context) error {
	_ = "STUB: not implemented"
	// packet encode
	return nil
}

// SetLastAt sets the last at to now
func (a *agentImpl) SetLastAt() { _ = "STUB: not implemented"; return }

// SetStatus sets the agent status
func (a *agentImpl) SetStatus(state int32) { _ = "STUB: not implemented"; return }

// Handle handles the messages from and to a client
func (a *agentImpl) Handle() { _ = "STUB: not implemented"; return }

// agent closed signal

// IPVersion returns the remote address ip version.
// net.TCPAddr and net.UDPAddr implementations of String()
// always construct result as <ip>:<port> on both
// ipv4 and ipv6. Also, to see if the ip is ipv6 they both
// check if there is a colon on the string.
// So checking if there are more than one colon here is safe.
func (a *agentImpl) IPVersion() string { _ = "STUB: not implemented"; return "" }

func (a *agentImpl) heartbeat() { _ = "STUB: not implemented"; return }

// chSend is never closed so we need this to don't block if agent is already closed

func (a *agentImpl) onSessionClosed(s session.Session) { _ = "STUB: not implemented"; return }

// SendHandshakeResponse sends a handshake response
func (a *agentImpl) SendHandshakeResponse() error { _ = "STUB: not implemented"; return nil }

func (a *agentImpl) SendHandshakeErrorResponse() error { _ = "STUB: not implemented"; return nil }

func (a *agentImpl) write() {
	_ = "STUB: not implemented"
	// clean func
	return
}

// Log the timeout error but continue processing

// close agent if low-level conn broke during write

func (a *agentImpl) writeToConnection(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func createConnectionSpan(ctx context.Context, conn net.Conn, op string) trace.Span {
	_ = "STUB: not implemented"
	return *new(trace.Span)
}

// SendRequest sends a request to a server
func (a *agentImpl) SendRequest(ctx context.Context, serverID, route string, v interface{}) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AnswerWithError answers with an error
func (a *agentImpl) AnswerWithError(ctx context.Context, mid uint, err error) {
	_ = "STUB: not implemented"
	return
}

func hbdEncode(heartbeatTimeout time.Duration, packetEncoder codec.PacketEncoder, dataCompression bool, serializerName string) {
	_ = "STUB: not implemented"
	return
}

func herdEncode(heartbeatTimeout time.Duration, packetEncoder codec.PacketEncoder, dataCompression bool, serializerName string) {
	_ = "STUB: not implemented"
	return
}

func encodeAndCompress(data interface{}, dataCompression bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *agentImpl) reportChannelSize() { _ = "STUB: not implemented"; return }
