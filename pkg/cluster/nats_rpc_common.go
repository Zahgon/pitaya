// Copyright (c) TFG Co. All Rights Reserved.
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

package cluster

import (
	nats "github.com/nats-io/nats.go"
)

func getChannel(serverType, serverID string) string { _ = "STUB: not implemented"; return "" }

func drainAndClose(nc *nats.Conn) error { _ = "STUB: not implemented"; return nil }

// If connection is already closed, just return

// Drain connection (this will flush any pending messages and prevent new ones)

// Even if drain fails, try to close (but only if not already closed)

// Wait for drain to complete with timeout

// Close will happen automatically after drain completes

// replaceNatsConnection handles the common logic for replacing NATS connections
// It stores old connection/subscription references, calls initFunc to set up the new connection,
// and then drains the old resources after the new connection is ready.
func replaceNatsConnection(
	oldConn *nats.Conn,
	oldSub *nats.Subscription,
	initFunc func() error,
	componentName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Re-initialize connection (pass true to indicate this is a replacement)

// Drain and close old connection and subscription after new one is set up

func setupNatsConn(connectString string, appDieChan chan bool, lameDuckReplacement func() error, options ...nats.Option) (*nats.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Recover from potential panics when sending on channels that may be closed
// This can happen if the callback is invoked after setupNatsConn has returned

// Channel might be closed or in invalid state, just log and continue
// This is safe because setupNatsConn has already handled the error

// If connection was never successfully established, prioritize initialConnectErrorCh
// to allow setupNatsConn to return quickly with an error

// During initial connection, send error to initialConnectErrorCh first

// If channel is not ready, fall through to appDieChan handling

// On Windows, Signal() with Interrupt works
// On Unix-like systems, this is equivalent to SIGINT

// If no appDieChan and connection was never established, try initialConnectErrorCh again

// Channel not ready, but we've already logged the error

// The old connection will eventually close (it's in lame duck mode),
// which will trigger ClosedHandler and appDieChan

// This is non-deterministic becase jitter TLS is different and we need to simplify
// the calculations. What we want to do is simply not block forever the call while
// we don't set a timeout so low that hinders our own reconnect config:
// 		maxReconnectTimeout = reconnectWait + reconnectJitter + reconnectTimeout
// 		connectionTimeout + (maxReconnectionAttemps * maxReconnectTimeout)
// Thus, the time.After considers 2 times this value
