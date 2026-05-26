package pitaya

import (
	"sync"
	"time"

	"github.com/grafana/sobek"
	pitayaclient "github.com/topfreegames/pitaya/v3/pkg/client"
	"github.com/topfreegames/pitaya/v3/pkg/session"
	"go.k6.io/k6/js/modules"
)

// Response is the type of the response returned by the server
type Response interface{}

// Client is the pitaya client
// It is used to connect to a pitaya server and send requests and notifies
// It is also used to consume pushes
type Client struct {
	vu             modules.VU
	client         pitayaclient.PitayaClient
	handshake      *session.HandshakeData
	responsesMutex sync.Mutex
	responses      map[uint]chan []byte
	pushesMutex    sync.Mutex
	pushes         map[string]chan []byte
	timeout        time.Duration
	metrics        *pitayaMetrics
	useTLS         bool
}

// Connect connects to the server
// addr is the address of the server to connect to
func (c *Client) Connect(addr string) error { _ = "STUB: not implemented"; return nil }

// IsConnected returns true if the client is connected to the server
func (c *Client) IsConnected() bool { _ = "STUB: not implemented"; return false }

// ConsumePush will return a promise that will be resolved when a push is received on the given route.
// The promise will be rejected if the timeout is reached before a push is received.
// The promise will be resolved with the push data.
func (c *Client) ConsumePush(route string, timeoutMs int) *sobek.Promise {
	_ = "STUB: not implemented"
	return nil
}

// Notify sends a notify to the server
// route is the route to send the notify to
// msg is the message to send
// returns an error if the notify could not be sent
func (c *Client) Notify(route string, msg interface{}) error { _ = "STUB: not implemented"; return nil }

// RequestB64 sends a request to the server using a base64 string
// route is the route to send the request to
// str is the string passed in request
// returns a promise that will be resolved when the response is received
// the promise will be rejected if the timeout is reached before a response is received
func (c *Client) RequestB64(route string, b64msg string) *sobek.Promise {
	_ = "STUB: not implemented" // TODO: add custom timeout
	return nil
}

// Request sends a request to the server
// route is the route to send the request to
// msg is the message to send
// returns a promise that will be resolved when the response is received
// the promise will be rejected if the timeout is reached before a response is received
func (c *Client) Request(route string, msg interface{}) *sobek.Promise {
	_ = "STUB: not implemented" // TODO: add custom timeout
	return nil
}

func (c *Client) pushRequestMetrics(route string, responseTime time.Duration, success bool, timeout bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Disconnect disconnects from the server
func (c *Client) Disconnect() { _ = "STUB: not implemented"; return }

func (c *Client) listen() { _ = "STUB: not implemented"; return }

// only keep one message in the channel, discard the rest

func (c *Client) getResponseChannelForID(id uint) chan []byte {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) removeResponseChannelForID(id uint) { _ = "STUB: not implemented"; return }

func (c *Client) getPushChannelForRoute(route string) chan []byte {
	_ = "STUB: not implemented"
	return nil
}

// makeHandledPromise will create a promise and return its resolve and reject methods,
// wrapped in such a way that it will block the eventloop from exiting before they are
// called even if the promise isn't resolved by the time the current script ends executing.
func (c *Client) makeHandledPromise() (*sobek.Promise, func(interface{}), func(interface{})) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// more stuff

// more stuff
