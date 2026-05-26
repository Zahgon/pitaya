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

package cluster

import (
	"context"
	"sync"
	"time"

	"github.com/topfreegames/pitaya/v3/pkg/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type etcdServiceDiscovery struct {
	cli                    *clientv3.Client
	syncServersInterval    time.Duration
	heartbeatTTL           time.Duration
	logHeartbeat           bool
	lastHeartbeatTime      time.Time
	leaseID                clientv3.LeaseID
	mapByTypeLock          sync.RWMutex
	serverMapByType        map[string]map[string]*Server
	serverMapByID          sync.Map
	etcdEndpoints          []string
	etcdUser               string
	etcdPass               string
	etcdPrefix             string
	etcdDialTimeout        time.Duration
	running                bool
	server                 *Server
	stopChan               chan bool
	stopLeaseChan          chan bool
	lastSyncTime           time.Time
	listeners              []SDListener
	revokeTimeout          time.Duration
	grantLeaseTimeout      time.Duration
	grantLeaseMaxRetries   int
	grantLeaseInterval     time.Duration
	shutdownDelay          time.Duration
	appDieChan             chan bool
	serverTypesBlacklist   []string
	syncServersParallelism int
	syncServersRunning     chan bool
}

// NewEtcdServiceDiscovery ctor
func NewEtcdServiceDiscovery(
	config config.EtcdServiceDiscoveryConfig,
	server *Server,
	appDieChan chan bool,
	cli ...*clientv3.Client,
) (ServiceDiscovery, error) {
	_ = "STUB: not implemented"
	return *new(ServiceDiscovery), nil
}

func (sd *etcdServiceDiscovery) configure(config config.EtcdServiceDiscoveryConfig) {
	_ = "STUB: not implemented"
	return
}

func (sd *etcdServiceDiscovery) watchLeaseChan(c <-chan *clientv3.LeaseKeepAliveResponse) {
	_ = "STUB: not implemented"
	return
}

// renewLease reestablishes connection with etcd
func (sd *etcdServiceDiscovery) renewLease() error { _ = "STUB: not implemented"; return nil }

func (sd *etcdServiceDiscovery) grantLease() error {
	_ = "STUB: not implemented"
	// grab lease
	return nil
}

// this will keep alive forever, when channel c is closed
// it means we probably have to rebootstrap the lease

// need to receive here as per etcd docs

func (sd *etcdServiceDiscovery) addServerIntoEtcd(server *Server) error {
	_ = "STUB: not implemented"
	return nil
}

func (sd *etcdServiceDiscovery) bootstrapServer(server *Server) error {
	_ = "STUB: not implemented"
	return nil
}

// AddListener adds a listener to etcd service discovery
func (sd *etcdServiceDiscovery) AddListener(listener SDListener) { _ = "STUB: not implemented"; return }

// AfterInit executes after Init
func (sd *etcdServiceDiscovery) AfterInit() { _ = "STUB: not implemented"; return }

func (sd *etcdServiceDiscovery) notifyListeners(act Action, sv *Server) {
	_ = "STUB: not implemented"
	return
}

func (sd *etcdServiceDiscovery) writeLockScope(f func()) { _ = "STUB: not implemented"; return }

func (sd *etcdServiceDiscovery) deleteServer(serverID string) { _ = "STUB: not implemented"; return }

func (sd *etcdServiceDiscovery) deleteLocalInvalidServers(actualServers []string) {
	_ = "STUB: not implemented"
	return
}

func getKey(serverID, serverType string) string { _ = "STUB: not implemented"; return "" }

func getServerFromEtcd(cli *clientv3.Client, serverType, serverID string) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetServersByType returns a slice with all the servers of a certain type
func (sd *etcdServiceDiscovery) GetServersByType(serverType string) (map[string]*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a new map to avoid concurrent read and write access to the
// map, this also prevents accidental changes to the list of servers
// kept by the service discovery.

// GetServers returns a slice with all the servers
func (sd *etcdServiceDiscovery) GetServers() []*Server { _ = "STUB: not implemented"; return nil }

func (sd *etcdServiceDiscovery) bootstrap() error { _ = "STUB: not implemented"; return nil }

// GetServer returns a server given it's id
func (sd *etcdServiceDiscovery) GetServer(id string) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sd *etcdServiceDiscovery) InitETCDClient() error { _ = "STUB: not implemented"; return nil }

// namespaced etcd :)

// Init starts the service discovery client
func (sd *etcdServiceDiscovery) Init() error { _ = "STUB: not implemented"; return nil }

// update servers

func parseEtcdKey(key string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func parseServer(value []byte) (*Server, error) { _ = "STUB: not implemented"; return nil, nil }

func (sd *etcdServiceDiscovery) printServers() { _ = "STUB: not implemented"; return }

// Struct that encapsulates a parallel/concurrent etcd get
// it spawns goroutines and receives work requests through a channel
type parallelGetterWork struct {
	serverType string
	serverID   string
	payload    []byte
}

type parallelGetter struct {
	cli         *clientv3.Client
	numWorkers  int
	wg          *sync.WaitGroup
	resultMutex sync.Mutex
	result      *[]*Server
	workChan    chan parallelGetterWork
}

func newParallelGetter(cli *clientv3.Client, numWorkers int) parallelGetter {
	_ = "STUB: not implemented"
	return *new(parallelGetter)
}

func (p *parallelGetter) start() { _ = "STUB: not implemented"; return }

func (p *parallelGetter) waitAndGetResult() []*Server { _ = "STUB: not implemented"; return nil }

func (p *parallelGetter) addWorkWithPayload(serverType, serverID string, payload []byte) {
	_ = "STUB: not implemented"
	return
}

func (p *parallelGetter) addWork(serverType, serverID string) { _ = "STUB: not implemented"; return }

// SyncServers gets all servers from etcd
func (sd *etcdServiceDiscovery) SyncServers(firstSync bool) error {
	_ = "STUB: not implemented"
	return nil
}

// delete invalid servers (local ones that are not in etcd)

// Spawn worker goroutines that will work in parallel

// Check whether the server type is blacklisted or not

// Add new work to the channel

// Wait until all goroutines are finished

// BeforeShutdown executes before shutting down and will remove the server from the list
func (sd *etcdServiceDiscovery) BeforeShutdown() { _ = "STUB: not implemented"; return }

// Sleep for a short while to ensure shutdown has propagated

// Shutdown executes on shutdown and will clean etcd
func (sd *etcdServiceDiscovery) Shutdown() error { _ = "STUB: not implemented"; return nil }

// revoke prevents Pitaya from crashing when etcd is not available
func (sd *etcdServiceDiscovery) revoke() error { _ = "STUB: not implemented"; return nil }

func (sd *etcdServiceDiscovery) addServer(sv *Server) { _ = "STUB: not implemented"; return }

func (sd *etcdServiceDiscovery) watchEtcdChanges() { _ = "STUB: not implemented"; return }

// Block here if SyncServers() is running and consume the watcher channel after it's finished, to avoid conflicts

func (sd *etcdServiceDiscovery) isServerTypeBlacklisted(svType string) bool {
	_ = "STUB: not implemented"
	return false
}

func (sd *etcdServiceDiscovery) IsConnected(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}
