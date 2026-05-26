package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/topfreegames/pitaya/v3/pkg/acceptor"
	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/config"
)

// MetagameServer ...
type MetagameServer struct {
	component.Base
}

// NewMetagameMock ...
func NewMetagameMock() *MetagameServer { _ = "STUB: not implemented"; return nil }

// CreatePlayerCheatArgs is the struct used as parameter for the CreatePlayerCheat handler
// Using the 'validate' tag it's possible to add validations on all struct fields.
// For reference on the default validator see https://github.com/go-playground/validator.
// Also, to enable this validation pipeline see docs/configuration.rst.
type CreatePlayerCheatArgs struct {
	Name         string `json:"name"`
	Email        string `json:"email" validate:"email"`
	SoftCurrency int    `json:"softCurrency" validate:"gte=0,lte=1000"`
	HardCurrency int    `json:"hardCurrency" validate:"gte=0,lte=200"`
}

// CreatePlayerCheatResponse ...
type CreatePlayerCheatResponse struct {
	Msg string `json:"msg"`
}

// CreatePlayerCheat ...
func (g *MetagameServer) CreatePlayerCheat(ctx context.Context, args *CreatePlayerCheatArgs) (*CreatePlayerCheatResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The default logger contains a requestId, the route being executed and the sessionId

// Do nothing. This is just an example of how pipelines can be helpful

// HandlerNoArgResponse ...
type HandlerNoArgResponse struct {
	Msg string `json:"msg"`
}

// HandlerNoArg is a simple handler that do not have any arguments
func (g *MetagameServer) HandlerNoArg(ctx context.Context) (*HandlerNoArgResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Simple example of a before pipeline that actually asserts the type of the
// in parameter.
// IMPORTANT: that this kind of pipeline will be hard to exist in real code
// as a pipeline function executes for every handler and each of them
// most probably have different parameter types.
func (g *MetagameServer) simpleBefore(ctx context.Context, in interface{}) (context.Context, interface{}, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// Simple example of an after pipeline. The 2nd argument is the handler response.
func (g *MetagameServer) simpleAfter(ctx context.Context, resp interface{}, err error) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var app pitaya.Pitaya

func main() {
	svType := flag.String("type", "metagameDemo", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")
	flag.Parse()

	port := 3251
	metagameServer := NewMetagameMock()

	config := config.NewDefaultPitayaConfig()
	config.DefaultPipelines.StructValidation.Enabled = true

	builder := pitaya.NewDefaultBuilder(*isFrontend, *svType, pitaya.Cluster, map[string]string{}, *config)
	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", port))
	builder.AddAcceptor(tcp)
	builder.HandlerHooks.BeforeHandler.PushBack(metagameServer.simpleBefore)
	builder.HandlerHooks.AfterHandler.PushBack(metagameServer.simpleAfter)
	app = builder.Build()

	defer app.Shutdown()

	app.Register(metagameServer,
		component.WithName("metagameHandler"),
	)

	app.Start()
}
