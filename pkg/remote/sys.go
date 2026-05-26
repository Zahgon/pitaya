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

package remote

import (
	"context"

	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/session"
)

// Sys contains logic for handling sys remotes
type Sys struct {
	component.Base
	sessionPool session.SessionPool
}

// NewSys returns a new Sys instance
func NewSys(sessionPool session.SessionPool) *Sys { _ = "STUB: not implemented"; return nil }

// BindSession binds the local session
func (s *Sys) BindSession(ctx context.Context, sessionData *protos.Session) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushSession updates the local session
func (s *Sys) PushSession(ctx context.Context, sessionData *protos.Session) (*protos.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Kick kicks a local user
func (s *Sys) Kick(ctx context.Context, msg *protos.KickMsg) (*protos.KickAnswer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
