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

package util

import (
	"context"
	"reflect"

	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/logger/interfaces"
	"github.com/topfreegames/pitaya/v3/pkg/protos"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
)

func getLoggerFromArgs(args []reflect.Value) interfaces.Logger {
	_ = "STUB: not implemented"
	return *new(interfaces.Logger)
}

// Pcall calls a method that returns an interface and an error and recovers in case of panic
func Pcall(method reflect.Method, args []reflect.Value) (rets interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to use logger from context here to help trace error cause

// r can have 0 length in case of notify handlers
// otherwise it will have 2 outputs: an interface and an error

// SliceContainsString returns true if a slice contains the string
func SliceContainsString(slice []string, str string) bool { _ = "STUB: not implemented"; return false }

// SerializeOrRaw serializes the interface if its not an array of bytes already
func SerializeOrRaw(serializer serialize.Serializer, v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FileExists tells if a file exists
func FileExists(filename string) bool { _ = "STUB: not implemented"; return false }

// GetErrorFromPayload gets the error from payload
func GetErrorFromPayload(serializer serialize.Serializer, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// GetErrorPayload creates and serializes an error payload
func GetErrorPayload(serializer serialize.Serializer, err error) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertProtoToMessageType converts a protos.MsgType to a message.Type
func ConvertProtoToMessageType(protoMsgType protos.MsgType) message.Type {
	_ = "STUB: not implemented"
	return *new(message.Type)
}

// CtxWithDefaultLogger inserts a default logger on ctx to be used on handlers and remotes.
// If using logrus, userId, route and requestId will be added as fields.
// Otherwise the pitaya logger will be used as it is.
func CtxWithDefaultLogger(ctx context.Context, route, userID string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// StartSpanFromRequest starts a tracing span from the request
func StartSpanFromRequest(
	ctx context.Context,
	serverID, route string,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Create a new context and span

// GetContextFromRequest gets the context from a request
func GetContextFromRequest(req *protos.Request, serverID string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
