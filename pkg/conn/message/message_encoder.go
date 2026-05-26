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

package message

// Encoder interface
type Encoder interface {
	IsCompressionEnabled() bool
	Encode(message *Message) ([]byte, error)
}

// MessagesEncoder implements MessageEncoder interface
type MessagesEncoder struct {
	DataCompression bool
}

// NewMessagesEncoder returns a new message encoder
func NewMessagesEncoder(dataCompression bool) *MessagesEncoder {
	_ = "STUB: not implemented"
	return nil
}

// IsCompressionEnabled returns wether the compression is enabled or not
func (me *MessagesEncoder) IsCompressionEnabled() bool { _ = "STUB: not implemented"; return false }

// Encode marshals message to binary format. Different message types is corresponding to
// different message header, message types is identified by 2-4 bit of flag field. The
// relationship between message types and message header is presented as follows:
// ------------------------------------------
// |   type   |  flag  |       other        |
// |----------|--------|--------------------|
// | request  |----000-|<message id>|<route>|
// | notify   |----001-|<route>             |
// | response |----010-|<message id>        |
// | push     |----011-|<route>             |
// ------------------------------------------
// The figure above indicates that the bit does not affect the type of message.
// See ref: https://github.com/topfreegames/pitaya/v3/blob/master/docs/communication_protocol.md
func (me *MessagesEncoder) Encode(message *Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// variant length encode

// Decode decodes the message
func (me *MessagesEncoder) Decode(data []byte) (*Message, error) {
	_ = "STUB: not implemented"
	return nil,

		// Decode unmarshal the bytes slice to a message
		// See ref: https://github.com/topfreegames/pitaya/v3/blob/master/docs/communication_protocol.md
		nil
}

func Decode(data []byte) (*Message, error) { _ = "STUB: not implemented"; return nil, nil }

// little end byte order
// WARNING: must can be stored in 64 bits integer
// variant length encode
