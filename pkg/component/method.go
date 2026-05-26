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

package component

import (
	"context"
	"reflect"

	"github.com/golang/protobuf/proto"
)

var (
	typeOfError    = reflect.TypeOf((*error)(nil)).Elem()
	typeOfBytes    = reflect.TypeOf(([]byte)(nil))
	typeOfContext  = reflect.TypeOf(new(context.Context)).Elem()
	typeOfProtoMsg = reflect.TypeOf(new(proto.Message)).Elem()
)

func isExported(name string) bool { _ = "STUB: not implemented"; return false }

// isRemoteMethod decide a method is suitable remote method
func isRemoteMethod(method reflect.Method) bool {
	_ = "STUB: not implemented"

	// Method must be exported.
	return false
}

// Method needs at least two ins: receiver and context.Context

// Method needs two outs: interface{}(that implements proto.Message), error

// isHandlerMethod decide a method is suitable handler method
func isHandlerMethod(method reflect.Method) bool {
	_ = "STUB: not implemented"

	// Method must be exported.
	return false
}

// Method needs two or three ins: receiver, context.Context and optional []byte or pointer.

// Method needs either no out or two outs: interface{}(or []byte), error

func suitableRemoteMethods(typ reflect.Type, nameFunc func(string) string) map[string]*Remote {
	_ = "STUB: not implemented"
	return nil
}

// rewrite remote name

func suitableHandlerMethods(typ reflect.Type, nameFunc func(string) string) map[string]*Handler {
	_ = "STUB: not implemented"
	return nil
}

// rewrite handler name
