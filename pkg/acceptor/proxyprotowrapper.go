package acceptor

import (
	"net"
	"sync"
)

// Listener is used to wrap an underlying listener,
// whose connections may be using the HAProxy Proxy Protocol.
// If the connection is using the protocol, the RemoteAddr() will return
// the correct client address.
type ProxyProtocolListener struct {
	net.Listener
	proxyProtocolEnabled *bool
}

// Accept waits for and returns the next connection to the listener.
func (p *ProxyProtocolListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	// Get the underlying connection
	return *new(net.Conn), nil
}

// Conn is used to wrap and underlying connection which
// may be speaking the Proxy Protocol. If it is, the RemoteAddr() will
// return the address of the client instead of the proxy address.
type Conn struct {
	net.Conn
	dstAddr              *net.Addr
	srcAddr              *net.Addr
	once                 sync.Once
	proxyProtocolEnabled *bool
}

func (p *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// RemoteAddr returns the address of the client if the proxy
// protocol is being used, otherwise just returns the address of
// the socket peer. If there is an error parsing the header, the
// address of the client is not returned, and the socket is closed.
// Once implication of this is that the call could block if the
// client is slow. Using a Deadline is recommended if this is called
// before Read()
func (p *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (p *Conn) checkPrefix() error { _ = "STUB: not implemented"; return nil }
