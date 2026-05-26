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

package service

import (
	"context"
	"errors"

	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/pipeline"
	"github.com/topfreegames/pitaya/v3/pkg/route"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
	"github.com/topfreegames/pitaya/v3/pkg/session"
)

var errInvalidMsg = errors.New("invalid message type provided")

func unmarshalHandlerArg(handler *component.Handler, serializer serialize.Serializer, payload []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalRemoteArg(remote *component.Remote, payload []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMsgType(msgTypeIface interface{}) (message.Type, error) {
	_ = "STUB: not implemented"
	return *new(message.Type), nil
}

func serializeReturn(ser serialize.Serializer, ret interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processHandlerMessage(
	ctx context.Context,
	rt *route.Route,
	handler *component.Handler,
	serializer serialize.Serializer,
	handlerHooks *pipeline.HandlerHooks,
	session session.Session,
	data []byte,
	msgTypeIface interface{},
	remote bool,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First unmarshal the handler arg that will be passed to
// both handler and pipeline functions

// This is a special case and should only happen with nats rpc client
// because we used nats request we have to answer to it or else a timeout
// will happen in the caller server and will be returned to the client
// the reason why we don't just Publish is to keep track of failed rpc requests
// with timeouts, maybe we can improve this flow
