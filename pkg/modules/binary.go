//go:build linux || darwin

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

package modules

import (
	"os/exec"
	"time"
)

// Binary is a pitaya module that starts a binary as a child process and
// pipes its stdout
type Binary struct {
	Base
	binPath                  string
	args                     []string
	gracefulShutdownInterval time.Duration
	cmd                      *exec.Cmd
	exitCh                   chan struct{}
}

// NewBinary creates a new binary module with the given path
func NewBinary(binPath string, args []string, gracefulShutdownInterval ...time.Duration) *Binary {
	_ = "STUB: not implemented"
	return nil
}

// GetExitChannel gets a channel that is closed when the binary dies
func (b *Binary) GetExitChannel() chan struct{} {
	_ = "STUB: not implemented"

	// Init initializes the binary
	return nil
}

func (b *Binary) Init() error { _ = "STUB: not implemented"; return nil }

// Shutdown shutdowns the binary module
func (b *Binary) Shutdown() error { _ = "STUB: not implemented"; return nil }
