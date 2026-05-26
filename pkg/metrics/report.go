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

package metrics

import (
	"context"
	"time"
)

// ReportTimingFromCtx reports the latency from the context
func ReportTimingFromCtx(ctx context.Context, reporters []Reporter, typ string, err error) {
	_ = "STUB: not implemented"
	return
}

// ReportMessageProcessDelayFromCtx reports the delay to process the messages
func ReportMessageProcessDelayFromCtx(ctx context.Context, reporters []Reporter, typ string) {
	_ = "STUB: not implemented"
	return
}

// ReportNumberOfConnectedClients reports the number of connected clients
func ReportNumberOfConnectedClients(reporters []Reporter, number int64) {
	_ = "STUB: not implemented"
	return
}

// ReportSysMetrics reports sys metrics
func ReportSysMetrics(reporters []Reporter, period time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ReportExceededRateLimiting reports the number of requests made
// after exceeded rate limiting in a connection
func ReportExceededRateLimiting(reporters []Reporter) { _ = "STUB: not implemented"; return }

func tagsFromContext(ctx context.Context) map[string]string { _ = "STUB: not implemented"; return nil }

func getTags(ctx context.Context, tags map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
