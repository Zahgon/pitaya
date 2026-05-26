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

package client

import (
	"crypto/tls"
	"time"

	protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
)

// Command struct. Save the input and output type and proto descriptor for each
// one.
type Command struct {
	input               string // input command name
	output              string // output command name
	inputMsgDescriptor  *desc.MessageDescriptor
	outputMsgDescriptor *desc.MessageDescriptor
}

// ProtoBufferInfo save all commands from a server.
type ProtoBufferInfo struct {
	Commands map[string]*Command
}

// ProtoClient struct
type ProtoClient struct {
	Client
	descriptorsNames        map[string]bool
	info                    ProtoBufferInfo
	docsRoute               string
	descriptorsRoute        string
	IncomingMsgChan         chan *message.Message
	expectedInputDescriptor *desc.MessageDescriptor
	ready                   bool
	closeChan               chan bool
}

// MsgChannel return the incoming message channel
func (pc *ProtoClient) MsgChannel() chan *message.Message { _ = "STUB: not implemented"; return nil }

// Receive a compressed byte slice and unpack it to a FileDescriptorProto
func unpackDescriptor(compressedDescriptor []byte) (*protobuf.FileDescriptorProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Receive an array of descriptors in binary format. The function creates the
// protobuffer from this data and associates it to the message.
func (pc *ProtoClient) buildProtosFromDescriptor(descriptorArray []*protobuf.FileDescriptorProto) error {
	_ = "STUB: not implemented"
	return nil
}

// Receives each entry from the Unmarshal json from the Docs and read the inputs and
// outputs associated with it. Return the output type, the input and the error.
func getOutputInputNames(command map[string]interface{}) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// we can have handlers that have no return specified.

// Get recursively all protos needed in a Unmarshal json.
func getKeys(info map[string]interface{}, keysSet map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// Receives one json string from the auto documentation, decode it and request
// to the server the protobuf descriptors. If the the  descriptors route are
// not set, this function identify the route responsible for providing the
// protobuf descriptors.
func (pc *ProtoClient) getDescriptors(data string) error { _ = "STUB: not implemented"; return nil }

// load predefined protos

// build commands reference

// get all proto types

// Return the basic structure for the ProtoClient struct.
func newProto(docslogLevel logrus.Level, requestTimeout ...time.Duration) *ProtoClient {
	_ = "STUB: not implemented"
	return nil
}

// NewProto returns a new protoclient with the auto documentation route.
func NewProto(docsRoute string, docslogLevel logrus.Level, requestTimeout ...time.Duration) *ProtoClient {
	_ = "STUB: not implemented"
	return nil
}

// NewWithDescriptor returns a new protoclient with the descriptors route and
// auto documentation route.
func NewWithDescriptor(descriptorsRoute string, docsRoute string, docslogLevel logrus.Level, requestTimeout ...time.Duration) *ProtoClient {
	_ = "STUB: not implemented"
	return nil
}

// LoadServerInfo load commands information from the server. Addr is the
// server address.
func (pc *ProtoClient) LoadServerInfo(addr string) error { _ = "STUB: not implemented"; return nil }

// request doc info

// Disconnect the client
func (pc *ProtoClient) Disconnect() { _ = "STUB: not implemented"; return }

// Wait for new messages from the server or the connection end. If the menssage
// has a response.Route, it decodes based on it. If not, it will try to decode
// the menssage using the last expected response.
func (pc *ProtoClient) waitForData() { _ = "STUB: not implemented"; return }

// ConnectTo connects to the server at addr, for now the only supported protocol is tcp
// this methods blocks as it also handles the messages from the server
func (pc *ProtoClient) ConnectTo(addr string, tlsConfig ...*tls.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportInformation export supported server commands information
func (pc *ProtoClient) ExportInformation() *ProtoBufferInfo { _ = "STUB: not implemented"; return nil }

// LoadInfo load commands information form ProtoBufferInfo
func (pc *ProtoClient) LoadInfo(info *ProtoBufferInfo) error { _ = "STUB: not implemented"; return nil }

// AddPushResponse add a push response. Must be ladded before LoadInfo.
func (pc *ProtoClient) AddPushResponse(route string, protoName string) {
	_ = "STUB: not implemented"
	return
}

// SendRequest sends a request to the server
func (pc *ProtoClient) SendRequest(route string, data []byte) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SendNotify sends a notify to the server
func (pc *ProtoClient) SendNotify(route string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
