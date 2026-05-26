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

package config

import (
	"github.com/mitchellh/mapstructure"

	"github.com/spf13/viper"
)

// Config is a wrapper around a viper config
type Config struct {
	viper.Viper
}

// NewConfig creates a new config with a given viper config if given
func NewConfig(cfgs ...*viper.Viper) *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) fillDefaultValues() { _ = "STUB: not implemented"; return }

// the max buffer size that nats will accept, if this buffer overflows, messages will begin to be dropped

// the sum of this config among all the frontend servers should always be less than
// the sum of pitaya.buffer.cluster.rpc.server.nats.messages, for covering the worst case scenario
// a single backend server should have the config pitaya.buffer.cluster.rpc.server.nats.messages bigger
// than the sum of the config pitaya.concurrency.handler.dispatch among all frontend servers

// UnmarshalKey unmarshals key into v
func (c *Config) UnmarshalKey(key string, rawVal interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func isStringMapInterface(val interface{}) bool { _ = "STUB: not implemented"; return false }

// A wrapper around mapstructure.Decode that mimics the WeakDecode functionality
func decode(input interface{}, config *mapstructure.DecoderConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// defaultDecoderConfig returns default mapstructure.DecoderConfig with support
// of time.Duration values & string slices
func defaultDecoderConfig(output interface{}, opts ...viper.DecoderConfigOption) *mapstructure.DecoderConfig {
	_ = "STUB: not implemented"
	return nil
}
