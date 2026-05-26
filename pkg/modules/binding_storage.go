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
	"time"

	"github.com/topfreegames/pitaya/v3/pkg/cluster"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/session"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ETCDBindingStorage module that uses etcd to keep in which frontend server each user is bound
type ETCDBindingStorage struct {
	Base
	cli             *clientv3.Client
	etcdEndpoints   []string
	etcdPrefix      string
	etcdDialTimeout time.Duration
	leaseTTL        time.Duration
	leaseID         clientv3.LeaseID
	thisServer      *cluster.Server
	sessionPool     session.SessionPool
	stopChan        chan struct{}
}

// NewETCDBindingStorage returns a new instance of BindingStorage
func NewETCDBindingStorage(server *cluster.Server, sessionPool session.SessionPool, conf config.ETCDBindingConfig) *ETCDBindingStorage {
	_ = "STUB: not implemented"
	return nil
}

func getUserBindingKey(uid, frontendType string) string { _ = "STUB: not implemented"; return "" }

// PutBinding puts the binding info into etcd
func (b *ETCDBindingStorage) PutBinding(uid string) error { _ = "STUB: not implemented"; return nil }

func (b *ETCDBindingStorage) removeBinding(uid string) error { _ = "STUB: not implemented"; return nil }

// GetUserFrontendID gets the id of the frontend server a user is connected to
// TODO: should we set context here?
// TODO: this could be way more optimized, using watcher and local caching
func (b *ETCDBindingStorage) GetUserFrontendID(uid, frontendType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *ETCDBindingStorage) setupOnSessionCloseCB() { _ = "STUB: not implemented"; return }

func (b *ETCDBindingStorage) setupOnAfterSessionBindCB() { _ = "STUB: not implemented"; return }

func (b *ETCDBindingStorage) watchLeaseChan(c <-chan *clientv3.LeaseKeepAliveResponse) {
	_ = "STUB: not implemented"
	return
}

func (b *ETCDBindingStorage) bootstrapLease() error {
	_ = "STUB: not implemented"
	// grab lease
	return nil
}

// this will keep alive forever, when channel c is closed
// it means we probably have to rebootstrap the lease

// need to receive here as per etcd docs

// Init starts the binding storage module
func (b *ETCDBindingStorage) Init() error { _ = "STUB: not implemented"; return nil }

// namespaced etcd :)

// Shutdown executes on shutdown and will clean etcd
func (b *ETCDBindingStorage) Shutdown() error { _ = "STUB: not implemented"; return nil }
