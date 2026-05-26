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

package pitaya

import (
	"github.com/topfreegames/pitaya/v3/pkg/component"
)

type regComp struct {
	comp component.Component
	opts []component.Option
}

// Register register a component with options
func (app *App) Register(c component.Component, options ...component.Option) {
	_ = "STUB: not implemented"
	return
}

// RegisterRemote register a remote component with options
func (app *App) RegisterRemote(c component.Component, options ...component.Option) {
	_ = "STUB: not implemented"
	return
}

func (app *App) startupComponents() {
	_ = "STUB: not implemented"
	// handler component initialize hooks
	return
}

// handler component after initialize hooks

// remote component initialize hooks

// remote component after initialize hooks

// register all components

// register all remote components

func (app *App) shutdownComponents() {
	_ = "STUB: not implemented"
	// reverse call `BeforeShutdown` hooks
	return
}

// reverse call `Shutdown` hooks

// reverse call `Shutdown` hooks
