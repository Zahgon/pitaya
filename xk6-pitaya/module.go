package pitaya

import (
	"github.com/grafana/sobek"
	"github.com/topfreegames/pitaya/v3/pkg/session"
	"go.k6.io/k6/js/modules"
)

type (
	// RootModule is the global module instance that will create Client
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		vu modules.VU
		*Client
		metrics *pitayaMetrics
	}
)

// Ensure the interfaces are implemented correctly
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance
func New() *RootModule { _ = "STUB: not implemented"; return nil }

// NewModuleInstance implements the modules.Module interface and returns
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	_ = "STUB: not implemented"
	return *new(modules.Instance)
}

// Exports implements the modules.Instance interface and returns
// the exports of the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	_ = "STUB: not implemented"
	return *new(modules.Exports)
}

// NewClient is the JS constructor function for the Client type.
// It returns a new Client instance for each VU.
// The first argument is an options object with the following fields:
// - handshakeData: the handshake data to send to the server
// - requestTimeoutMs: the timeout for requests in milliseconds
// - logLevel: the log level to use
func (mi *ModuleInstance) NewClient(call sobek.ConstructorCall) *sobek.Object {
	_ = "STUB: not implemented"
	return nil
}

type options struct {
	HandshakeData    *session.HandshakeData `json:"handshakeData"`
	RequestTimeoutMs int                    `json:"requestTimeoutMs"`
	UseTLS           bool                   `json:"useTLS"`
}

// newOptionsFrom validates and instantiates an options struct from its map representation
// as obtained by calling a Goja's Runtime.ExportTo.
func newOptionsFrom(argument map[string]interface{}) (*options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Instantiate a JSON decoder which will error on unknown
// fields. As a result, if the input map contains an unknown
// option, this function will produce an error.
