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

package pitaya

import (
	"github.com/topfreegames/pitaya/v3/pkg/interfaces"
)

type moduleWrapper struct {
	module interfaces.Module
	name   string
}

type sessionModuleWrapper struct {
	module interfaces.SessionModule
	name   string
}

// RegisterModule registers a module, by default it register after registered modules
func (app *App) RegisterModule(module interfaces.Module, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterModuleAfter registers a module after all registered modules
func (app *App) RegisterModuleAfter(module interfaces.Module, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterModuleBefore registers a module before all registered modules
func (app *App) RegisterModuleBefore(module interfaces.Module, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetModule gets a module with a name
func (app *App) GetModule(name string) (interfaces.Module, error) {
	_ = "STUB: not implemented"
	return *new(interfaces.Module), nil
}

func (app *App) alreadyRegistered(name string) error { _ = "STUB: not implemented"; return nil }

// startModules starts all modules in order
func (app *App) startModules() { _ = "STUB: not implemented"; return }

func (app *App) startModuleSessionDraining() { _ = "STUB: not implemented"; return }

func (app *App) maxModuleSessionCount() int64 { _ = "STUB: not implemented"; return 0 }

// shutdownModules starts all modules in reverse order
func (app *App) shutdownModules() { _ = "STUB: not implemented"; return }
