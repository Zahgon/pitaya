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

package errors

// ErrUnknownCode is a string code representing an unknown error
// This will be used when no error code is sent by the handler
const ErrUnknownCode = "PIT-000"

// ErrInternalCode is a string code representing an internal Pitaya error
const ErrInternalCode = "PIT-500"

// ErrNotFoundCode is a string code representing a not found related error
const ErrNotFoundCode = "PIT-404"

// ErrBadRequestCode is a string code representing a bad request related error
const ErrBadRequestCode = "PIT-400"

// ErrClientClosedRequest is a string code representing the client closed request error
const ErrClientClosedRequest = "PIT-499"

// ErrClosedRequest is a string code representing the closed request error
const ErrClosedRequest = "PIT-498"

// Error is an error with a code, message and metadata
type Error struct {
	Code     string
	Message  string
	Metadata map[string]string
}

// NewError ctor
func NewError(err error, code string, metadata ...map[string]string) *Error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func mergeMetadatas(pitayaErr *Error, metadata map[string]string) {
	_ = "STUB: not implemented"
	return
}

// CodeFromError returns the code of error.
// If error is nil, return empty string.
// If error is not a pitaya error, returns unkown code
func CodeFromError(err error) string { _ = "STUB: not implemented"; return "" }
