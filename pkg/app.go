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

package pitaya

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/topfreegames/pitaya/v3/pkg/acceptor"
	"github.com/topfreegames/pitaya/v3/pkg/cluster"
	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/errors"
	"github.com/topfreegames/pitaya/v3/pkg/groups"
	"github.com/topfreegames/pitaya/v3/pkg/interfaces"
	logging "github.com/topfreegames/pitaya/v3/pkg/logger/interfaces"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/router"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
	"github.com/topfreegames/pitaya/v3/pkg/service"
	"github.com/topfreegames/pitaya/v3/pkg/session"
	"github.com/topfreegames/pitaya/v3/pkg/worker"
	"go.opentelemetry.io/otel/trace"
)

// ServerMode represents a server mode
type ServerMode byte

const (
	_ ServerMode = iota
	// Cluster represents a server running with connection to other servers
	Cluster
	// Standalone represents a server running without connection to other servers
	Standalone
)

// Pitaya App interface
type Pitaya interface {
	GetDieChan() chan bool
	SetDebug(debug bool)
	SetHeartbeatTime(interval time.Duration)
	GetServerID() string
	GetMetricsReporters() []metrics.Reporter
	GetServer() *cluster.Server
	GetServerByID(id string) (*cluster.Server, error)
	GetServersByType(t string) (map[string]*cluster.Server, error)
	GetServers() []*cluster.Server
	GetSessionFromCtx(ctx context.Context) session.Session
	Start()
	SetDictionary(dict map[string]uint16) error
	AddRoute(serverType string, routingFunction router.RoutingFunc) error
	Shutdown()
	StartWorker()
	RegisterRPCJob(rpcJob worker.RPCJob) error
	Documentation(getPtrNames bool) (map[string]interface{}, error)
	IsRunning() bool

	RPC(ctx context.Context, routeStr string, reply proto.Message, arg proto.Message) error
	RPCTo(ctx context.Context, serverID, routeStr string, reply proto.Message, arg proto.Message) error
	ReliableRPC(
		routeStr string,
		metadata map[string]interface{},
		reply, arg proto.Message,
	) (jid string, err error)
	ReliableRPCWithOptions(
		routeStr string,
		metadata map[string]interface{},
		reply, arg proto.Message,
		opts *config.EnqueueOpts,
	) (jid string, err error)

	SendPushToUsers(route string, v interface{}, uids []string, frontendType string) ([]string, error)
	SendKickToUsers(uids []string, frontendType string) ([]string, error)

	GroupCreate(ctx context.Context, groupName string) error
	GroupCreateWithTTL(ctx context.Context, groupName string, ttlTime time.Duration) error
	GroupMembers(ctx context.Context, groupName string) ([]string, error)
	GroupBroadcast(ctx context.Context, frontendType, groupName, route string, v interface{}) error
	GroupContainsMember(ctx context.Context, groupName, uid string) (bool, error)
	GroupAddMember(ctx context.Context, groupName, uid string) error
	GroupRemoveMember(ctx context.Context, groupName, uid string) error
	GroupRemoveAll(ctx context.Context, groupName string) error
	GroupCountMembers(ctx context.Context, groupName string) (int, error)
	GroupRenewTTL(ctx context.Context, groupName string) error
	GroupDelete(ctx context.Context, groupName string) error

	Register(c component.Component, options ...component.Option)
	RegisterRemote(c component.Component, options ...component.Option)

	RegisterModule(module interfaces.Module, name string) error
	RegisterModuleAfter(module interfaces.Module, name string) error
	RegisterModuleBefore(module interfaces.Module, name string) error
	GetModule(name string) (interfaces.Module, error)

	GetNumberOfConnectedClients() int64
	IsReady(ctx context.Context) bool
}

// App is the base app struct
type App struct {
	acceptors         []acceptor.Acceptor
	config            config.PitayaConfig
	debug             bool
	dieChan           chan bool
	heartbeat         time.Duration
	onSessionBind     func(session.Session)
	router            *router.Router
	rpcClient         cluster.RPCClient
	rpcServer         cluster.RPCServer
	metricsReporters  []metrics.Reporter
	running           bool
	serializer        serialize.Serializer
	server            *cluster.Server
	serverMode        ServerMode
	serviceDiscovery  cluster.ServiceDiscovery
	startAt           time.Time
	worker            *worker.Worker
	remoteService     *service.RemoteService
	handlerService    *service.HandlerService
	handlerComp       []regComp
	remoteComp        []regComp
	modulesMap        map[string]interfaces.Module
	modulesArr        []moduleWrapper
	sessionModulesArr []sessionModuleWrapper
	groups            groups.GroupService
	sessionPool       session.SessionPool
}

// NewApp is the base constructor for a pitaya app instance
func NewApp(
	serverMode ServerMode,
	serializer serialize.Serializer,
	acceptors []acceptor.Acceptor,
	dieChan chan bool,
	router *router.Router,
	server *cluster.Server,
	rpcClient cluster.RPCClient,
	rpcServer cluster.RPCServer,
	worker *worker.Worker,
	serviceDiscovery cluster.ServiceDiscovery,
	remoteService *service.RemoteService,
	handlerService *service.HandlerService,
	groups groups.GroupService,
	sessionPool session.SessionPool,
	metricsReporters []metrics.Reporter,
	config config.PitayaConfig,
) *App {
	_ = "STUB: not implemented"
	return nil
}

