package pitaya

import (
	"github.com/topfreegames/pitaya/v3/pkg/acceptor"
	"github.com/topfreegames/pitaya/v3/pkg/cluster"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/conn/codec"
	"github.com/topfreegames/pitaya/v3/pkg/conn/message"
	"github.com/topfreegames/pitaya/v3/pkg/groups"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/metrics/models"
	"github.com/topfreegames/pitaya/v3/pkg/pipeline"
	"github.com/topfreegames/pitaya/v3/pkg/router"
	"github.com/topfreegames/pitaya/v3/pkg/serialize"
	"github.com/topfreegames/pitaya/v3/pkg/session"
	"github.com/topfreegames/pitaya/v3/pkg/worker"
)

// Builder holds dependency instances for a pitaya App
type Builder struct {
	acceptors        []acceptor.Acceptor
	postBuildHooks   []func(app Pitaya)
	Config           config.PitayaConfig
	DieChan          chan bool
	PacketDecoder    codec.PacketDecoder
	PacketEncoder    codec.PacketEncoder
	MessageEncoder   *message.MessagesEncoder
	Serializer       serialize.Serializer
	Router           *router.Router
	RPCClient        cluster.RPCClient
	RPCServer        cluster.RPCServer
	MetricsReporters []metrics.Reporter
	Server           *cluster.Server
	ServerMode       ServerMode
	ServiceDiscovery cluster.ServiceDiscovery
	Groups           groups.GroupService
	SessionPool      session.SessionPool
	Worker           *worker.Worker
	RemoteHooks      *pipeline.RemoteHooks
	HandlerHooks     *pipeline.HandlerHooks
}

// PitayaBuilder Builder interface
type PitayaBuilder interface {
	// AddPostBuildHook adds a post-build hook to the builder, a function receiving a Pitaya instance as parameter.
	AddPostBuildHook(hook func(app Pitaya))
	Build() Pitaya
}

// NewBuilderWithConfigs return a builder instance with default dependency instances for a pitaya App
// with configs defined by a config file (config.Config) and default paths (see documentation).
func NewBuilderWithConfigs(
	isFrontend bool,
	serverType string,
	serverMode ServerMode,
	serverMetadata map[string]string,
	conf *config.Config,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// NewDefaultBuilder return a builder instance with default dependency instances for a pitaya App,
// with default configs
func NewDefaultBuilder(isFrontend bool, serverType string, serverMode ServerMode, serverMetadata map[string]string, pitayaConfig config.PitayaConfig) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// NewBuilder return a builder instance with default dependency instances for a pitaya App,
// with configs explicitly defined
func NewBuilder(isFrontend bool,
	serverType string,
	serverMode ServerMode,
	serverMetadata map[string]string,
	config config.PitayaConfig,
) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// AddAcceptor adds a new acceptor to app
func (builder *Builder) AddAcceptor(ac acceptor.Acceptor) { _ = "STUB: not implemented"; return }

// AddPostBuildHook adds a post-build hook to the builder, a function receiving a Pitaya instance as parameter.
func (builder *Builder) AddPostBuildHook(hook func(app Pitaya)) { _ = "STUB: not implemented"; return }

// Build returns a valid App instance
func (builder *Builder) Build() Pitaya { _ = "STUB: not implemented"; return *new(Pitaya) }

// NewDefaultApp returns a default pitaya app instance
func NewDefaultApp(isFrontend bool, serverType string, serverMode ServerMode, serverMetadata map[string]string, config config.PitayaConfig) Pitaya {
	_ = "STUB: not implemented"
	return *new(Pitaya)
}

func configureDefaultPipelines(handlerHooks *pipeline.HandlerHooks) {
	_ = "STUB: not implemented"
	return
}

func addDefaultPrometheus(config config.MetricsConfig, customMetrics models.CustomMetricsSpec, reporters []metrics.Reporter, serverType string) []metrics.Reporter {
	_ = "STUB: not implemented"
	return nil
}

func addDefaultStatsd(config config.MetricsConfig, reporters []metrics.Reporter, serverType string) []metrics.Reporter {
	_ = "STUB: not implemented"
	return nil
}
