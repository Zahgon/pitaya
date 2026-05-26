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

package acceptor

import (
	"crypto/tls"
	"net"
)

// TCPAcceptor struct
type TCPAcceptor struct {
	addr          string
	connChan      chan PlayerConn
	listener      net.Listener
	running       bool
	certs         []tls.Certificate
	proxyProtocol bool
}

type tcpPlayerConn struct {
	net.Conn
	remoteAddr net.Addr
}

func (t *tcpPlayerConn) RemoteAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// GetNextMessage reads the next message available in the stream
	new(net.Addr)
}

func (t *tcpPlayerConn) GetNextMessage() (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the header has no data, we can consider it as a closed connection

// NewTCPAcceptor creates a new instance of tcp acceptor
func NewTCPAcceptor(addr string, certs ...string) *TCPAcceptor {
	_ = "STUB: not implemented"
	return nil
}

func NewTLSAcceptor(addr string, certs ...tls.Certificate) *TCPAcceptor {
	_ = "STUB: not implemented"
	return nil
}

// GetAddr returns the addr the acceptor will listen on
func (a *TCPAcceptor) GetAddr() string { _ = "STUB: not implemented"; return "" }

// GetConnChan gets a connection channel
func (a *TCPAcceptor) GetConnChan() chan PlayerConn {
	_ = "STUB: not implemented"

	// Stop stops the acceptor
	return nil
}

func (a *TCPAcceptor) Stop() { _ = "STUB: not implemented"; return }

func (a *TCPAcceptor) hasTLSCertificates() bool { _ = "STUB: not implemented"; return false }

// ListenAndServe using tcp acceptor
func (a *TCPAcceptor) ListenAndServe() { _ = "STUB: not implemented"; return }

// ListenAndServeTLS listens using tls
func (a *TCPAcceptor) ListenAndServeTLS(cert, key string) { _ = "STUB: not implemented"; return }

// Create base listener
func (a *TCPAcceptor) createBaseListener() net.Listener {
	_ = "STUB: not implemented"
	// Create raw listener
	return *new(net.Listener)
}

// Wrap listener in ProxyProto

// ListenAndServeTLS listens using tls
func (a *TCPAcceptor) listenAndServeTLS(listener net.Listener) net.Listener {
	_ = "STUB: not implemented"
	return *new(net.Listener)
}

func (a *TCPAcceptor) EnableProxyProtocol() { _ = "STUB: not implemented"; return }

func (a *TCPAcceptor) serve() { _ = "STUB: not implemented"; return }

func (a *TCPAcceptor) IsRunning() bool { _ = "STUB: not implemented"; return false }

func (a *TCPAcceptor) GetConfiguredAddress() string { _ = "STUB: not implemented"; return "" }