// GetDieChan gets the channel that the app sinalizes when its going to die
func (app *App) GetDieChan() chan bool {
	_ = "STUB: not implemented"

	// SetDebug toggles debug on/off
	return nil
}

func (app *App) SetDebug(debug bool) {
	_ = "STUB: not implemented"

	// SetHeartbeatTime sets the heartbeat time
	return
}

func (app *App) SetHeartbeatTime(interval time.Duration) { _ = "STUB: not implemented"; return }

// GetServerID returns the generated server id
func (app *App) GetServerID() string { _ = "STUB: not implemented"; return "" }

// GetMetricsReporters gets registered metrics reporters
func (app *App) GetMetricsReporters() []metrics.Reporter { _ = "STUB: not implemented"; return nil }

// GetServer gets the local server instance
func (app *App) GetServer() *cluster.Server {
	_ = "STUB: not implemented"

	// GetServerByID returns the server with the specified id
	return nil
}

func (app *App) GetServerByID(id string) (*cluster.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetServersByType get all servers of type
func (app *App) GetServersByType(t string) (map[string]*cluster.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetServers get all servers
func (app *App) GetServers() []*cluster.Server { _ = "STUB: not implemented"; return nil }

// IsRunning indicates if the Pitaya app has been initialized. Note: This
// doesn't cover acceptors, only the pitaya internal registration and modules
// initialization.
func (app *App) IsRunning() bool {
	_ = "STUB: not implemented"

	// SetLogger logger setter
	return false
}

func SetLogger(l logging.Logger) { _ = "STUB: not implemented"; return }

func (app *App) initSysRemotes() { _ = "STUB: not implemented"; return }

func (app *App) periodicMetrics() { _ = "STUB: not implemented"; return }

// Start starts the app
func (app *App) Start() { _ = "STUB: not implemented"; return }

// set the service discovery as the last module to be started to ensure
// all modules have been properly initialized before the server starts
// receiving requests from other pitaya servers

// stop server

func (app *App) listen() { _ = "STUB: not implemented"; return }

// create global ticker instance, timer precision could be customized
// by SetTimerPrecision

// SetDictionary sets routes map
func (app *App) SetDictionary(dict map[string]uint16) error { _ = "STUB: not implemented"; return nil }

// AddRoute adds a routing function to a server type
func (app *App) AddRoute(
	serverType string,
	routingFunction router.RoutingFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown send a signal to let 'pitaya' shutdown itself.
func (app *App) Shutdown() { _ = "STUB: not implemented"; return }

// prevent closing closed channel

// Error creates a new error with a code, message and metadata
func Error(err error, code string, metadata ...map[string]string) *errors.Error {
	_ = "STUB: not implemented"
	return nil
}

// GetSessionFromCtx retrieves a session from a given context
func (app *App) GetSessionFromCtx(ctx context.Context) session.Session {
	_ = "STUB: not implemented"
	return *new(session.Session)
}

// GetDefaultLoggerFromCtx returns the default logger from the given context
func GetDefaultLoggerFromCtx(ctx context.Context) logging.Logger {
	_ = "STUB: not implemented"
	return *new(logging.Logger)
}

// AddMetricTagsToPropagateCtx adds a key and metric tags that will
// be propagated through RPC calls. Use the same tags that are at
// 'pitaya.metrics.additionalLabels' config
func AddMetricTagsToPropagateCtx(
	ctx context.Context,
	tags map[string]string,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// AddToPropagateCtx adds a key and value that will be propagated through RPC calls
func AddToPropagateCtx(ctx context.Context, key string, val interface{}) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetFromPropagateCtx adds a key and value that came through RPC calls
func GetFromPropagateCtx(ctx context.Context, key string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ExtractSpan retrieves an OpenTelemetry span context from the given context
// The span context can be received directly or via an RPC call
func ExtractSpan(ctx context.Context) (trace.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext), nil
}

// Documentation returns handler and remotes documentacion
func (app *App) Documentation(getPtrNames bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddGRPCInfoToMetadata adds host, external host and
// port into metadata
func AddGRPCInfoToMetadata(
	metadata map[string]string,
	region string,
	host, port string,
	externalHost, externalPort string,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Descriptor returns the protobuf message descriptor for a given message name
func Descriptor(protoName string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// StartWorker configures, starts and returns pitaya worker
func (app *App) StartWorker() {
	_ = "STUB: not implemented"

	// RegisterRPCJob registers rpc job to execute jobs with retries
	return
}

func (app *App) RegisterRPCJob(rpcJob worker.RPCJob) error { _ = "STUB: not implemented"; return nil }

// GetNumberOfConnectedClients returns the number of connected clients
func (app *App) GetNumberOfConnectedClients() int64 { _ = "STUB: not implemented"; return 0 }

// IsReady checks if pitaya is ready and able to serve requests
func (app *App) IsReady(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// Check NATS RPC Client connection

// Check NATS RPC Server connection

// Check ETCD connection
