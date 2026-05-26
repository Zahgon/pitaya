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

package tracing

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	otelTrace "go.opentelemetry.io/otel/trace"
)

func castValueToCarrier(val interface{}) (propagation.MapCarrier, error) {
	_ = "STUB: not implemented"
	return *new(propagation.MapCarrier), nil
}

// ExtractSpan retrieves an OpenTelemetry span context from the given context.Context
// The span context can be received directly (inside the context) or via an RPC call
// (encoded in a carrier)
func ExtractSpan(ctx context.Context) (otelTrace.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(otelTrace.SpanContext), nil
}

// InjectSpan retrieves an OpenTelemetry span from the current context and creates a new context
// with it encoded in text map format inside the propagatable context content
func InjectSpan(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// StartSpan starts a new span with a given parent context, operation name, and attributes.
// It returns a context with the created span.
func StartSpan(
	parentCtx context.Context,
	opName string,
	attributes ...attribute.KeyValue,
) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

// FinishSpan finishes a span retrieved from the given context and logs the error if it exists
func FinishSpan(ctx context.Context, err error) { _ = "STUB: not implemented"; return }
