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

package pipeline

import (
	"context"
)

type (
	// HandlerTempl is a function that has the same signature as a handler and will
	// be called before or after handler methods
	HandlerTempl func(ctx context.Context, in interface{}) (c context.Context, out interface{}, err error)

	// AfterHandlerTempl is a function for the after handler, receives both the handler response
	// and the error returned
	AfterHandlerTempl func(ctx context.Context, out interface{}, err error) (interface{}, error)

	// Channel contains the functions to be called before the handler method is executed
	Channel struct {
		Handlers []HandlerTempl
	}

	// AfterChannel contains the functions to be called after the handler method is executed
	AfterChannel struct {
		Handlers []AfterHandlerTempl
	}

	// Hooks contains before and after channels
	Hooks struct {
		BeforeHandler *Channel
		AfterHandler  *AfterChannel
	}

	HandlerHooks struct {
		Hooks
	}

	RemoteHooks struct {
		Hooks
	}
)

// NewHandlerHooks ctor
func NewHandlerHooks() *HandlerHooks { _ = "STUB: not implemented"; return nil }

// NewRemoteHooks ctor
func NewRemoteHooks() *RemoteHooks { _ = "STUB: not implemented"; return nil }

// NewChannel ctor
func NewChannel() *Channel { _ = "STUB: not implemented"; return nil }

// NewAfterChannel ctor
func NewAfterChannel() *AfterChannel { _ = "STUB: not implemented"; return nil }

// ExecuteBeforePipeline calls registered handlers
func (p *Channel) ExecuteBeforePipeline(ctx context.Context, data interface{}) (context.Context, interface{}, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// ExecuteAfterPipeline calls registered handlers
func (p *AfterChannel) ExecuteAfterPipeline(ctx context.Context, res interface{}, err error) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushFront should not be used after pitaya is running
func (p *Channel) PushFront(h HandlerTempl) { _ = "STUB: not implemented"; return }

// PushBack should not be used after pitaya is running
func (p *Channel) PushBack(h HandlerTempl) { _ = "STUB: not implemented"; return }

// Clear should not be used after pitaya is running
func (p *Channel) Clear() { _ = "STUB: not implemented"; return }

// PushFront should not be used after pitaya is running
func (p *AfterChannel) PushFront(h AfterHandlerTempl) { _ = "STUB: not implemented"; return }

// PushBack should not be used after pitaya is running
func (p *AfterChannel) PushBack(h AfterHandlerTempl) { _ = "STUB: not implemented"; return }

// Clear should not be used after pitaya is running
func (p *AfterChannel) Clear() { _ = "STUB: not implemented"; return }
